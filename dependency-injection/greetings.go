package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// It's hard to test this guy
func OldGreet(name string) {
	fmt.Printf("Good morning, %s", name)
}

// Let's inject the dependency of printing
func Greeting(writer io.Writer, name string) {
	msg := fmt.Sprintf("Hello, %s", name)
	fmt.Fprintf(writer, msg)
} 

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greeting(w, "world")
}

func main() {
	Greeting(os.Stdout, "World")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
} 