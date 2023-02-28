package http

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"log"
)

func Login(phoneNumber string) *LoginResponse {
	client := req.C()
	url := fmt.Sprintf("http://localhost:8081/userlogin?phone_number=%s", phoneNumber)
	res, err := client.R().Post(url)
	if err != nil {
		log.Println("Error during Login Call:")
		log.Println(err)
		return nil
	}
	var loginResponse LoginResponse
	err = json.Unmarshal([]byte(res.String()), &loginResponse)
	if err != nil || !loginResponse.Ok {
		tgErr := unmarshalErrorResponse(res.String())
		log.Println("Error during Login Call:")
		log.Println(tgErr)
		return nil
	}
	return &loginResponse
}

func ConfirmLogin(userToken string, otp string) *ConfirmLoginResult {
	client := req.C()
	url := fmt.Sprintf("http://localhost:8081/user%s/authCode?code=%s", userToken, otp)
	res, err := client.R().Post(url)
	if err != nil {
		log.Println("Error during ConfirmLogin Call:")
		log.Println(err)
		return nil
	}
	var confirmLoginResponse ConfirmLoginResult
	err = json.Unmarshal([]byte(res.String()), &confirmLoginResponse)
	if err != nil || !confirmLoginResponse.Ok {
		log.Println("Error during ConfirmLogin response unmarshaling or call ko")
		return nil
	}
	return &confirmLoginResponse
}
