package service

import (
	"context"
	"errors"
	"security/model"
)

var (
	ErrClientNotExist = errors.New("clientId is not exist")
	ErrClientSecret   = errors.New("invalid clientSecret")
)

// Service define a service interface
type ClientDetailService interface {
	GetClientDetailByClientId(ctx context.Context, clientId string, clientSecret string) (*model.ClientDetails, error)
}

type InMemoryClientDetailService struct {
	clientDetailsDict map[string]*model.ClientDetails
}

func NewInMemoryClientDetailService(clientDetailsList []*model.ClientDetails) *InMemoryClientDetailService {
	clientDetailsDict := make(map[string]*model.ClientDetails)

	if clientDetailsList != nil {
		for _, value := range clientDetailsList {
			clientDetailsDict[value.ClientId] = value
		}
	}

	return &InMemoryClientDetailService{
		clientDetailsDict: clientDetailsDict,
	}
}

func (service *InMemoryClientDetailService) GetClientDetailByClientId(ctx context.Context, clientId string, clientSecret string) (*model.ClientDetails, error) {

	// 根据 clientId 获取 clientDetails
	clientDetails, ok := service.clientDetailsDict[clientId]
	if ok {
		// 比较 clientSecret 是否正确
		if clientDetails.ClientSecret == clientSecret {
			return clientDetails, nil
		} else {
			return nil, ErrClientSecret
		}
	} else {
		return nil, ErrClientNotExist
	}
}
