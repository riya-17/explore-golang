package main
import "fmt"

const Pi = 3.14

func main() {
    fmt.Println("\nConstants\n")
    // **********************************************************************
    // Constants
    // **********************************************************************
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const a = 4
    var b int
    c := 5

    if a == 4 {
        b = c
        fmt.Println("b's value =", b)
    }

    /*
    will throw error as we are trying to assign value to constant (variable a)
    if a == 4 {
        a = c
        fmt.Println("b's value =", b)
    }
    */
}
