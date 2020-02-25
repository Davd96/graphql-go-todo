package models

type UserResponse struct {
	ID   int32  `db:"id"`
	Name string `db:"name"`
}

type CreateUser struct {
	Input struct {
		Name string
	}
}

type UpdateUser struct {
	Input struct {
		ID   int32  `db:"id"`
		Name string `db:"name"`
	}
}

type DeleteUser struct {
	Input struct {
		ID int32
	}
}
