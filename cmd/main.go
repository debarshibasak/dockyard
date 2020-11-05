package main

import (
	"fmt"
	"github.com/adaptive-scale/dockyard/internal/configuration"
	"github.com/adaptive-scale/dockyard/internal/documentmanager"
	"github.com/adaptive-scale/dockyard/internal/server"
	"io/ioutil"
	"os"
	"path"
)

func main() {

	c := configuration.GetConfiguration()

	dManager := documentmanager.New(c)
	dManager.Reset()

	for r, f := range dManager.ListAllFiles() {
		menu, js := dManager.GenerateJS(f)
		tmpl := dManager.GenerateIndexHTML(menu, js)
		mkdir(r)
		write(r, tmpl)
	}

	fmt.Println("completed")

	if c.Serve {
		if err := server.New(documentmanager.OutputDir, ":10000").Start(); err != nil {
			panic(err)
		}
	}
}

func write(r string, tmpl string) {
	err := ioutil.WriteFile(path.Join(r, "index.html"), []byte(tmpl), os.FileMode(0777))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func mkdir( r string) {
	err := os.MkdirAll(r, os.FileMode(0700))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}