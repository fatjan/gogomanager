package dto

import (
	"errors"
	"regexp"
	"strings"

	"github.com/fatjan/gogomanager/internal/pkg/exceptions"
	"golang.org/x/crypto/bcrypt"
)

type ActionType string

const (
	Create ActionType = "create"
	Login  ActionType = "login"
)

type AuthRequest struct {
	Name     string
	Email    string
	Password string
	Action   string
}

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (d *AuthRequest) ValidatePayloadAuth() error {
	if d.Email == "" {
		return exceptions.ErrorBadRequest
	}
	if d.Password == "" {
		return exceptions.ErrorBadRequest
	}
	if d.Action == "" {
		return exceptions.ErrorBadRequest
	}
	if !isValidEmail(d.Email) {
		return exceptions.ErrorBadRequest
	}
	if !isValidPasswordLength(d.Password, 8, 32) {
		return exceptions.ErrorBadRequest
	}
	if !isValidAction(d.Action) {
		return exceptions.ErrorBadRequest
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPasswordLength(password string, minLength, maxLength int) bool {
	passwordLength := len(password)
	return passwordLength >= minLength && passwordLength <= maxLength
}

func isValidAction(action string) bool {
	validActions := map[string]bool{
		"create": true,
		"login":  true,
	}
	return validActions[action]
}

func (d *AuthRequest) SetName() {
	parts := strings.Split(d.Email, "@")
	localPart := parts[0]
	name := strings.ReplaceAll(localPart, ".", " ")
	d.Name = name
}

func (d *AuthRequest) HashPassword() error {
	resultHash, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("ERROR HASHING")
	}
	d.Password = string(resultHash)

	return nil
}

func (d *AuthRequest) ComparePassword(password string) error {
	errCompare := bcrypt.CompareHashAndPassword([]byte(password), []byte(d.Password))
	if errCompare != nil {
		return errors.New("password or username is wrong")
	}
	return nil
}
