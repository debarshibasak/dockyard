package main

import (
	"flag"
	"fmt"
	"github.com/adaptive-scale/dockyard/internal/documentmanager"
	"github.com/adaptive-scale/dockyard/internal/server"
	"io/ioutil"
	"os"
	"path"
)

func main() {

	location := flag.String("location", "", "location of documentation")
	branding := flag.String("branding", "Acme", "branding of documentation")
	theme := flag.String("theme", "default", "only default supported at this point")
	serve := flag.Bool("serve", false, "generate and serve the documentation")
	flag.Parse()

	if *location == "" {
		fmt.Println("=========================")
		fmt.Println("==> location is empty <==")
		fmt.Println("=========================")
		os.Exit(1)
	}

	dManager := documentmanager.New(*location)

	files, err := dManager.ListAllFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for r, f := range files {
		fmt.Println("-------------------------------")
		fmt.Println("==> Now generating "+r)
		start, menu, active := dManager.GetMenu(f)
		js := dManager.GenerateJS(start, active)
		tmpl, err := dManager.GenerateIndexHTML(*theme, *branding, menu, js)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = os.MkdirAll(r, os.FileMode(0700))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = ioutil.WriteFile(path.Join(r, "index.html"), []byte(tmpl), os.FileMode(0777))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("")

	}

	fmt.Println("completed")

	if *serve {
		if err := server.New(documentmanager.OutputDir, ":10000").Start(); err != nil {
			panic(err)
		}
	}
}