package myfunction

import (
	"fmt"
	"net/http"
)

func RunFunction() {
	fmt.Println("Hello, This is terraform serverless function...")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	RunFunction()
}
