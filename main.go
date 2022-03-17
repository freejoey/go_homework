package main

import (
	"homework/account"
	"log"
)

func main() {
	log.Printf("starting...")

	log.Printf("test db connection: %t", account.GetAccountCount() >= 0)
}
