package main


// import (
// 	"log"
// 	"fmt"
// 	// "strings"
// 	// "sort"
// 	"os"
// 	"io/ioutil"
// 	// "strconv"
// )

// //this is a comment
// /*
// this is a multiline comment
// */

// func main(){
// 	// fmt.Println("Hello World")

// 	// var age int = 40

// 	// var favNum float64 = 1.6180339


// 	// fmt.Println(age, favNum)

// 	// var numOne = 1.000
// 	// var num99 = .9999

// 	// fmt.Println(numOne - num99)

// 	// const pi float64 = 3.14159265

// 	// var (
// 	// 	varA = 2
// 	// 	varB = 3
// 	// )

// 	// var myName string = "Name"

// 	// fmt.Println(len(myName + " is a name."))

// 	// var isOld bool = true

// 	// rect1 := Rectangle{0, 5, 10, 10}

// 	// fmt.Println(rect1.height)
// 	// 	fmt.Println(rect1.area())

// 	// sampString := "Hello World"

// 	// fmt.Println(strings.Contains(sampString, "lo"))
// 	// fmt.Println(strings.Index(sampString, "lo"))
// 	// fmt.Println(strings.Count(sampString, "l"))
// 	// fmt.Println(strings.Replace(sampString, "l", "x", 3))

// 	file, err := os.Create("samp.txt")

// 	if err != nil{
// 		log.Fatal(err)
// 	}

// 	file.WriteString("This is some text")

// 	file.Close()

// 	stream, err := ioutil.ReadFile("samp.txt")

// 	if err!= nil{
// 		log.Fatal(err)
// 	}

// 	readString := string(stream)

// 	fmt.Println(readString)

// }

// type Rectangle struct{
// 	leftX float64
// 	topY float64
// 	height float64
// 	width float64
// }

// //use this to attach method onto struct
// func (rect *Rectangle) area() float64{
// 	return rect.width * rect.height
// }