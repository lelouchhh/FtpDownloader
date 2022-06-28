package main

import (
	"flag"
	"fmt"
	"../pkg/extract"
	"../pkg/parse"
	"os"
)

func main() {
	if len(os.Args) == 1{
		fmt.Println("Используйте 'cmd.exe -h' для получения справки")
		os.Exit(1)
	}
	delete := flag.Bool("delete", true, "Удалить zip файл?")
	destination := flag.String("output", "./", "Место для разархивации файла")
	path := flag.String("path", "./", "Путь до папки")


	flag.Parse()

	fmt.Println(parse.GetDir(*path))
	extract.Extract(parse.GetDir(*path), *path, *destination)
	if !*delete{
		extract.Remover(*path)
	}

}