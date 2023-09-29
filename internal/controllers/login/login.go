package login

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

func LoginController(router *gin.RouterGroup) {
	router.GET("/", renderLoginPage)
	router.POST("/", handleLogin)
	router.POST("/email", emailValidator)
}

type User struct {
	email    string
	password string
}

var userList = []User{
	{"test@test.com", "test"},
}

type LoginInput struct {
	Name      string
	Type      string
	Label     string
	Value     string
	Invalid   bool
	Error     string
	Validator string
}

var EmailInput = LoginInput{
	Name:      "email",
	Type:      "text",
	Label:     "Email",
	Invalid:   false,
	Error:     "",
	Validator: "/login/email",
}
var PasswordInput = LoginInput{
	Name:    "password",
	Type:    "password",
	Label:   "Password",
	Invalid: false,
	Error:   "",
}

type FormLogin struct {
	Invalid bool
	Error   string
	Inputs  []LoginInput
}

func renderLoginPage(c *gin.Context) {
	inputs := []LoginInput{EmailInput, PasswordInput}
	form := FormLogin{
		Invalid: false,
		Error:   "",
		Inputs:  inputs,
	}
	c.HTML(http.StatusOK, "views/loginPage", form)
}

func emailValidator(c *gin.Context) {
	email := c.Request.FormValue("email")
	emailInput := getEmailInput(email)
	c.HTML(http.StatusOK, "components/login/input", emailInput)
}

func handleLogin(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	if validateUser(email, password) {
		c.SetCookie("LOGGED", "true", 3600, "/", "localhost", false, true)

		if header := c.GetHeader("Hx-Request"); header == "true" {
			c.Header("HX-Location", "/")
		} else {
			c.Header("location", "/")
		}
		c.Status(http.StatusOK)
	} else {
		emailInput := getEmailInput(email)
		inputs := []LoginInput{emailInput, PasswordInput}
		form := FormLogin{
			Invalid: true,
			Error:   "Invalid user or password.",
			Inputs:  inputs,
		}
		c.HTML(http.StatusOK, "components/login/form", form)

	}
}

func includeUser(email string) bool {
	for _, user := range userList {
		if user.email == email {
			return true
		}
	}
	return false
}
func validateUser(email string, password string) bool {
	for _, user := range userList {
		if user.email == email && user.password == password {
			return true
		}
	}
	return false
}

func getEmailInput(email string) LoginInput {
	emailInput := EmailInput
	emailInput.Value = email
	if _, err := mail.ParseAddress(email); err != nil {
		emailInput.Invalid = true
		emailInput.Error = "Please, enter a valid email!"
	} else if !includeUser(email) {
		emailInput.Invalid = true
		emailInput.Error = "User not found."
	}
	return emailInput
}
