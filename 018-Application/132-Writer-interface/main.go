package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground \n")

	je := json.NewEncoder(os.Stdout)
	je.Encode(`{name: "John", age: 31, city: "New York", salary: 1.01}`)

}
