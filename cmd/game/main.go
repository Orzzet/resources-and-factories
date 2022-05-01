package main

import (
	"fmt"
	"resourcesAndFactories/pkg/storage/badger"
)

func main() {
	db, err := badger.New()
	if err != nil {
		fmt.Println("couldn't initialize database")
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
}
