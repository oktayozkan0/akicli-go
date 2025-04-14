package api

import (
	"bytes"
	"encoding/json"

	"github.com/oktayozkan0/akicli-go/utils"
)

type APIResponse struct {
	StatusCode int
	RawContent []byte
	Response   LoginResponse
}

type LoginResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

func (a *API) Login(email, password string) (*APIResponse, error) {
	u := LoginPath
	reqBody, err := json.Marshal(
		map[string]interface{}{
			"email":    email,
			"password": password,
		},
	)

	if err != nil {
		return nil, err
	}
	req, err := a.client.Post(u, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	body, err := utils.ResponseAsBytes(req)
	if err != nil {
		return nil, err
	}
	var response LoginResponse
	json.Unmarshal(body, &response)
	a.client.Token = "Token " + response.Token
	apiResponse := APIResponse{
		StatusCode: req.StatusCode,
		RawContent: body,
		Response:   response,
	}
	return &apiResponse, nil
}
