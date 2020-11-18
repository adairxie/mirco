package service

import (
	"context"
	"errors"
	"security/model"
)

var (
	ErrUserNotExist = errors.New("username is not exist")
	ErrPassword     = errors.New("invalid password")
)

// Service define a service interface
type UserDetailService interface {
	// Get UserDetails By username
	GetUserDetailByUsername(ctx context.Context, username, password string) (*model.UserDetails, error)
}

type InMemoryUserDetailService struct {
	userDetailsDict map[string]*model.UserDetails
}

func (service *InMemoryUserDetailService) GetUserDetailByUsername(ctx context.Context, username, password string) (*model.UserDetails, error) {

	// 根据username 获取用户信息
	userDetails, ok := service.userDetailsDict[username]
	if ok {
		if userDetails.Password == password {
			return userDetails, nil
		} else {
			return nil, ErrPassword
		}
	} else {
		return nil, ErrUserNotExist
	}
}

func NewInMemoryUserDetailsService(userDetailsList []*model.UserDetails) *InMemoryUserDetailService {
	userDetailsDict := make(map[string]*model.UserDetails)

	if userDetailsList != nil {
		for _, value := range userDetailsList {
			userDetailsDict[value.Username] = value
		}
	}

	return &InMemoryUserDetailService{
		userDetailsDict: userDetailsDict,
	}
}
