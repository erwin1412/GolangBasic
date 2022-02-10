package main

import "fmt"
type User struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	isActivate bool
}

type Group struct{
	Name string
	Admin User
	Users []User
	isAvailable bool
}

func (user User) display() string{
	return fmt.Sprintf("Name %s %s , email %s",user.FirstName,user.LastName,user.Email)
}

func main(){

	//cara 1
	user1 := User{1,"Er","Win","erwin3@gmail.com",true}
	result := user1.display()
	fmt.Println(result)

	//cara 2
	user2 := User{2,"Ac","Ca","AcCa@gmail.com",true}
	fmt.Println(user2.display())
}

