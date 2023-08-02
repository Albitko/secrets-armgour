package sender

import (
	"fmt"
)

type httpAPI interface {
	SendCredentials(serviceName, serviceLogin, servicePassword, meta string) error
	CreateText(title, body, meta string) error
}

type sender struct {
	api httpAPI
}

func (s *sender) CreateText(title, body, meta string) error {
	fmt.Println(title, body, meta)
	err := s.api.CreateText(title, body, meta)
	return err
}

func (s *sender) CreateCredentials(serviceName, serviceLogin, servicePassword, meta string) error {
	fmt.Println(serviceName, serviceLogin, servicePassword, meta)
	err := s.api.SendCredentials(serviceName, serviceLogin, servicePassword, meta)
	return err
}

func New(api httpAPI) *sender {
	return &sender{
		api: api,
	}
}
