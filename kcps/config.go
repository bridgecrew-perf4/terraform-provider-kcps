package kcps

import "github.com/uesyn/gokcps"

// configuration structure for KCPS API Client
type Config struct {
	APIURL    string
	APIKey    string
	SecretKey string
	VerifySSL bool
}

func (c *Config) Client() *gokcps.KCPSClient {
	return gokcps.NewAsyncClient(c.APIURL, c.APIKey, c.SecretKey, c.VerifySSL)
}
