package main

import "fmt"
type User struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	isActivate bool
}

func main(){
	user1 := User{3,"Er3","Win3","erwin3@gmail.com",true}
	user2 := User{2,"Er2","Win2","erwin2@gmail.com",true}
	
	display1 := DisplayUser(user1)
	display2 := DisplayUser(user2)
	fmt.Println(display1)
	fmt.Println(display2)
}

func DisplayUser (user User)string{
	hasil := fmt.Sprintf("Name %s %s , email %s",user.FirstName,user.LastName,user.Email)
	return hasil
}