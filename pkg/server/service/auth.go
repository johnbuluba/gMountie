package service

import (
	"context"
	"gmountie/pkg/common"
	"gmountie/pkg/server/config"
	"gmountie/pkg/utils/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthService is a service for authentication
type AuthService interface {
	// Authorize checks if the user is authorized
	Authorize(ctx context.Context, method string) (bool, error)
}

// --------------------------- Factory ---------------------------

// NewAuthServiceFromConfig creates a new AuthService from the config
func NewAuthServiceFromConfig(cfg *config.Config) AuthService {
	switch cfg.Auth.GetType() {
	case config.AuthConfigTypeNone:
		log.Log.Warn("no authentication is enabled")
		return &NoneAuthService{}
	case config.AuthConfigTypeBasic:
		// Create users map
		authConfig := cfg.Auth.(*config.BasicAuthConfig)
		users := make(map[string]string)
		for _, user := range authConfig.Users {
			users[user.Username] = user.Password
		}
		log.Log.Info("basic authentication is enabled")
		return NewBasicAuthService(users)
	default:
		return &NoneAuthService{}
	}
}

// --------------------------- Implementations ---------------------------

// ----------- NoneAuthService -----------

// NoneAuthService is a service that does not perform any authentication
type NoneAuthService struct{}

// Authorize always returns true
func (a *NoneAuthService) Authorize(ctx context.Context, method string) (bool, error) {
	return true, nil
}

// ----------- BasicAuthService -----------

// BasicAuthService is a service that performs basic authentication
type BasicAuthService struct {
	users map[string]string
}

// NewBasicAuthService creates a new BasicAuthService
func NewBasicAuthService(users map[string]string) *BasicAuthService {
	return &BasicAuthService{users: users}
}

// Authorize checks if the user is authorized
func (a *BasicAuthService) Authorize(ctx context.Context, _ string) (bool, error) {
	// Get the user and password from the metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, status.Errorf(codes.Internal, "metadata is not provided")
	}

	user := md.Get(common.MetadataAuthBasicUsername)
	password := md.Get(common.MetadataAuthBasicPassword)
	if len(user) == 0 || len(password) == 0 {
		return false, status.Errorf(codes.Unauthenticated, "user or password is not provided")
	}

	// Check if the user exists
	if pass, ok := a.users[user[0]]; ok {
		// Check if the password is correct
		if pass == password[0] {
			return true, nil
		}
	}

	return false, status.Errorf(codes.Unauthenticated, "invalid user or password")
}
