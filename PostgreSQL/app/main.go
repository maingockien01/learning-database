package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main () {
    generatePool()
    defer closePool()

    reader := bufio.NewReader (os.Stdin)
    fmt.Println("Welcome to Simple Book library app!")
    for  {
        text, _ := reader.ReadString('\n') 
        text = strings.Replace(text, "\n", "", -1)
        if text == "quit" {
            fmt.Println("App is quiting")
            break
        }

        command := ReadLine(text)

        result := command.PerformCommand()
        fmt.Println("Result: ")
        fmt.Println(result)
    }
}

