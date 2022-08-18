package model

type User struct {
	Name     string `bson:"name" json:"name,omitempty"`
	Phone    string `bson:"phone" json:"phone,omitempty"`
	Birthday string `bson:"birthday" json:"birthday,omitempty"`
	Gender   uint8  `bson:"gender" json:"gender,omitempty"`
}
