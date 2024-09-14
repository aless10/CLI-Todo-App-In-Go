package main

import (
	"github.com/aless10/CLI-Todo-App-In-Go/cli"
	"github.com/aless10/CLI-Todo-App-In-Go/cmd"
	"github.com/aless10/CLI-Todo-App-In-Go/storage"
	"log"
	"os"
	"path/filepath"
)

func main() {
	todos_storage_path := os.Getenv("TODOS_STORAGE")
	if todos_storage_path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		todos_storage_path = filepath.Join(home, ".todos.json")
	}
	todos_storage := storage.NewStorage[todos.Todos](todos_storage_path)
	todos := &todos.Todos{}
	todos_storage.Load(todos)
	cmdFlags := cli.NewCmdFlags()
	cmdFlags.Execute(todos)
	todos_storage.Save(*todos)
}
