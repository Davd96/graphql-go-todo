package models

import (
	"fmt"
)

type User struct {
	ID   int32  `db:"id"`
	Name string `db:"name"`
}

type NewUser struct {
	Name string
}

type UserService struct{}

func (c UserService) GetAll() []User {
	var user []User
	err := GetConnection().Select(&user, "SELECT id,name FROM public.user")
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}

func (c UserService) GetById(id int32) []User {
	var user []User
	err := GetConnection().Select(&user, fmt.Sprintf("SELECT id,name FROM public.user WHERE id = %v", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}

func (c UserService) Save(user NewUser) *User {

	_, err := GetConnection().NamedExec(`INSERT INTO public.user (name) VALUES (:name)`,
		map[string]interface{}{
			"name": user.Name,
		})
	if err != nil {
		fmt.Println(err.Error())
	}

	return &User{ID: 10, Name: "test"}
}
