package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/XoliqberdiyevBehruz/wtc_backend/services/auth"
	types_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/user_admin"
)

func CreateSuperUser(store types_admin.UserStore) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("First Name: ")
	firstName, _ := reader.ReadString('\n')

	fmt.Print("Last Name: ")
	lastName, _ := reader.ReadString('\n')

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')

	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	hashPass, err := auth.GenerateHashPassword(password)
	if err != nil {
		log.Println(err)
	}
	user := types_admin.UserCreatePayload{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Password:  hashPass,
	}
	err = store.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("✅ Superuser created successfully!")
}
