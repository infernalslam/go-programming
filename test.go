package main
import "fmt"

func main () {
    str:= ""
    age:= 0
    fmt.Print("Hello What's your name : ")
    fmt.Scanf("%s", &str)
    fmt.Println("Hello ::", str)

    fmt.Print("Age?")
    fmt.Scanf("%d", &age)
    fmt.Println("Age ===", age)

}
