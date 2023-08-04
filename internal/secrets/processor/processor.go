package processor

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type repository interface {
	InsertCard(card entity.UserCard) error
	InsertCredentials(credentials entity.UserCredentials) error
	InsertBinary(bin entity.UserBinary, data []byte) error
	InsertText(text entity.UserText) error

	UpdateCard(index int, card entity.UserCard) error
	UpdateCredentials(index int, credentials entity.UserCredentials) error
	UpdateBinary(index int, bin entity.UserBinary, data []byte) error
	UpdateText(index int, text entity.UserText) error

	SelectUserData(data string) (interface{}, error)
	GetUserData(data, id string) (interface{}, error)
	DeleteUserData(data, id string) error

	RegisterUser(login, pass string) error
	GetUserPasswordHash(login string) (string, error)
}

type processor struct {
	repo repository
}

func (p *processor) LoginUser(auth entity.UserAuth) error {
	storedHash, err := p.repo.GetUserPasswordHash(auth.Login)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(auth.Password))
	if err == nil {
		return nil
	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		return entity.ErrInvalidCredentials
	}
	return err
}

func (p *processor) RegisterUser(auth entity.UserAuth) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = p.repo.RegisterUser(auth.Login, string(hashedPassword))
	return err
}

func (p *processor) CardEdit(index string, card entity.UserCard) error {
	intIndex, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Error parsing index:", err)
		return err
	}
	err = p.repo.UpdateCard(intIndex, card)
	return err
}

func (p *processor) BinaryEdit(index string, binary entity.UserBinary) error {
	decodedContent, err := base64.StdEncoding.DecodeString(binary.B64Content)
	if err != nil {
		fmt.Println("Error decoding content:", err)
		return err
	}
	intIndex, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Error parsing index:", err)
		return err
	}
	err = p.repo.UpdateBinary(intIndex, binary, decodedContent)
	return err
}

func (p *processor) TextEdit(index string, text entity.UserText) error {
	intIndex, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Error parsing index:", err)
		return err
	}
	err = p.repo.UpdateText(intIndex, text)
	return err
}

func (p *processor) CredentialsEdit(index string, credentials entity.UserCredentials) error {
	intIndex, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Error parsing index:", err)
		return err
	}
	err = p.repo.UpdateCredentials(intIndex, credentials)
	return err
}

func (p *processor) DeleteUserData(data, id string) error {
	err := p.repo.DeleteUserData(data, id)
	return err
}

func (p *processor) GetUserData(data, id string) (interface{}, error) {
	res, err := p.repo.GetUserData(data, id)
	if data == "binary" {
		binRes := res.(entity.UserBinary)
		binRes.B64Content = base64.StdEncoding.EncodeToString([]byte(binRes.B64Content))
		res = binRes
	}
	return res, err
}

func (p *processor) ListUserData(data string) (interface{}, error) {
	res, err := p.repo.SelectUserData(data)
	return res, err
}

func (p *processor) BinaryCreation(binary entity.UserBinary) error {
	decodedContent, err := base64.StdEncoding.DecodeString(binary.B64Content)
	if err != nil {
		fmt.Println("Error decoding content:", err)
		return err
	}
	err = p.repo.InsertBinary(binary, decodedContent)
	return err
}

func (p *processor) TextCreation(text entity.UserText) error {
	err := p.repo.InsertText(text)
	return err
}

func (p *processor) CredentialsCreation(credentials entity.UserCredentials) error {
	err := p.repo.InsertCredentials(credentials)
	return err
}

func (p *processor) CardCreation(card entity.UserCard) error {
	err := p.repo.InsertCard(card)
	return err
}

func New(repo repository) *processor {
	return &processor{
		repo: repo,
	}
}
