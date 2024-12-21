package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file:///Users/macbook/AllRepository/goLang/belajar-golang-rest-api/cmd/migrate/migrations",
		"postgres://postgres:tangerang12@localhost:5432/belajargolangrestapi?sslmode=disable",
	)
	if err != nil {
		log.Fatal("error initializie migration : ", err)
	}

	cmd := os.Args[len(os.Args) - 1]

	if cmd == "up" {
		if err := m.Up(); err != nil { 
			if err == migrate.ErrNoChange {
				log.Println("migrate up : no changes") 
			} else {
				log.Fatal("migrate up success")
			}
		} else {
			log.Println("migrate succes")	
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil  {
			if err == migrate.ErrNoChange {
				log.Println("migrate down : no changes")
			} else {
				log.Fatal("migrate down error : ", err) 
			}
		} else {
			log.Println("migrate down success")
		}
	}
}