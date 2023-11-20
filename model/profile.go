package model

type Profile struct {
	Address      *Address      `bson:"address" json:"address"`
	PersonalInfo *PersonalInfo `bson:"personalInfo" json:"personalInfo"`
}

type PersonalInfo struct {
	AgeRange *AgeRange `bson:"ageRange" json:"ageRange,omitempty"`
	Gender   *Gender   `bson:"gender" json:"gender,omitempty"`
}

type AgeRange string
type Gender string
