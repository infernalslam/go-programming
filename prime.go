package main // ประกาศ package ครับผม
import "fmt" // อันนี้ประกาศเพื่อใช้ในการเเสดงผล รับค่า อะไรทำนองนั้นนะครับ


func main() {
	fmt.Println("Hello world This go - progrmming By Tak 😚")
    
    var input int
	fmt.Print("> prime_factors  ")
	fmt.Scanf("%d\n", &input)
	fmt.Println("ผลรับของการรับค่า ", input)

    var prime = Prime(input)
    if (prime) {
        fmt.Print(input)
    } else if (prime == false) {
        fmt.Print("con..")
        divide(input)
    }

    fmt.Print(prime)
	
    fmt.Println("")
}

func Prime (input int) bool {
    var check int = 0
    var checkVail bool
    fmt.Print(input)
    for i:= input - 1; i > 1; i-- {
        if (input % i == 0) {
            check = 1
        } else {
            check = 2
        }
    }

    if (check == 1) {
        // fmt.Println("NotPrime_func")
        checkVail = false
    } else if (check == 2 || input == 2) {
        // fmt.Println("Prime_func")
        checkVail = true
    }

    return checkVail
}

func divide (input int) int {

    for ; i < 10;  {}

    return 0
}

