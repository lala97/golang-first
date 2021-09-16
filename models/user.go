package model

import (
	"go_crud/database"
	entity "go_crud/entities"
	"html"
	"strings"
)

type User entity.User

func (user *User) FindAll() (*[]User, error) {
	var users []User
	err := database.Connector.Joins("JOIN roles ON roles.id=users.role_id").Preload("roles").Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (user *User) FindById(id int32) (*User, error) {
	err := database.Connector.Model(User{}).First(&user, id).Error
	if err != nil {
		return &User{}, err
	}
	return user, err
}

func (user *User) Delete(id int32) error {
	err := database.Connector.Model(&User{}).Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Create(newUser *User) (*User, error) {
	newUser.Name = html.EscapeString(strings.TrimSpace(newUser.Name))
	newUser.Surname = html.EscapeString(strings.TrimSpace(newUser.Surname))

	err := database.Connector.Model(newUser).Create(newUser).Error
	if err != nil {
		return &User{}, err
	}
	return user, err
}

func (user *User) Update(id int32) (*User, error) {
	result := database.Connector.Model(&user).Where("id = ?", id).UpdateColumns(
		map[string]interface{}{
			"name":    user.Name,
			"surname": user.Surname,
		},
	)

	if result.Error != nil {
		return &User{}, result.Error
	}

	//return updated user
	err := database.Connector.Model(&user).First(&user, id).Error
	if err != nil {
		return &User{}, err
	}
	return user, err
}
