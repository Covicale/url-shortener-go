package main

import (
	"fmt"
	"log"

	"github.com/covicale/url-shortener-go/internal/api"
	"github.com/covicale/url-shortener-go/internal/config"
	"github.com/covicale/url-shortener-go/internal/db"
)

func main() {
	db, err := db.Connect(*config.Env.DB)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected succesfully.")

	address := fmt.Sprintf(":%v", config.Env.Port)
	server := api.NewAPIServer(address, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
