package processor

import (
	"encoding/base64"
	"fmt"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type repository interface {
	InsertCard(card entity.UserCard) error
	InsertCredentials(credentials entity.UserCredentials) error
	InsertBinary(bin entity.UserBinary, data []byte) error
	InsertText(text entity.UserText) error

	SelectUserData(data string) (interface{}, error)
}

type processor struct {
	repo repository
}

func (p processor) ListUserData(data string) (interface{}, error) {
	res, err := p.repo.SelectUserData(data)
	return res, err
}

func (p processor) BinaryCreation(binary entity.UserBinary) error {
	decodedContent, err := base64.StdEncoding.DecodeString(binary.B64Content)
	if err != nil {
		fmt.Println("Error decoding content:", err)
		return err
	}
	err = p.repo.InsertBinary(binary, decodedContent)
	return err
}

func (p processor) TextCreation(text entity.UserText) error {
	err := p.repo.InsertText(text)
	return err
}

func (p processor) CredentialsCreation(credentials entity.UserCredentials) error {
	err := p.repo.InsertCredentials(credentials)
	return err
}

func (p processor) CardCreation(card entity.UserCard) error {
	err := p.repo.InsertCard(card)
	return err
}

func New(repo repository) *processor {
	return &processor{
		repo: repo,
	}
}
