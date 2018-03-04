package tests

import (
	"context"
	"testing"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// 57412c43cdf95e5dbef55f5cc56501cbf2536f90
func TestTags(t *testing.T) {
	ctx := context.Background()
	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: "57412c43cdf95e5dbef55f5cc56501cbf2536f90",
		},
	)))

	// 4 ~ 7 页是需要的版本号
	tags, _, err := client.Repositories.ListTags(ctx, "golang", "go", &github.ListOptions{
		Page:    4,
		PerPage: 30,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(tags))

	for _, tag := range tags {
		t.Log(tag.GetName())
	}
}
