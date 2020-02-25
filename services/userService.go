package services

import (
	"fmt"

	"github.com/Davd96/graphql-go-todo/models"
)

type UserService struct{}

func (c UserService) GetAll() []models.UserResponse {
	var user []models.UserResponse
	err := GetConnection().Select(&user, "SELECT id,name FROM public.user")
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}

func (c UserService) GetById(id int32) []models.UserResponse {
	var user []models.UserResponse
	err := GetConnection().Select(&user, fmt.Sprintf("SELECT id,name FROM public.user WHERE id = %v", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}

func (c UserService) Save(user models.CreateUser) *models.UserResponse {
	var insertedUser models.UserResponse
	query := fmt.Sprintf("INSERT INTO public.user (name) VALUES ('%s') returning id, name", user.Input.Name)

	err := GetConnection().QueryRow(query).Scan(&insertedUser.ID, &insertedUser.Name)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &insertedUser
}

func (c UserService) Update(user models.UpdateUser) *models.UserResponse {
	var updatedUser models.UserResponse
	query := fmt.Sprintf("UPDATE public.user SET %s WHERE id = %v returning id, name", makeUpdateQuery(user.Input), user.Input.ID)
	err := GetConnection().QueryRow(query).Scan(&updatedUser.ID, &updatedUser.Name)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &updatedUser
}

func (c UserService) Delete(user models.DeleteUser) *models.UserResponse {
	var deletedUser models.UserResponse

	query := fmt.Sprintf("DELETE FROM public.user WHERE id=%v returning id, name", user.Input.ID)

	err := GetConnection().QueryRow(query).Scan(&deletedUser.ID, &deletedUser.Name)

	if err != nil {
		fmt.Println(err.Error())
	}

	return &deletedUser
}
