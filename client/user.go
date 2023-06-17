package client

import (
	"bytes"
	"encoding/json"

	// "errors"
	"github.com/Krisna20046/config"
	// "github.com/Krisna20046/model"
	// "io/ioutil"
	"net/http"
)

type UserClient interface {
	Login(username, password string) (respCode int, err error)
	Register(nama, email, username, password, no_hp, jenis_kelamin, alamat string) (respCode int, err error)
}

type userClient struct {
}

func NewUserClient() *userClient {
	return &userClient{}
}

func (u *userClient) Login(username, password string) (respCode int, err error) {
	datajson := map[string]string{
		"username": username,
		"password": password,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/user/login"), bytes.NewBuffer(data))
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if err != nil {
		return -1, err
	} else {
		return resp.StatusCode, nil
	}
}

func (u *userClient) Register(nama, email, username, password, no_hp, jenis_kelamin, alamat string) (respCode int, err error) {
	datajson := map[string]string{
		"nama":          nama,
		"email":         email,
		"username":      username,
		"password":      password,
		"no_hp":         no_hp,
		"jenis_kelamin": jenis_kelamin,
		"alamat":        alamat,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/user/register"), bytes.NewBuffer(data))
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if err != nil {
		return -1, err
	} else {
		return resp.StatusCode, nil
	}
}
