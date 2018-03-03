package install_go

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func getGoVersions() ([]string, error) {
	vs := make([]string, 0)
	versionName := filepath.Join("cache", "VERSION")
	if fi, e := os.Stat(versionName); os.IsNotExist(e) {
		if err := getGoVersionsFormRemote(); err != nil {
			return nil, err
		}
	} else {
		// 缓存文件过期
		if time.Now().Sub(fi.ModTime()) > time.Hour*24*30 {
			os.Remove(versionName)
			if err := getGoVersionsFormRemote(); err != nil {
				return nil, err
			}
		}
	}

	b, err := ioutil.ReadFile(versionName)
	if err != nil {
		return nil, fmt.Errorf("读取版本缓存文件错误: %v", err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		vs = append(vs, scanner.Text())
	}

	return vs, nil
}

func getGoVersionsFormRemote() error {
	ctx := context.Background()
	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: "57412c43cdf95e5dbef55f5cc56501cbf2536f90",
		},
	)))

	os.Mkdir("cache", os.ModePerm)
	cacheFile, err := os.Create(filepath.Join("cache", "VERSION"))
	if err != nil {
		fmt.Printf("创建缓存文件错误: %v\n", err)
	}
	defer cacheFile.Close()

	w := bufio.NewWriter(cacheFile)

	page := 4
	size := 30
	// 4 ~ 7 页是需要的版本号,前面几页的版本号不需要
	for {
		if page > 7 {
			break
		}

		tags, _, err := client.Repositories.ListTags(ctx, "golang", "go", &github.ListOptions{
			Page:    page,
			PerPage: size,
		})
		if err != nil {
			// 读取任何一页错误,就全部错误,需要保持完整的版本号列表
			return err
		}

		page++

		for _, tag := range tags {
			tagName := tag.GetName()
			if strings.HasPrefix(tagName, "go") {
				fmt.Fprintln(w, tagName)
			}
		}

		time.Sleep(200 * time.Millisecond)
	}

	return w.Flush()
}
