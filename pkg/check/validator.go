package check

import (
	"errors"
	"time"
	"regexp"
	"net/mail"
	"encoding/json"
	"fmt"
	"net/http"
)


func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}


	
	// func ValidateEmail(email string) error {
	// 	emailRegex := `^[a-zA-Z0-9._%+-]+@(?:gmail|email)+(?:com|ru)$`
	// 	regex := regexp.MustCompile(emailRegex)
	// 	if regex.MatchString(email) {
	// 		return nil
	// 	} else {
	// 		return errors.New("email is not valid")
	// 	}
	// }

	func ValidateEmail(address string) (error) {
		_, err := mail.ParseAddress(address)
		if err != nil {
			return  errors.New("email is not valid")
			
		}
		return nil
	}
	
	type EmailVerificationResponse struct {
		Data struct {
			Result string `json:"result"`
		} `json:"data"`
	}
	
	func CheckEmail(email string) (error) {
		apiKey := "a78afa97d76af0e3364a3eb68ed12aae83e247a0"       
	
		url := fmt.Sprintf("https://api.hunter.io/v2/email-verifier?email=%s&api_key=%s", email, apiKey)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		defer resp.Body.Close()
	
		var verificationResponse EmailVerificationResponse
		err = json.NewDecoder(resp.Body).Decode(&verificationResponse)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return err
		}


		if verificationResponse.Data.Result == "undeliverable" {
			fmt.Println("Email address does not exist or is undeliverable")

			return errors.New("Email address does not exist or is undeliverable")
		} else if verificationResponse.Data.Result == "deliverable" {
			fmt.Println("Email address is valid")
			return nil
		} else {
			fmt.Println("Unable to verify email address")
			return errors.New("Unable to verify email address")
		}
	}
	
	func ValidatePassword(password string) error {
		lowercaseRegex := `[a-z]`
		hasLowercase, _ := regexp.MatchString(lowercaseRegex, password)
		
		uppercaseRegex := `[A-Z]`
		hasUppercase, _ := regexp.MatchString(uppercaseRegex, password)

		digitRegex := `[0-9]`
		hasDigit, _ := regexp.MatchString(digitRegex, password)
		
		symbolRegex := `[!@#$%^&*()-_+=~\[\]{}|\\:;"'<>,.?\/]`
		hasSymbol, _ := regexp.MatchString(symbolRegex, password)
	
		if hasLowercase && hasUppercase && hasDigit && hasSymbol && len(password) >= 8 {
			return nil
		}
	
		return errors.New("password does not meet the criteria")
	}
	
	
	
	
	func ValidatePhone(phone string) error {
		
		phoneRegex := `^\+998\d{9}$`
	
		regex := regexp.MustCompile(phoneRegex)
	
		if regex.MatchString(phone) {
			return nil
		} else {
			return errors.New("phone number is not valid")
		}
	}
	
	var ORDER_STATUS = []string{
		"new", "in-process", "finished", "canceled",
	}


	func CheckOrderStatus(status string) error {
		for i:=0;i<4;i++ {
			if ORDER_STATUS[i] == status {
				return nil
			}
		}
		return errors.New("error: Invalid order status")
	}
