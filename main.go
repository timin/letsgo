package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Create a reader to read input from stdin
    reader := bufio.NewReader(os.Stdin)
    
    // Prompt the user for their name
    fmt.Print("Please enter your name: ")
    
    // Read the input until newline character
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("An error occurred while reading input:", err)
        return
    }
    
    // Trim whitespace and newline characters from input
    name := strings.TrimSpace(input)
    
    // Print the greeting
    fmt.Printf("Hello %s\n", name)
}
