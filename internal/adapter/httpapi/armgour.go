package httpapi

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type httpAPI struct {
	armgourURL string
	client     *resty.Client
}

func (a *httpAPI) CreateBinary(title, b64Content, meta string) error {
	binary := entity.UserBinary{
		Title:      title,
		B64Content: b64Content,
		Meta:       meta,
	}
	resp, err := a.client.R().SetBody(binary).Post(a.armgourURL + "/v1/secrets/binary/create")
	fmt.Println(resp.String())
	return err
}

func (a *httpAPI) SendCredentials(serviceName, serviceLogin, servicePassword, meta string) error {
	userCredentials := entity.UserCredentials{
		ServiceLogin:    serviceLogin,
		ServiceName:     serviceName,
		ServicePassword: servicePassword,
		Meta:            meta,
	}
	resp, err := a.client.R().SetBody(userCredentials).Post(a.armgourURL + "/v1/secrets/credentials/create")
	fmt.Println(resp.String())
	return err
}

func (a *httpAPI) CreateText(title, body, meta string) error {
	text := entity.UserText{
		Title: title,
		Body:  body,
		Meta:  meta,
	}
	resp, err := a.client.R().SetBody(text).Post(a.armgourURL + "/v1/secrets/text/create")
	fmt.Println(resp.String())
	return err
}

func (a *httpAPI) CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error {
	card := entity.UserCard{
		CardHolder:         cardHolder,
		CardNumber:         cardNumber,
		CardValidityPeriod: cardValidityPeriod,
		CvcCode:            cvcCode,
		Meta:               meta,
	}
	resp, err := a.client.R().SetBody(card).Post(a.armgourURL + "/v1/secrets/card/create")
	fmt.Println(resp.String())
	return err
}

func New(serverURL string) *httpAPI {
	r := resty.New()
	return &httpAPI{
		armgourURL: serverURL,
		client:     r,
	}
}
