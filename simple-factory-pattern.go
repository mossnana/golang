// Simple Factory Pattern from 
// จาก blog https://medium.com/@phayao/design-pattern-101-factory-pattern-a0a3f89cfc23

package main

import (
	"fmt"
)

// Pizza object
type Pizza struct {
	pizzaType string
	cheese int
	sausage int
	tomato int
}

// จัดเตรียมอัตราส่วนของพิซซ่า ถ้าเป็น พิซซ่าชีสให้ Cheese เป็น อัตราส่วน 2
func (pizza *Pizza) preparePizza() *Pizza {
	if(pizza.pizzaType == "cheese") {
		pizza.cheese = 2
		pizza.sausage = 1
		pizza.tomato = 1
	} else {
		pizza.cheese = 1
		pizza.sausage = 1
		pizza.tomato = 1
	}
	return pizza
}

// คำสั่ง สั่งพิซซ่า เพื่อไปเรียกการเตรียมวัตถุดิบให้กับพิซซ่าแต่ละแบบ โดยไม่ต้องมาแก้ไข เงื่อนไขวัตถุดิบใน orderPizza()
func (pizza *Pizza) orderPizza() bool {
	pizza.preparePizza()
	return true
}

func main() {
	fmt.Println("Test Factory Pattern")
  // สั่งพิซซ่าธรรมดา
	pizza := Pizza {
		pizzaType: "regular",
	}
  	pizza.orderPizza()
  
  // สั่งพิซซ่าชีส
	pizza2 := Pizza {
		pizzaType: "cheese",
	}
	pizza2.orderPizza()

	// Regular Pizza
	fmt.Println("Pizza Cheese", pizza.cheese)
	fmt.Println("Pizza Sausage", pizza.sausage)
	fmt.Println("Pizza Tomato", pizza.tomato)

	// Cheese Pizza
	fmt.Println("Pizza2 Cheese", pizza2.cheese)
	fmt.Println("Pizza2 Sausage", pizza2.sausage)
	fmt.Println("Pizza2 Tomato", pizza2.tomato)
}
