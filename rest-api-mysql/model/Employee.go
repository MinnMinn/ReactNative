package model

type Employee struct {
	Id    int		`json:"id,omitempty" bson:"_id,omitempty"`
	Name  string	`json:"name" bson:"name"`
	City string		`json:"city" bson:"city"`
}