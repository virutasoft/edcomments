package main

import (
	"flag"
	"log"

	"github.com/golang-es/edcomments/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la base de datos")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Inició la migración de la BD...")
		migration.Migrate()
		log.Println("Finalizó la migración de la BD.")
	}
}

// ./edcomments --migrate yes
