// Package httpapi - package with clients for
// armGour server
package httpapi

import (
	"strconv"

	"github.com/go-resty/resty/v2"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type httpAPI struct {
	armgourURL string
	client     *resty.Client
}

// LoginUser - method that authorize user
func (a *httpAPI) LoginUser(login, password string) error {
	userCredentials := entity.UserAuth{
		Login:    login,
		Password: password,
	}
	resp, err := a.client.R().SetBody(userCredentials).Post(a.armgourURL + "/v1/user/login")
	if resp.StatusCode() == 401 {
		return entity.ErrInvalidCredentials
	}
	return err
}

// RegisterUser - method that register user in service
func (a *httpAPI) RegisterUser(login, password string) error {
	userCredentials := entity.UserAuth{
		Login:    login,
		Password: password,
	}
	resp, err := a.client.R().SetBody(userCredentials).Post(a.armgourURL + "/v1/user/register")
	if resp.StatusCode() == 409 {
		return entity.ErrLoginAlreadyInUse
	}
	return err
}

// EditCredentials - method that provide credentials entity modification
func (a *httpAPI) EditCredentials(index int, serviceName, serviceLogin, servicePassword, meta string) error {
	userCredentials := entity.UserCredentials{
		ServiceLogin:    serviceLogin,
		ServiceName:     serviceName,
		ServicePassword: servicePassword,
		Meta:            meta,
	}
	_, err := a.client.R().SetBody(userCredentials).Put(a.armgourURL + "/v1/secrets/credentials/" + strconv.Itoa(index))
	return err
}

// EditText - - method that provide text entity modification
func (a *httpAPI) EditText(index int, title, body, meta string) error {
	text := entity.UserText{
		Title: title,
		Body:  body,
		Meta:  meta,
	}
	_, err := a.client.R().SetBody(text).Put(a.armgourURL + "/v1/secrets/text/" + strconv.Itoa(index))
	return err
}

// EditCard - method that provide cards entity modification
func (a *httpAPI) EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error {
	card := entity.UserCard{
		CardHolder:         cardHolder,
		CardNumber:         cardNumber,
		CardValidityPeriod: cardValidityPeriod,
		CvcCode:            cvcCode,
		Meta:               meta,
	}
	_, err := a.client.R().SetBody(card).Put(a.armgourURL + "/v1/secrets/card/" + strconv.Itoa(index))
	return err
}

// EditBinary - method that provide binary entity modification
func (a *httpAPI) EditBinary(index int, title, b64Content, meta string) error {
	binary := entity.UserBinary{
		Title:      title,
		B64Content: b64Content,
		Meta:       meta,
	}
	_, err := a.client.R().SetBody(binary).Put(a.armgourURL + "/v1/secrets/binary/" + strconv.Itoa(index))
	return err
}

// DeleteUserSecrets - method that delete users secret
func (a *httpAPI) DeleteUserSecrets(secretType string, idx int) error {
	_, err := a.client.R().Delete(a.armgourURL + "/v1/secrets/" + secretType + "/" + strconv.Itoa(idx))
	return err
}

// GetSecret - method that return user secret
func (a *httpAPI) GetSecret(secretType, user string, idx int) (string, error) {
	resp, err := a.client.R().Get(
		a.armgourURL + "/v1/secrets/get/" + secretType + "/" + strconv.Itoa(idx) + "/" + user)
	return resp.String(), err
}

// ListSecrets - method that return users secrets list
func (a *httpAPI) ListSecrets(data, user string) (string, error) {
	resp, err := a.client.R().Get(a.armgourURL + "/v1/secrets/list/" + data + "/" + user)
	return resp.String(), err
}

// CreateBinary - method that save binary in service
func (a *httpAPI) CreateBinary(title, b64Content, meta, user string) error {
	binary := entity.UserBinary{
		Title:      title,
		B64Content: b64Content,
		Meta:       meta,
	}
	_, err := a.client.R().SetBody(binary).Post(a.armgourURL + "/v1/secrets/binary/" + user)
	return err
}

// SendCredentials - save service credentials
func (a *httpAPI) SendCredentials(serviceName, serviceLogin, servicePassword, meta, user string) error {
	userCredentials := entity.UserCredentials{
		ServiceLogin:    serviceLogin,
		ServiceName:     serviceName,
		ServicePassword: servicePassword,
		Meta:            meta,
	}
	_, err := a.client.R().SetBody(userCredentials).Post(a.armgourURL + "/v1/secrets/credentials/" + user)
	return err
}

// CreateText - save text in service
func (a *httpAPI) CreateText(title, body, meta, user string) error {
	text := entity.UserText{
		Title: title,
		Body:  body,
		Meta:  meta,
	}
	_, err := a.client.R().SetBody(text).Post(a.armgourURL + "/v1/secrets/text/" + user)
	return err
}

// CreateCard - save card data in service
func (a *httpAPI) CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta, user string) error {
	card := entity.UserCard{
		CardHolder:         cardHolder,
		CardNumber:         cardNumber,
		CardValidityPeriod: cardValidityPeriod,
		CvcCode:            cvcCode,
		Meta:               meta,
	}
	_, err := a.client.R().SetBody(card).Post(a.armgourURL + "/v1/secrets/card/" + user)
	return err
}

func New(serverURL string) *httpAPI {
	r := resty.New()
	return &httpAPI{
		armgourURL: serverURL,
		client:     r,
	}
}
