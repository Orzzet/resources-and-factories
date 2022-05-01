package main

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"log"
	"net/http"
	"resourcesAndFactories/pkg/domain/services"
	"resourcesAndFactories/pkg/http/rest"
	badgerStorage "resourcesAndFactories/pkg/storage/badger"
	"time"
)

func run() error {
	db, err := badgerStorage.New()
	if err != nil {
		return err
	}
	defer func(db *badger.DB) {
		err := db.Close()
		if err != nil {
			log.Println("failed closing database")
			log.Fatal(err)
			return
		}
	}(db)
	service := services.New(db)
	err = service.EngineStartup()
	if err != nil {
		return err
	}

	go func() {
		for true {
			time.Sleep(1 * time.Second)
			err := service.EngineTick()
			if err != nil {
				log.Println("failed processing next tick")
				log.Fatal(err)
				return
			}
		}

	}()

	handler := rest.New(service)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8850", handler.Router); err != nil {
		log.Fatal("failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("-- Resources and Factories --")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
