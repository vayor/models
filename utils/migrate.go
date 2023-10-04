package main

import (
	"os"
	"fmt"
)

func main() {

	fmt.Println(os.Getenv("PROD_DB_PASSWORD"))

}