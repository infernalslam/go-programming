package main // ประกาศ package ครับผม
import "fmt" // อันนี้ประกาศเพื่อใช้ในการเเสดงผล รับค่า อะไรทำนองนั้นนะครับ


func main() {
	fmt.Println("Hello world This go - progrmming By Tak 😚")
    
	// // ประกาศตัวแปร 
	// var input float64
	// fmt.Print("Input number : ")
	// // input ต้องใช้ Scanf
	// fmt.Scanf("%f", &input)
	// fmt.Println("ผลรับของการรับค่า ", input)


	var myname string
    fmt.Print("What is your name?  ")
    fmt.Scanf("%s", &myname)
    fmt.Println("Hello", myname)

	
}
