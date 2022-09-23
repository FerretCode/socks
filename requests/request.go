package request

import (
	"io"
	"net/http"
)

type Config struct {
	Token string `env:"SOCKS_TOKEN"`
	Domain string `env:"SOCKS_DOMAIN"`
}

type ResponseParser struct {
	Response http.Response
}

func (rp ResponseParser) ParseRequest(http.Response) ([]byte, error) {
	body, err := io.ReadAll(rp.Response.Body)

	if err != nil {
		return nil, err 
	}

	return body, nil
}
