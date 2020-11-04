package main

import (
	"flag"
	"fmt"
	"github.com/adaptive-scale/inventorize/dockyard/internal/documentmanager"
	"io/ioutil"
	"os"
	"path"
)

func main() {

	location := flag.String("location", "", "location of documentation")
	flag.Parse()

	if *location == "" {
		fmt.Println("location is empty")
		os.Exit(1)
	}

	dManager := documentmanager.New(*location)

	files, err := dManager.ListAllFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for r, f := range files {

		fmt.Println("==> Now generating "+r)
		start, menu, active := dManager.GetMenu(f)
		js := dManager.GenerateJS(start, active)
		tmpl, err := dManager.GenerateIndexHTML(menu, js)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.MkdirAll(r, os.FileMode(0700))

		err = ioutil.WriteFile(path.Join(r, "index.html"), []byte(tmpl), os.FileMode(0777))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}



	fmt.Println("completed")
}

type Index struct {
	JS   string
	CSS  string
	Menu string
}
