package model

type Preferences struct {
	StoreIDs []string `bson:"storeIds" json:"storeIds"`
}
