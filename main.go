package main

import (
	"flag"
	"log"
)

var who = flag.String("who", "world", "Say hello to who")

func main() {
	flag.Parse()
	log.Println("Hello,", *who)
}
