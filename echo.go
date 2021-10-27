package main

import (
	"fmt"
	"log"
)

func main() {
	var n []byte
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели слудющие данные: %s\n", n)
}
