package netease

import (
	"github.com/bingohuang/163yun-go-sdk/cloudcomb"
)

type Config struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (c *Config) Client() (*cloudcomb.CloudComb, error) {
	client := cloudcomb.NewCC(c.AppKey, c.AppSecret)

	token, err := client.UserToken()
	if err != nil {
		return nil, err
	}
	client.Token = token

	return client, nil
}
