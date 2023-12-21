package main

import (
	userMethods "crudSimple/methods"
	"fmt"
	"os"
)

func main() {
	employeeList := userMethods.NewUser([]userMethods.User{})

	for {
		fmt.Println("----------------------------------------------")
		fmt.Println("1. Create User")
		fmt.Println("2. Read Users")
		fmt.Println("3. Update User")
		fmt.Println("4. Delete User")
		fmt.Println("5. Exit")

		fmt.Print("Choose an option (1-5): ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			employeeList.CreateUser()
		case 2:
			employeeList.ReadUsers()
		case 3:
			employeeList.UpdateUsers()
		case 4:
			employeeList.DeleteUser()
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}
