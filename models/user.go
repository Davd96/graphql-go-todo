package models

import (
	"fmt"
)

type UserResponse struct {
	ID   int32  `db:"id"`
	Name string `db:"name"`
}

type User struct {
	ID   int32   `db:"id"`
	Name *string `db:"name"`
}

type UserService struct{}

func (c UserService) GetAll() []UserResponse {
	var user []UserResponse
	err := GetConnection().Select(&user, "SELECT id,name FROM public.user")
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}

func (c UserService) GetById(id int32) []UserResponse {
	var user []UserResponse
	err := GetConnection().Select(&user, fmt.Sprintf("SELECT id,name FROM public.user WHERE id = %v", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}

func (c UserService) Save(user User) *UserResponse {
	var insertedUser UserResponse
	query := fmt.Sprintf("INSERT INTO public.user (name) VALUES ('%s') returning id, name", *user.Name)

	err := GetConnection().QueryRow(query).Scan(&insertedUser.ID, &insertedUser.Name)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &insertedUser
}

func (c UserService) Update(user User) *UserResponse {
	var updatedUser UserResponse
	query := fmt.Sprintf("UPDATE public.user SET %s WHERE id = %v returning id, name", makeUpdateQuery(user), user.ID)
	err := GetConnection().QueryRow(query).Scan(&updatedUser.ID, &updatedUser.Name)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &updatedUser
}

func (c UserService) Delete(user User) {

}
