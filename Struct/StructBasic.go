package main

import "fmt"

type User struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	isActivate bool
}

func main() {
	// cara 1 Manual
	user1 := User{}
	user1.ID = 1
	user1.FirstName = "Er1"
	user1.LastName = "Win1"
	user1.Email = "erwin@gmail.com"
	user1.isActivate = true

	//cara2 menggunakan MAP
	user2 := User{
		ID : 2,
		FirstName : "Er2",
		LastName : "Win2",
		Email : "erwin2@gmail.com",
		isActivate : true,
	}
	// cara 3
	user3 := User{3,"Er3","Win3","erwin3@gmail.com",true}
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)

}