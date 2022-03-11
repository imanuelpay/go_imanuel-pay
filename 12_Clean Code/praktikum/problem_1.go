package main

type User struct {
	id       int
	username string
	password string
}

type UserService struct {
	ListUser []User
}

func (userService UserService) getAllUsers() []User {
	return userService.ListUser
}

func (userService UserService) getUserById(id int) User {
	for _, user := range userService.ListUser {
		if id == user.id {
			return user
		}
	}

	return User{}
}
