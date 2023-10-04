package main

import (
	"os"
	"log"
)

func main() {

	log.Println(os.Getenv("DB_PASSWORD"))

}