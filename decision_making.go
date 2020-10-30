package main
import "fmt"

func main() {
    x, y := 6, 4

    fmt.Println("\n\nConditions\n")
    // **********************************************************************
    // Conditions
    // **********************************************************************
    if x > 0 {
        fmt.Println("x is big")
    }

    if x > y {
        fmt.Println("x is bigger than y")
    } else {
        fmt.Println("y is bigger than x")
    }

    if x > 7 {
        fmt.Println("x is big")
    } else if y > 1 && y < 5 {
        fmt.Println("y is just right")
    } else {
        fmt.Println("y is big")
    }


    fmt.Println("\n\nSwitch Condition\n")
    // **********************************************************************
    // Switch: we can even add conditions in cases, fallthrough as we desire
    // **********************************************************************
    switch x{
    case 6:
        fmt.Println("switch result : seven")
    case 8:
        fmt.Println("switch result : eight")
    default:
        fmt.Println("none of them\n")
    }

    // fallthrough
    switch x{
    case 6:
        fmt.Println("switch result : seven")
        fallthrough
    case 8:
        fmt.Println("switch result : eight")
    default:
        fmt.Println("none of them\n")
    }

    // conditional
    switch {
    case x == 6:
        fmt.Println("switch result : seven")
    case x > 10:
        fmt.Println("switch result : eight")
    default:
        fmt.Println("none of them\n")
    }

    // type-specific: Type switch is used when you want to compare types.
    // In this switch, the case contains the type which is going to compare 
    // with the type present in the switch expression.
}
