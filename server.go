package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	w.Write([]byte(fmt.Sprintf("Hello, %s. I am %s years old.", name, age)))
}

func secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	w.Write([]byte(fmt.Sprintf("User: %s, Password: %s", user, pass)))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	fmt.Fprint(w, string(data))
}
