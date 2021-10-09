package main

import (
	"log"

	"github.com/onituka/agile-project-management/project-management/infrastructure/router"
)

func main() {
	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}
}
