package main

import (
	"PRIMEboard/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recovering...")
			os.Exit(1)
		}
	}()
	fmt.Println("Start server on 8080")
	http.HandleFunc("/get", handlers.HandleGet)
	http.HandleFunc("/set", handlers.HandleSet)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server started!")
}
