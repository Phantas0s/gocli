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

	data := struct {
		Name string
	}{"nwecli"}

	templateDir := "template"
	os.Mkdir(data.Name, 0775)
	err := filepath.Walk(templateDir,
		func(path string, info os.FileInfo, err error) error {
			newPath := strings.ReplaceAll(path, templateDir, data.Name)

			if err != nil {
				return err
			}

			if info.IsDir() {
				os.MkdirAll(newPath, 0775)
			}

			if !inArray(templates, path) && !info.IsDir() {
				input, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println(err)
					return nil
				}

				err = ioutil.WriteFile(newPath, input, 0775)
				if err != nil {
					fmt.Println("Error creating", newPath)
					fmt.Println(err)
					return nil
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	t, err := template.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}

	for _, v := range templates {
		_, file := filepath.Split(v)
		destination, err := os.Create(strings.ReplaceAll(v, templateDir, data.Name))
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
