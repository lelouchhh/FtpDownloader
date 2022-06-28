package extract

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Extract(list []string, path, dest string) {
	for _, file := range list {

		dst := dest+file[0:len(file)-7]
		fmt.Println("path+file", path, file)
		archive, err := zip.OpenReader(path+file)
		if err != nil {
			panic(err)
		}
		defer archive.Close()

		for _, f := range archive.File {
			filePath := filepath.Join(dst, f.Name)
			fmt.Println("unzipping file ", filePath)

			if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
				fmt.Println("invalid file path")
				return
			}
			if f.FileInfo().IsDir() {
				fmt.Println("creating directory...")
				os.MkdirAll(filePath, os.ModePerm)
				continue
			}

			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				panic(err)
			}

			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				panic(err)
			}

			fileInArchive, err := f.Open()
			if err != nil {
				panic(err)
			}

			if _, err := io.Copy(dstFile, fileInArchive); err != nil {
				panic(err)
			}

			dstFile.Close()
			fileInArchive.Close()
		}
	}
}
func Remover(name string) error{
	err := os.Remove(name)
	if err != nil {
		return err
	}
	return err
}