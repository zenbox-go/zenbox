package install_go

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	pb "gopkg.in/cheggaaa/pb.v1"
)

func unpack(dest, name string, fi os.FileInfo, r io.Reader) error {
	if strings.HasPrefix(name, "go/") {
		name = name[len("go/"):]
	}

	path := filepath.Join(dest, name)
	if fi.IsDir() {
		return os.MkdirAll(path, fi.Mode())
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fi.Mode())
	if err != nil {
		return err
	}
	defer f.Close()

	bar := pb.New64(fi.Size()).SetUnits(pb.U_BYTES)
	bar.Prefix(name)
	bar.Start()

	w := io.MultiWriter(f, bar)

	_, err = io.Copy(w, r)

	bar.Finish()
	return err
}

func unpackTar(src, dest string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	archive, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer archive.Close()

	tarReader := tar.NewReader(archive)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if err := unpack(dest, header.Name, header.FileInfo(), tarReader); err != nil {
			return err
		}
	}

	return nil
}

func unpackZip(src, dest string) error {
	zr, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	for _, f := range zr.File {
		fr, err := f.Open()
		if err != nil {
			return err
		}
		if err := unpack(dest, f.Name, f.FileInfo(), fr); err != nil {
			return err
		}
		fr.Close()
	}

	return nil
}
