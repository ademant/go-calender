package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "go-calender/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
//	"fmt"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	gorm.Model
	User     string `json:"user"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

type AccountPub struct {
	User     string `json:"user"`
	Email    string `json:"email"`
	Token    string `json:"token";sql:"-"`
}

func (ac *Account) DBAccountInit() {
	// ensure admin user exist
	// initial password
	initPass := "adMin"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(initPass), bcrypt.DefaultCost)
	initAccount := Account{User:"admin",Password:string(hashedPassword)}
	GetDB().Where(Account{User:"admin"}).FirstOrCreate(&initAccount)
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	temp := &Account{}
//Username must be unique
	err := GetDB().Table("accounts").Where("user = ?", account.User).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.User != "" {
		return u.Message(false, "Username already in use."), false
	}
	
	//Email must be unique
	if account.Email != "" {
	//check for errors and duplicate emails
		err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return u.Message(false, "Connection error. Please retry"), false
		}
		if temp.Email != "" {
			return u.Message(false, "Email address already in use by another user."), false
		}
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string]interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(user, password string) (map[string]interface{}) {

	account := &Account{}
	GetDB().Where(Account{User:user}).First(&account)
	if account.ID == 0 {
			return u.Message(false, "User not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response
	GetDB().Save(&account)

	pubacc := &AccountPub{}
	pubacc.User = account.User;
	pubacc.Email = account.Email;
	pubacc.Token = account.Token;
	resp := u.Message(true, "Logged In")
	resp["account"] = pubacc
	return resp
}

func Logout(user uint) (map[string]interface{}) {
	
	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", user).First(acc)
	if acc.User == "" { //User not found!
		return nil
	}
	acc.Token = ""
	GetDB().Save(&acc)
	resp := u.Message(true, "Logged Out")
	resp["account"] = acc.User 
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.User == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
