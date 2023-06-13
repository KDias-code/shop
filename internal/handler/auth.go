package handler

import (
	"encoding/json"
	"fmt"
	"github.com/KDias-code/internal/model"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	username = "dkubay"
	password = "diasKUBAY02_"
	sender   = "Halyk Shop"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input model.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"id": id,
	}

	writeJSONResponse(w, http.StatusOK, response)
}

type signInInput struct {
	Number   string `json:"number" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Number, input.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"token": token,
	}

	writeJSONResponse(w, http.StatusOK, response)
}

func (h *Handler) Otps(w http.ResponseWriter, r *http.Request) {
	userId, err := h.getUserId(r)
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	var input model.UpdateUserInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Authorization.Update(userId, input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := statusResponse{Status: "success"}
	writeJSONResponse(w, http.StatusOK, response)
}

func (h *Handler) Sms(w http.ResponseWriter, r *http.Request) {
	var input model.User
	recipient := input.Number

	rand.NewSource(time.Now().UnixNano())
	verificationCode := rand.Intn(9000) + 1000
	h.services.Authorization.RndSave(verificationCode, recipient, input)

	message := fmt.Sprintf("Your verification code is: %s", verificationCode)

	url1 := "https://smsc.kz/sys/send.php"

	params := url.Values{}
	params.Set("login", username)
	params.Set("psw", password)
	params.Set("phones", recipient)
	params.Set("mes", message)
	params.Set("sender", sender)
	params.Set("cost", "1")

	resp, err := http.Get(url1 + "?" + params.Encode())
	if err != nil {
		fmt.Println("Failed to connect to the API(SMS).")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the API response.")
		return
	}
	if h.services.Authorization.SmsCheck(verificationCode, recipient, input) != nil {
		result := string(body)
		if result == "OK" {
			fmt.Println("SMS sent successfully!")
		} else {
			fmt.Printf("SMS sending failed. Error: %s\n", result)
		}
	}

}
