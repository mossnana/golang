package main

// import เฉพาะตัวที่จำเป็นเท่านั้น
import (
	"fmt"
	"log"
)

func main(){


	// Statement
	fmt.Println("Hello, Permpoon")
	// Hello, Permpoon
	log.Println("Hello, log");
	// 2019/01/25 14:38:04 Hello, log
	println("Print Built-in function")
	// Print Built-in function
	// *ไม่แนะนำให้ใช้*


	// Variable
	var permpoon string
	// ประกาศแล้ว ต้องเอาไปใช้
	permpoon = "Permpoon"
	fmt.Printf("Hello again, %s. \n", permpoon)
	// Hello again, Permpoon.
	var name = "Permpoon Chao"
	// name จะเป็น string ไปโดยอัตโนมัติ
	fmt.Printf("My Name is %s \n", name)
	// My Name is Permpoon Chao
	// กำหนด name = 10 จะ Error ทันที
	nickName := "Moss"
	// ใช้ := แทน var นำหน้าได้
	fmt.Printf("My Nickname is %s\n", nickName)
	// My Nickname is Moss
	_ = nickName
	// ใช้ _ = nickName ในตอนยังไม่ได้ใช้ nickName มันจะไม่ Error


	// Example#1 ตัวอย่างใช้ control flow: if, else if, else, switch...case
	fmt.Print("Input a fruit: ")
	var fruit string
	// fmt.Scanln(&fruit) *
	// ส่งค่า Pointer ของ Fruit เข้าไป

	// check string เปล่าแบบที่ 1
	/*if fruit == "" {
		fmt.Println("Null Check by String Value")
		return
	}

	// check string เปล่าแบบที่ 2
	if len(fruit) == 0{
		fmt.Println("Null Check by String Length")
		return
	}*/

	// switch...case ใน Golang ไม่ต้องใส่ Break
	switch fruit {
	case "apple":
		fmt.Println("Apple")
	case "banana":
		fmt.Println("Banana")
	case "lemon":
		fmt.Println("Lemon")
	default:
		fmt.Println("Others")
	}
	/*
	ใน case สามารถ ใส่ statement ได้ เช่น
	case variableX < 50 :
		...
	*/



	// Array
	var a [5]int
	b := [3]int{1, 2, 3}

	// Assign Array
	a[0] = 10
	// ตัว a[1] ถ้าไม่กำหนดจะเป็น 0
	a[2] = 5
	a[3] = 1
	a[4] = 8
	fmt.Println(a)
	// [10 0 5 1 8]
	fmt.Println(len(a))
	// 5

	// ++i ไม่มี

	// ลูปค่าใน Array แบบที่ 1
	for i :=0; i< len(a); i++ {
		fmt.Println(a[i])
	}

	// ลูปค่าใน Array แบบที่ 2
	for i:= range a {
		fmt.Println(a[i])
	}

	// ลูปค่าใน Array แบบที่ 3
	// v คือ value ใน array
	for _, v:= range a {
		fmt.Println(v)
	}

	/*
	10
	0
	5
	1
	8
	*/

	fmt.Println(a[2:4])
	// [5 1]
	fmt.Println(a[2:])
	// [5 1 8]
	fmt.Println(a[:4])
	// [10 0 5 1]


	fmt.Println(b)
	// [1 2 3]


	// Slide เพิ่มขนาดนั้น
	c := make([]int, 5)
	c[0] = 10
	c[2] = 30
	c[3] = 40
	fmt.Println(c)
	// [10 0 30 40 0]
	c = append(c, 90)
	fmt.Println(c)
	// [10 0 30 40 0 90]


	// Map
	d := make(map[string]string)
	d["hello"] = "Permpoon"
	d["name"] = "test"
	fmt.Println(d["hello"])

	// search in map
	x, ok := d["name"]
	// ค่า ok จะเป็น Boolean เจอเป็น True ไม่เจอเป็น False
	if ok {
		fmt.Println(x)
	} else {
		fmt.Println(false)
	}

	// Delete ค่า
	delete(d, "hello")

	// Pointer
	e := 10
	ptrE := &e
	fmt.Println(ptrE)
	// 0xc000016210
	// เป็นค่าในหน่วยความจำ
	fmt.Println(*ptrE)
	// 10




	// 57:17 / 2:04:00
}
