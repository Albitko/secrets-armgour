package sender

import (
	"encoding/base64"
	"fmt"
	"os"
)

type httpAPI interface {
	SendCredentials(serviceName, serviceLogin, servicePassword, meta string) error
	CreateText(title, body, meta string) error
	CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
	CreateBinary(title, b64Content, meta string) error
}

type sender struct {
	api httpAPI
}

func (s *sender) CreateBinary(title, dataPath, meta string) error {
	content, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	b64Content := base64.StdEncoding.EncodeToString(content)
	fmt.Println("Encoded Content:", b64Content)
	err = s.api.CreateBinary(title, b64Content, meta)
	return err
}

func (s *sender) CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error {
	fmt.Println(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta)
	err := s.api.CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta)
	return err
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
