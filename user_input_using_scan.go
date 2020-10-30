package main
import "fmt"

func main() {
    // **********************************************************************
    // Scan function: to take input from user
    // **********************************************************************
    fmt.Println("\n\nScan\n")
    var p int
    var q int
    fmt.Print("enter p's value: ")
    fmt.Scan(&p)
    fmt.Print("enter q's value: ")
    fmt.Scan(&q)
    fmt.Println("p's value : ", p)
    fmt.Println("q's value : ", q)
}
