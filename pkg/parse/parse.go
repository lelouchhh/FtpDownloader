package parse

import (
	"fmt"
	"io/ioutil"
	"log"
)

func GetDir(path string) []string{
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var list []string
	for _, f := range files {
		fmt.Println(f.Name())
		if f.Name()[len(f.Name())-3:len(f.Name())] == "zip" {
			list = append(list, f.Name())
		}
	}
	fmt.Println(list)
	return list
}
