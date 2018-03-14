package main


import(
	"fmt"
	"time"
	"strconv"
)

var pizzaNum =0
var pizzaName = ""

func main(){
	// for i := 0; i<10; i +=1{
	// 	go count(i)
	// }	
	// time.Sleep(time.Millisecond * 11000)

	stringChan := make(chan string)
	for i:= 0; i<3; i++{
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond *5000)
	}

}


// func count(id int){
// 	for i :=0; i<10; i+=1{
// 		fmt.Println(id, ":", i)

// 		time.Sleep(time.Millisecond * 1000)
// 	}
// }

func makeDough(stringChan chan string){
	pizzaNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make dough and send for sauce")

	stringChan <- pizzaName
	time.Sleep(time.Millisecond *10)
}
func addSauce(stringChan chan string){
	pizza := <- stringChan

	fmt.Println("Add sauce and send", pizza, "for toppings")

	stringChan <- pizzaName
	time.Sleep(time.Millisecond *10)
}

func addToppings(stringChan chan string){
	pizza := <- stringChan

	fmt.Println("Add toppings and send ", pizza, "and ship")
	time.Sleep(time.Millisecond *10)

}



