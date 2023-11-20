package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// roles Paid, Premium, Marketing, Admin, SuperAdmin
// role bits
// 1 + 2 + 4 + 8 + 16 = 31
// for admin, total is 15
const (
	RoleBitsPaid = 1 << iota
	RoleBitsPremium
	RoleBitsMarketing
	RoleBitsAdmin
	RoleBitsSuperUser
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Email        string             `bson:"email"`
	FirstName    string             `bson:"firstName"`
	LastName     string             `bson:"lastName"`
	Password     string             `bson:"password"`
	Preferences  *Preferences       `bson:"preferences"`
	Profile      *Profile           `bson:"profile"`
	RoleBits     int32              `bson:"roleBits"`
	Subscription *Subscription      `bson:"subscription"`
	AuthTS       time.Time          `bson:"authTS,omitempty"`
	CreateTS     time.Time          `bson:"createTS"`
	UpdateTS     time.Time          `bson:"updateTS"`
}

type Subscription struct {
	Active   bool        `bson:"active" json:"active"`
	Stripe   *StripeUser `bson:"stripe,omitempty" json:"stripe,omitempty"`
	UpdateTS time.Time   `bson:"updateTS,omitempty"`
}

func (User) IsEntity() {}
