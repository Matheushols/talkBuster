package main

import (
	"fmt"
	"net/http"
	"talkBuster/router"
)

func main() {
	r := router.Start()
	fmt.Println("Api is running!")

	http.ListenAndServe(":8080", r)
}
