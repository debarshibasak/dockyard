package main

import (
	"fmt"
	"github.com/adaptive-scale/dockyard/internal/configuration"
	"github.com/adaptive-scale/dockyard/internal/documentmanager"
	"github.com/adaptive-scale/dockyard/internal/server"
	"sync"
)

func main() {

	c := configuration.GetConfiguration()

	dManager := documentmanager.New(c)
	dManager.Generate()

	fmt.Println("completed")

	var wait sync.WaitGroup
	wait.Add(1)

	if c.Watch {

		go dManager.Watch()
	}

	if c.Serve {
		go func() {
			if err := server.New(documentmanager.OutputDir, ":10000").Start(); err != nil {
				panic(err)
			}
		}()
	}

	wait.Wait()
}
