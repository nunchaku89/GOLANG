package model // import "model"

import (
	nullable "gopkg.in/guregu/null.v3"
)

// Person : 
type Person struct {
	Idx   nullable.Int	`json:"P_Idx"`
	Name  nullable.String	`json:"P_Name"`
	Email nullable.String	`json:"P_Email"`
}

// Paging - 
type Paging struct {
	Limit nullable.Int				`json:"limit"`
	Offset nullable.Int				`json:"offset"`
}
