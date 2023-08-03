package sender

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type httpAPI interface {
	SendCredentials(serviceName, serviceLogin, servicePassword, meta string) error
	CreateText(title, body, meta string) error
	CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
	CreateBinary(title, b64Content, meta string) error
	ListSecrets(data string) (string, error)
	GetSecret(secretType string, idx int) (string, error)
	DeleteUserSecrets(secretType string, idx int) error
	EditCredentials(index int, serviceName, serviceLogin, servicePassword, meta string) error
	EditText(index int, title, body, meta string) error
	EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
	EditBinary(index int, title, b64Content, meta string) error
}

type sender struct {
	api httpAPI
}

func (s *sender) EditCredentials(index int, serviceName, serviceLogin, servicePassword, meta string) error {
	err := s.api.EditCredentials(index, serviceName, serviceLogin, servicePassword, meta)
	return err
}

func (s *sender) EditBinary(index int, title, dataPath, meta string) error {
	content, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	b64Content := base64.StdEncoding.EncodeToString(content)
	err = s.api.EditBinary(index, title, b64Content, meta)
	return err
}

func (s *sender) EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error {
	err := s.api.EditCard(index, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta)
	return err
}

func (s *sender) EditText(index int, title, body, meta string) error {
	err := s.api.EditText(index, title, body, meta)
	return err
}

func (s *sender) DeleteUserSecrets(secretType string, idx int) error {
	return s.api.DeleteUserSecrets(secretType, idx)
}

func (s *sender) GetUserSecrets(secretType string, idx int) (interface{}, error) {
	resp, err := s.api.GetSecret(secretType, idx)
	var res interface{}
	switch secretType {
	case "credentials":
		var cred entity.UserCredentials
		err = json.Unmarshal([]byte(resp), &cred)
		if err != nil {
			return cred, err
		}
		res = cred
	case "binary":
		var bin entity.UserBinary
		err = json.Unmarshal([]byte(resp), &bin)
		if err != nil {
			return bin, err
		}
		res = bin
	case "text":
		var text entity.UserText
		err = json.Unmarshal([]byte(resp), &text)
		if err != nil {
			return text, err
		}
		res = text
	case "card":
		var card entity.UserCard
		err = json.Unmarshal([]byte(resp), &card)
		if err != nil {
			return card, err
		}
		res = card
	}
	return res, err
}

func (s *sender) ListUserSecrets(data string) (interface{}, error) {
	resp, err := s.api.ListSecrets(data)
	var res interface{}
	switch data {
	case "credentials":
		var cred []entity.CutCredentials
		err = json.Unmarshal([]byte(resp), &cred)
		if err != nil {
			return cred, err
		}
		res = cred
	case "binary":
		var bin []entity.CutBinary
		err = json.Unmarshal([]byte(resp), &bin)
		if err != nil {
			return bin, err
		}
		res = bin
	case "text":
		var text []entity.CutText
		err = json.Unmarshal([]byte(resp), &text)
		if err != nil {
			return text, err
		}
		res = text
	case "card":
		var card []entity.CutCard
		err = json.Unmarshal([]byte(resp), &card)
		if err != nil {
			return card, err
		}
		res = card
	}
	return res, err
}

func (s *sender) CreateBinary(title, dataPath, meta string) error {
	content, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	b64Content := base64.StdEncoding.EncodeToString(content)
	err = s.api.CreateBinary(title, b64Content, meta)
	return err
}

func (s *sender) CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error {
	err := s.api.CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta)
	return err
}

func (s *sender) CreateText(title, body, meta string) error {
	err := s.api.CreateText(title, body, meta)
	return err
}

func (s *sender) CreateCredentials(serviceName, serviceLogin, servicePassword, meta string) error {
	err := s.api.SendCredentials(serviceName, serviceLogin, servicePassword, meta)
	return err
}

func New(api httpAPI) *sender {
	return &sender{
		api: api,
	}
}
