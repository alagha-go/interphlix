package accounts

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type GoogleAccount struct {
	ID									primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	GoogleID							string										`json:"id,omitempty" bson:"id,omitempty"`
	Email								string										`json:"email,omitempty" bson:"email,omitempty"`
	EmailVerified						bool										`json:"verified_email,omitempty" bson:"verified_email,omitempty"`
	Name								string										`json:"name,omitempty" bson:"name,omitempty"`
	GivenName							string										`json:"given_name,omitempty" bson:"given_name,omitempty"`
	FamilyName							string										`json:"family_name,omitempty" bson:"family_name,omitempty"`
	Picture								string										`json:"picture,omitempty" bson:"picture,omitempty"`
	Locale								string										`json:"locale,omitempty" bson:"locale,omitempty"`
}

type Account struct {
	ID									primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	GoogleID							string										`json:"id,omitempty" bson:"id,omitempty"`
	SignUpMethod						string										`json:"signup_method,omitempty" bson:"signup_method,omitempty"`
	Email								string										`json:"email,omitempty" bson:"email,omitempty"`
	EmailVerified						bool										`json:"email_verified,omitempty" bson:"email_verified,omitempty"`
	UserName							string										`json:"user_name,omitempty" bson:"user_name,omitempty"`
	FirstName							string										`json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName							string										`json:"last_name,omitempty" bson:"last_name,omitempty"`
	Photo								string										`json:"photo,omitempty" bson:"photo,omitempty"`
	Token								*Token										`json:"token,omitempty" bson:"token,omitempty"`
	NewPassword							string										`json:"new_password,omitempty" bson:"new_password,omitempty"`
	TimeCreated							time.Time									`json:"time_created,omitempty" bson:"time_created,omitempty"`
	Type								string										`json:"type,omitempty" bson:"type,omitempty"`
	Locale								string										`json:"locale,omitempty" bson:"locale,omitempty"`
	Password							string										`json:"password,omitempty" bson:"password,omitempty"`
	Deleted								bool										`json:"deleted,omitempty" bson:"deleted,omitempty"`
}

type Token struct {
	AccessToken							string										`json:"access_token,omitempty" bson:"access_token,omitempty"`
	TokenType							string										`json:"token_type,omitempty" bson:"token_type,omitempty"`
	RefreshToken						string										`json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	Expiry								time.Time									`json:"expiry,omitempty" bson:"expiry,omitempty"`
	Raw									interface{}									`json:"raw,omitempty" bson:"raw,omitempty"`
}