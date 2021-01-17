package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "fut-bot ", log.LstdFlags)
	l.Println("Starting service...")
}
