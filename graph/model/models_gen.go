// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Link struct {
	ID      string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string `json:"title,omitempty" bson:"title,omitempty"`
	Address string `json:"address,omitempty" bson:"address,omitempty"`
	User    *User  `json:"user,omitempty" bson:"user,omitempty"`
}

type Login struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type NewLink struct {
	Title   string `json:"title,omitempty" bson:"title,omitempty"`
	Address string `json:"address,omitempty" bson:"address,omitempty"`
}

type NewUser struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type RefreshTokenInput struct {
	Token string `json:"token,omitempty" bson:"token,omitempty"`
}

type User struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}
