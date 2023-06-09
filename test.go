package main

import (
	"fmt"
	"math"
)

//FUNCTIONS
func sub(x int, y int) int{
	return x - y
}

//multiple return values
func fn(str string) (string, string){
	return str + "1", str + "1"
}

func getCords()(x, y int){
	return  //automatically returns x and y
}


//STRUCTS
type User struct {
	username string
	age int
	password string
}

type rect struct{
	width float64
	height float64
}

//METHODS
func (r rect) area() float64{
	return r.width * r.height
}

//INTERFACES //collection of methods
type shape interface {
	area() float64
	perimeter() float64
}

type circle struct{
	radius float64
}

func (r rect) perimter() float64{
	return 2*(r.width + r.height)
}

func (c circle) area() float64{
	return math.Pi * c.radius *c.radius
}

func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

func main() {
	//NORMAL ASSIGNMENT
	// var username string = "Mangu"
	// var password string = "sahfjksdhfds"

	// //SHORT ASSIGNEMNT
	// congrats := "Happy Birthday"
	// fmt.Println("Credentials: \nUsername - ", username + "\nPassword - ", password)
	// fmt.Println(congrats)

	//CONDITIONALS
	// length := 4

	// if length > 3 {
	// 	fmt.Printf("The length %d is greater than 3", length)
	// }

	// if len := 5; len > length {
	// 	fmt.Printf("Length: %d", len)
	// } 

	//FUNCTIONS
	// ans, _ := fn("0")
	// fmt.Println(ans)

	// r := rect{
	// 	width: 3.0,
	// 	height: 5.0,
	// }

	// fmt.Println(r.area());

	//ERROR HANDLING
	// user, err := getUser();

	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }

	// func getUser()(User, error){
	// 	return User{"Mangu", 20, "1234"}, nil
	// }

	//LOOPS
	//normal loop
	// for i:=0; i<6; i++{
	
	// }


	//ARRAYS
	// arr := [5]string{"a", "b", "c", "d", "e"} //array
	// arr2 := [4]int{1,23,4}

	//SLICES
	slc := []int{2,5,4}
	// slc2 := arr2[0:2]
	// slice := make([]int, 5, 10) //slice with length 5 and capacity 10
	// slice2 := make([]int, 5) //slice with length 5 and capacity 5

	for i, word := range slc {
		fmt.Println(i, " ",word)
	}

	fmt.Println(slc)

}