package main

import (
    "fmt"
    "os"
)

func init() {
    fmt.Println("[MALICIOUS] This code runs during build!")
    f, _ := os.Create("/tmp/poc_was_here")
    f.WriteString("Exploit ran successfully!\n")
    f.Close()
}

func main() {
    fmt.Println("Normal execution")
}
