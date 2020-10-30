package main
import "fmt"

func main() {
    fmt.Println("\n\nLoop\n")
    // **********************************************************************
    // For loop: while, do-while as well
    // **********************************************************************
    for i:=0; i<10; i++ {
        if i == 7 {
            fmt.Println("break")
            break
        }
        fmt.Println(i)
    }
    fmt.Println("\n")

    for i:=0; i<10; i++ {
        if i == 7 {
            fmt.Println("continue")
            continue
        }
        fmt.Println(i)
    }

    fmt.Println("\n\nwhile\n")
    // while from for loop
    j := 10
    for j<20 {
        fmt.Println(j)
        j++
    }

    fmt.Println("\n\nDo While\n")
    // Do-while from for loop
    k :=21
    for {
        fmt.Println(k)
        k++
        if k > 30 {
            break
        }
    }
}
