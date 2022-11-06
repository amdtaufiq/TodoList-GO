package helper

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// Bind request error checking
// Return array of string that contain error messages
func FormatValidationError(bindError error) []string {
	errorMessages := []string{}
	var validatorErr validator.ValidationErrors

	if errors.As(bindError, &validatorErr) {
		for _, e := range bindError.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
	} else {
		errorMessages = append(errorMessages, bindError.Error())
	}

	return errorMessages
}

func UploadFile(file *multipart.FileHeader) (path string, err error) {
	src, err := file.Open()
	fileByte, err := ioutil.ReadAll(src)
	fileType := http.DetectContentType(fileByte)

	if fileType == "application/pdf" {
		path = "uploads/" + strconv.FormatInt(time.Now().Local().Unix(), 10) + ".pdf"
	} else if fileType == "text/plain; charset=utf-8" {
		path = "uploads/" + strconv.FormatInt(time.Now().Local().Unix(), 10) + ".txt"
	}

	err = ioutil.WriteFile(path, fileByte, 0777)
	defer src.Close()
	return
}

func APIResponse(message string, code int, status bool, data interface{}) Response {
	jsonResponse := Response{
		StatusCode: code,
		Success:    status,
		Message:    message,
		Data:       data,
	}

	return jsonResponse
}
