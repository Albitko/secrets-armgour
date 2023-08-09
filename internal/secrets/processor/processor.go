// Package processor - main use case of server side
// process all information with client interconnection
package processor

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

//go:generate mockery --name repository
type repository interface {
	InsertCard(ctx context.Context, card entity.UserCard, user string) error
	InsertCredentials(ctx context.Context, credentials entity.UserCredentials, user string) error
	InsertBinary(ctx context.Context, bin entity.UserBinary, data []byte, user string) error
	InsertText(ctx context.Context, text entity.UserText, user string) error

	UpdateCard(ctx context.Context, index int, card entity.UserCard) error
	UpdateCredentials(ctx context.Context, index int, credentials entity.UserCredentials) error
	UpdateBinary(ctx context.Context, index int, bin entity.UserBinary, data []byte) error
	UpdateText(ctx context.Context, index int, text entity.UserText) error

	SelectUserData(ctx context.Context, data, string string) (interface{}, error)
	GetUserData(ctx context.Context, data, id, user string) (interface{}, error)
	DeleteUserData(ctx context.Context, data, id string) error

	RegisterUser(ctx context.Context, login, pass string) error
	GetUserPasswordHash(ctx context.Context, login string) (string, error)
}

type processor struct {
	repo repository
}

// LoginUser - process login requests to repository
func (p *processor) LoginUser(ctx context.Context, auth entity.UserAuth) error {
	storedHash, err := p.repo.GetUserPasswordHash(ctx, auth.Login)
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

// RegisterUser - save user credentials as service user in repository
func (p *processor) RegisterUser(ctx context.Context, auth entity.UserAuth) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = p.repo.RegisterUser(ctx, auth.Login, string(hashedPassword))
	return err
}

// CardEdit - edit card data, pass request to db
func (p *processor) CardEdit(ctx context.Context, index string, card entity.UserCard) error {
	intIndex, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Error parsing index:", err)
		return err
	}
	err = p.repo.UpdateCard(ctx, intIndex, card)
	return err
}

// BinaryEdit - edit binary data, pass request to db
func (p *processor) BinaryEdit(ctx context.Context, index string, binary entity.UserBinary) error {
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
	err = p.repo.UpdateBinary(ctx, intIndex, binary, decodedContent)
	return err
}

// TextEdit - edit text data, pass request to db
func (p *processor) TextEdit(ctx context.Context, index string, text entity.UserText) error {
	intIndex, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Error parsing index:", err)
		return err
	}
	err = p.repo.UpdateText(ctx, intIndex, text)
	return err
}

// CredentialsEdit - edit credentials data, pass request to db
func (p *processor) CredentialsEdit(ctx context.Context, index string, credentials entity.UserCredentials) error {
	intIndex, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Error parsing index:", err)
		return err
	}
	err = p.repo.UpdateCredentials(ctx, intIndex, credentials)
	return err
}

// DeleteUserData - delete all user data, pass request to db
func (p *processor) DeleteUserData(ctx context.Context, data, id string) error {
	err := p.repo.DeleteUserData(ctx, data, id)
	return err
}

// GetUserData - find all user data by index and user, pass request to db
func (p *processor) GetUserData(ctx context.Context, data, id, user string) (interface{}, error) {
	res, err := p.repo.GetUserData(ctx, data, id, user)
	if data == "binary" {
		binRes := res.(entity.UserBinary)
		binRes.B64Content = base64.StdEncoding.EncodeToString([]byte(binRes.B64Content))
		res = binRes
	}
	return res, err
}

// ListUserData - find all user data by user, pass request to db
func (p *processor) ListUserData(ctx context.Context, data, user string) (interface{}, error) {
	res, err := p.repo.SelectUserData(ctx, data, user)
	return res, err
}

// BinaryCreation - save user binary data, pass request to db
func (p *processor) BinaryCreation(ctx context.Context, binary entity.UserBinary, user string) error {
	decodedContent, err := base64.StdEncoding.DecodeString(binary.B64Content)
	if err != nil {
		fmt.Println("Error decoding content:", err)
		return err
	}
	err = p.repo.InsertBinary(ctx, binary, decodedContent, user)
	return err
}

// TextCreation - save user text data, pass request to db
func (p *processor) TextCreation(ctx context.Context, text entity.UserText, user string) error {
	err := p.repo.InsertText(ctx, text, user)
	return err
}

// CredentialsCreation - save user credentials, pass request to db
func (p *processor) CredentialsCreation(ctx context.Context, credentials entity.UserCredentials, user string) error {
	err := p.repo.InsertCredentials(ctx, credentials, user)
	return err
}

// CardCreation - save user cards data, pass request to db
func (p *processor) CardCreation(ctx context.Context, card entity.UserCard, user string) error {
	err := p.repo.InsertCard(ctx, card, user)
	return err
}

// New - create server side use case
func New(repo repository) *processor {
	return &processor{
		repo: repo,
	}
}
