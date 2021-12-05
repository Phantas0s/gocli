package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	templates := []string{
		"template/main.go",
		"template/go.mod",
		"template/cmd/example.go",
		"template/cmd/root.go",
	}

	if len(os.Args) < 2 {
		panic("You need to give the path of the new CLI as argument.")
	}

	cliPath := os.Args[1]
	_, name := filepath.Split(cliPath)

	data := struct {
		Name string
	}{name}

	templateDir := "template"
	os.Mkdir(cliPath+string(filepath.Separator)+data.Name, 0750)
	walkFunc := func(cliPath string) func(path string, file os.FileInfo, err error) error {
		return func(path string, file os.FileInfo, err error) error {
			tailPath := strings.Split(path, string(filepath.Separator))[1:]
			newPath := cliPath + string(filepath.Separator) + filepath.Join(tailPath...)
			if file.IsDir() {
				os.MkdirAll(newPath, file.Mode())
			}

			if !inArray(templates, path) && !file.IsDir() {
				input, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println(err)
					return nil
				}

				err = ioutil.WriteFile(newPath, input, file.Mode())
				if err != nil {
					fmt.Println("Error creating", newPath)
					fmt.Println(err)
					return nil
				}
			}

			return nil
		}
	}
	err := filepath.Walk(templateDir, walkFunc(cliPath))
	if err != nil {
		log.Println(err)
	}

	t, err := template.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}

	for _, v := range templates {
		tailPath := strings.Split(v, string(filepath.Separator))[1:]
		destination, err := os.Create(cliPath + string(filepath.Separator) + filepath.Join(tailPath...))
		defer destination.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = t.ExecuteTemplate(destination, tailPath[len(tailPath)-1], data)
	}

	if err != nil {
		panic(err)
	}
}

func inArray(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
