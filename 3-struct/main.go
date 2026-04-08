package main

import (
	"demo/CLI/file"
	"demo/CLI/storage"
	"fmt"
)

func main() {
	fileDb, err := file.NewFile("bins.json")
	if err != nil {
		fmt.Println(err)
	}
	store := storage.NewStorage(fileDb)
	store.WriteStorage()
	store.ReadStorage()
}
