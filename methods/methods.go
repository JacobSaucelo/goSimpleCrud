package userMethods

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Users struct {
	users []User
}

var lastID int

func NewUser(userData []User) Users {
	return Users{
		users: userData,
	}
}

func (u *Users) CreateUser() {
	ClearScreen()

	fmt.Print("Enter user name: ")
	name := ReadInput()
	fmt.Print("Enter user age: ")
	ageStr := ReadInput()
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Invalid age. Please enter a valid number.")
		return
	}

	lastID++

	slug := User{ID: lastID, Name: name, Age: age}
	u.users = append(u.users, slug)

	fmt.Println("User created successfully.")
}

func (u *Users) ReadUsers() {
	ClearScreen()
	fmt.Println("Users:")
	for _, userItem := range u.users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", userItem.ID, userItem.Name, userItem.Age)
	}
}

func (u *Users) UpdateUsers() {
	ClearScreen()
	u.ReadUsers()

	fmt.Println("-------------------------------------")
	fmt.Print("Enter user ID to update: ")
	idStr := ReadInput()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid number.")
		return
	}

	var foundIndex int = -1
	for index, user := range u.users {
		if user.ID == id {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("User not found.")
		return
	}

	fmt.Print("Enter new user name: ")
	newName := ReadInput()
	fmt.Print("Enter new user age: ")
	newAgeStr := ReadInput()
	newAge, err := strconv.Atoi(newAgeStr)
	if err != nil {
		fmt.Println("Invalid age. Please enter a valid number.")
		return
	}

	u.users[foundIndex].Name = newName
	u.users[foundIndex].Age = newAge

	fmt.Println("User updated successfully.")
}

func (u *Users) DeleteUser() {
	ClearScreen()

	u.ReadUsers()

	fmt.Println("-------------------------------------")
	fmt.Print("Enter user ID to delete: ")
	idStr := ReadInput()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid number.")
		return
	}

	var foundIndex int = -1
	for index, user := range u.users {
		if user.ID == id {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("User not found.")
		return
	}

	//order matters (ordered but slower)
	// u.users = append(u.users[:foundIndex], u.users[foundIndex+1:]...)

	u.users[foundIndex] = u.users[len(u.users)-1]
	u.users = u.users[:len(u.users)-1]

	fmt.Println("User deleted successfully.")
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
