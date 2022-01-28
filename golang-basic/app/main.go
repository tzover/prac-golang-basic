package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s san !", r.URL.Path[1:])
}

func main() {
	fmt.Println("Hello golang from docker!")

	var msg string = "Hello World"

	
	fmt.Println(msg)
	variable()
	constant()
	calc()
	pointer()
	fmt.Println("add(3 + 7) : ", addFunction(3, 7))
	array()
	process()
	structure()
	interfaceFunc()
	routine()
	
	// serverを立てる
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
