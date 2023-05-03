package gapi

import (
	"github.com/delaram-gholampoor-sagha/Deli-Authenticator/config"
	"github.com/delaram-gholampoor-sagha/Deli-Authenticator/pb"
	"github.com/delaram-gholampoor-sagha/Deli-Authenticator/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	config         config.Config
	userService    services.UserService
	userCollection *mongo.Collection
}

func NewGrpcUserServer(config config.Config, userService services.UserService, userCollection *mongo.Collection) (*UserServer, error) {
	userServer := &UserServer{
		config:         config,
		userService:    userService,
		userCollection: userCollection,
	}

	return userServer, nil
}
