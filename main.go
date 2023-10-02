package main

import (
	"flag"
	"fmt"
	"learning-go/api"
	"learning-go/storage"
	"log"
)

func main(){
	listenAddr := flag.String("listenAddress", ":3000", "the server address")
	flag.Parse()
	store := storage.NewMemoryStorage()
	server := api.NewServer(*listenAddr, store) 
	fmt.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())
}