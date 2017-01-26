package main // ประกาศ package ครับผม
import "fmt" // อันนี้ประกาศเพื่อใช้ในการเเสดงผล รับค่า อะไรทำนองนั้นนะครับ


func main() {
	fmt.Println("Hello world This go - progrmming By Tak 😚")


	hello()
	
	
}

func hello () {
    var myname string
    fmt.Print("What is your name?  ")
    fmt.Scanf("%s", &myname)
    fmt.Println(myname)
}
