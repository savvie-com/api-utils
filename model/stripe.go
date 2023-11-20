package model

type StripeCustomer struct {
	ID    string `bson:"id" json:"id"`
	Email string `bson:"email" json:"email"`
	Name  string `bson:"name" json:"name"`
}

type StripeSubscription struct {
	ID     string `bson:"id" json:"id"`
	Status string `bson:"status,omitempty" json:"status"`
}

type StripeUser struct {
	CustomerID   string              `bson:"customerId" json:"customerId"`
	Subscription *StripeSubscription `bson:"subscription,omitempty" json:"subscription,omitempty"`
}
