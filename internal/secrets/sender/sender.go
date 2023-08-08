package sender

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Albitko/secrets-armgour/internal/entity"
	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type httpAPI interface {
	SendCredentials(serviceName, serviceLogin, servicePassword, meta, user string) error
	CreateText(title, body, meta, user string) error
	CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta, user string) error
	CreateBinary(title, b64Content, meta, user string) error
	ListSecrets(data, user string) (string, error)
	GetSecret(secretType, user string, idx int) (string, error)
	DeleteUserSecrets(secretType string, idx int) error
	EditCredentials(index int, serviceName, serviceLogin, servicePassword, meta string) error
	EditText(index int, title, body, meta string) error
	EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
	EditBinary(index int, title, b64Content, meta string) error

	RegisterUser(login, password string) error
	LoginUser(login, password string) error
}

type sender struct {
	api httpAPI
}

// LoginUser - use http client for user login
func (s *sender) LoginUser(login, password string) error {
	err := s.api.LoginUser(login, password)
	return err
}

// RegisterUser - use http client for user register
func (s *sender) RegisterUser(login, password string) error {
	err := s.api.RegisterUser(login, password)
	return err
}

// EditCredentials - use http client for change user credentials
func (s *sender) EditCredentials(index int, serviceName, serviceLogin, servicePassword, meta string) error {
	err := s.api.EditCredentials(index, serviceName, serviceLogin, servicePassword, meta)
	return err
}

// EditBinary - use http client for change user binary data
func (s *sender) EditBinary(index int, title, dataPath, meta string) error {
	var key, encTitle, encContent, encMeta string
	content, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	b64Content := base64.StdEncoding.EncodeToString(content)

	key, _, err = encrypt.GetCliSecrets()
	if err != nil {
		fmt.Println(err)
		return err
	}
	encTitle, err = encrypt.EncryptMessage([]byte(key), title)
	if err != nil {
		fmt.Println(err)
		return err
	}
	encContent, err = encrypt.EncryptMessage([]byte(key), b64Content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	encMeta, err = encrypt.EncryptMessage([]byte(key), meta)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = s.api.EditBinary(index, encTitle, encContent, encMeta)
	return err
}

// EditCard - use http client for change user card data
func (s *sender) EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error {
	err := s.api.EditCard(index, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta)
	return err
}

// EditText - use http client for change user text data
func (s *sender) EditText(index int, title, body, meta string) error {
	err := s.api.EditText(index, title, body, meta)
	return err
}

// DeleteUserSecrets - use http client for delete user data
func (s *sender) DeleteUserSecrets(secretType string, idx int) error {
	return s.api.DeleteUserSecrets(secretType, idx)
}

// GetUserSecrets - use http client for get user data
func (s *sender) GetUserSecrets(secretType, user string, idx int) (interface{}, error) {
	resp, err := s.api.GetSecret(secretType, user, idx)
	var res interface{}
	switch secretType {
	case entity.Credentials:
		var cred entity.UserCredentials
		err = json.Unmarshal([]byte(resp), &cred)
		if err != nil {
			return cred, err
		}
		res = cred
	case entity.Binary:
		var bin entity.UserBinary
		err = json.Unmarshal([]byte(resp), &bin)
		if err != nil {
			return bin, err
		}
		res = bin
	case entity.Text:
		var text entity.UserText
		err = json.Unmarshal([]byte(resp), &text)
		if err != nil {
			return text, err
		}
		res = text
	case entity.Card:
		var card entity.UserCard
		err = json.Unmarshal([]byte(resp), &card)
		if err != nil {
			return card, err
		}
		res = card
	default:
		return res, fmt.Errorf("unsupported data type")
	}
	return res, err
}

// ListUserSecrets - use http client for list all user data
func (s *sender) ListUserSecrets(data, user string) (interface{}, error) {
	resp, err := s.api.ListSecrets(data, user)
	var res interface{}
	switch data {
	case entity.Credentials:
		var cred []entity.CutCredentials
		err = json.Unmarshal([]byte(resp), &cred)
		if err != nil {
			return cred, err
		}
		res = cred
	case entity.Binary:
		var bin []entity.CutBinary
		err = json.Unmarshal([]byte(resp), &bin)
		if err != nil {
			return bin, err
		}
		res = bin
	case entity.Text:
		var text []entity.CutText
		err = json.Unmarshal([]byte(resp), &text)
		if err != nil {
			return text, err
		}
		res = text
	case entity.Card:
		var card []entity.CutCard
		err = json.Unmarshal([]byte(resp), &card)
		if err != nil {
			return card, err
		}
		res = card
	default:
		return res, fmt.Errorf("unsupported data type")
	}
	return res, err
}

// CreateBinary - use http client for create user binary data
func (s *sender) CreateBinary(title, dataPath, meta string) error {
	content, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	b64Content := base64.StdEncoding.EncodeToString(content)
	key, user, err := encrypt.GetCliSecrets()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	encTitle, err := encrypt.EncryptMessage([]byte(key), title)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	encContent, err := encrypt.EncryptMessage([]byte(key), b64Content)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	encMeta, err := encrypt.EncryptMessage([]byte(key), meta)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	err = s.api.CreateBinary(encTitle, encContent, encMeta, user)
	return err
}

// CreateCard - use http client for create user card data
func (s *sender) CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta, user string) error {
	err := s.api.CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta, user)
	return err
}

// CreateText - use http client for create user text data
func (s *sender) CreateText(title, body, meta, user string) error {
	err := s.api.CreateText(title, body, meta, user)
	return err
}

// CreateCredentials - use http client for create user credentials data
func (s *sender) CreateCredentials(serviceName, serviceLogin, servicePassword, meta, user string) error {
	err := s.api.SendCredentials(serviceName, serviceLogin, servicePassword, meta, user)
	return err
}

// New - cli use case
func New(api httpAPI) *sender {
	return &sender{
		api: api,
	}
}
