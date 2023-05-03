package services

import "github.com/delaram-gholampoor-sagha/Deli-Authenticator/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.DBResponse, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}
