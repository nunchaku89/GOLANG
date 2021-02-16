package model // import "model"

import (
	"github.com/dgrijalva/jwt-go"
	nullable "gopkg.in/guregu/null.v3"
)

// JwtCustomClaims -
type JwtCustomClaims struct {
	Idx   int    `json:"J_Idx"`
	Name  string `json:"J_Name"`
	Email string `json:"J_Email"`
	jwt.StandardClaims
}

// Person -
type Person struct {
	Idx   nullable.Int    `json:"P_Idx"`
	Name  nullable.String `json:"P_Name"`
	Email nullable.String `json:"P_Email"`
}

// Paging -
type Paging struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
}
