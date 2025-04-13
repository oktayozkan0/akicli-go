package api

import (
	"bytes"
	"encoding/json"

	"github.com/oktayozkan0/akicli-go/utils"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func (a *API) Login(email, password string) error {
	u := LoginPath
	reqBody, err := json.Marshal(
		map[string]interface{}{
			"email":    email,
			"password": password,
		},
	)

	if err != nil {
		return err
	}
	req, err := a.client.Post(u, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	body, err := utils.ResponseAsBytes(req)
	if err != nil {
		return err
	}
	var response LoginResponse
	json.Unmarshal(body, &response)
	a.client.Token = "Token " + response.Token
	return nil
}
