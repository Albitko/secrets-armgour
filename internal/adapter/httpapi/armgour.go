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

func (a *httpAPI) SendCredentials(serviceName, serviceLogin, servicePassword, meta string) error {
	userCredentials := entity.UserCredentials{
		ServiceLogin:    serviceLogin,
		ServiceName:     serviceName,
		ServicePassword: servicePassword,
		Meta:            meta,
	}
	resp, _ := a.client.R().SetBody(userCredentials).Post(a.armgourURL + "/v1/secrets/credentials/create")
	fmt.Println(resp.String())
	return nil
}

func New(serverURL string) *httpAPI {
	r := resty.New()
	return &httpAPI{
		armgourURL: serverURL,
		client:     r,
	}
}
