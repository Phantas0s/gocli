package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	templates := []string{
		"template/main.go",
		"template/go.mod",
		"template/cmd/example.go",
		"template/cmd/root.go",
	}

	if len(os.Args) == 0 {
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
			newPath := cliPath + string(filepath.Separator) + file.Name()
			if file.IsDir() {
				os.MkdirAll(newPath, 0750)
			}

			if !inArray(templates, path) && !file.IsDir() {
				input, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println(err)
					return nil
				}

				err = ioutil.WriteFile(newPath, input, 0750)
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
		_, file := filepath.Split(v)
		destination, err := os.Create(cliPath + string(filepath.Separator) + file)
		defer destination.Close()
		if err != nil {
			return
		}
		err = t.ExecuteTemplate(destination, file, data)
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
