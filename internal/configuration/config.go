package configuration

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Configuration struct {
	Location string
	Branding string
	Theme string
	Serve bool
	Watch bool
}


func GetConfiguration() *Configuration {

	location := flag.String("location", "", "location of documentation")
	branding := flag.String("branding", "Acme", "branding of documentation")
	theme := flag.String("theme", "default", "only default supported at this point")
	serve := flag.Bool("serve", false, "generate and serve the documentation")
	watch := flag.Bool("watch", false, "watch for document change and regenerate when changed")
	version := flag.Bool("version", false, "shows the version of the release")

	flag.Parse()

	if *version {
		fmt.Println("0.1.1")
		os.Exit(0)
	}

	if *location == "" {
		err := errors.New(`
=========================
==> location is empty <==
=========================
`)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	return &Configuration{
		Branding:*branding,
		Location:*location,
		Theme:*theme,
		Serve:*serve,
		Watch:*watch,
	}
}
