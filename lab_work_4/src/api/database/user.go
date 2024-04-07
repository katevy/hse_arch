package database

import "api/models"

type UserDB interface {
	CreateUser(user *models.User) error
}
