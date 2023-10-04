package main

import (
	"os"
	"fmt"
)

func main() {

	fmt.Println(os.Getenv("DB_PASSWORD"))

}