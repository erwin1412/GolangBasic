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

func main(){
	user1 := User{3,"Er","Win","erwin3@gmail.com",true}
	user2 := User{2,"Ac","Ca","erwin2@gmail.com",true}

	users := []User{user1,user2}

	Kelompok := Group{"Gamers" , user1 , users , true}

	displayGroup(Kelompok)
}

func displayGroup (Kelompok1 Group){
	fmt.Printf("Name : %s",Kelompok1.Name)
	fmt.Println()
	fmt.Printf("Admin : %s%s",Kelompok1.Admin.FirstName,Kelompok1.Admin.LastName)
	fmt.Println()
	fmt.Printf("Jumlah Member : %d",len(Kelompok1.Users))
	fmt.Println()
	fmt.Println("Members : ")
	for _,user := range Kelompok1.Users{
			fmt.Println(user.FirstName)
	}
}
