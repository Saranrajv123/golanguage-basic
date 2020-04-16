package main

import "fmt"

// execute after the main function end
// panic happen after the defer execution

func main() {
//     fmt.Println("first")
//        defer fmt.Println("middle")
//         fmt.Println("last")

        //a := "start"
        //defer fmt.Println(a, "a")
        //a = "end"
        //
        //http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        //    writer.Write([]byte("Hello go"))
        //})
        //
        //err := http.ListenAndServe(":8080", nil)
        //if err != nil {
        //    panic("error")
        //}

    fmt.Println("start")
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Error", err)
            fmt.Println("inside defer if loop")
        }
    }()

    panic("something went wrong")
    fmt.Println("end")





}