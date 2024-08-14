package client

import (
	"config"
	"log"
	"net/http"
)

type UseCase interface {
	MakeRequest(TCKimlik *TCKimlik) bool
	GenerateRandomString(Length int) string
}

type Client struct {
	HttpClient *http.Client
	Config     *config.ServiceConfig
}

func NewClient() UseCase {
	serviceConfig, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	c := &http.Client{}
	return &Client{
		HttpClient: c,
		Config:     serviceConfig,
	}
}
