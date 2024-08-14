package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (c Client) MakeRequest(TCKimlik *TCKimlik) bool {
	envelope := Envelope{
		XSI:    c.Config.XSI,
		XSD:    c.Config.XSD,
		Soap12: c.Config.Soap12,
		Body: TCKimlikNoDogrula{
			TCKimlik: *TCKimlik,
		},
	}

	xmlPayload, err := xml.Marshal(envelope)
	if err != nil {
		msg := fmt.Sprintf("xmlPayload error : %s", err.Error())
		log.Println(msg)
		return false
	}

	payload := bytes.NewReader(xmlPayload)
	req, err := http.NewRequest(c.Config.HTTPMethod, c.Config.BaseURL, payload)
	if err != nil {
		msg := fmt.Sprintf("http request error : %s", err.Error())
		log.Println(msg)
		return false
	}
	req.Header.Add("Content-Type", c.Config.ContentType)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		msg := fmt.Sprintf("http client error : %s", err.Error())
		log.Println(msg)
		return false
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		msg := fmt.Sprintf("http response body read error : %s", err.Error())
		log.Println(msg)
		return false
	}

	response := new(TCKimlikNoDogrulaResponse)
	if err := xml.Unmarshal([]byte(body), &response); err != nil {
		msg := fmt.Sprintf("http response body unmarshal error : %s", err.Error())
		log.Println(msg)
		return false
	}

	return response.Body.Result.Value
}

func (c Client) GenerateRandomString(Length int) string {
	randomString := make([]byte, Length)
	for i := range randomString {
		randomString[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(randomString)
}
