package model

import "time"

type Candidate struct {
	ID         uint   `json:"id" gorm:"primary_key; not null; auto_increment" validate:"required"`
	Name       string `json:"name" validate:"required,min=3,max=50"`
	DOB        string `json:"dob" validate:"required,min=3,max=50"`
	BioLink    string `json:"bioLink" validate:"required,min=5"`
	ImageLink  string `json:"imageLink" validate:"required,min=5"`
	Policy     string `json:"policy" validate:"required,min=10"`
	VotedCount *uint  `json:"votedCount" gorm:"default:0"`
}

type Vote struct {
	ID          uint   `json:"id" gorm:"primary_key; not null; auto_increment"`
	NationalID  string `json:"nationalId"`
	CandidateID uint   `json:"candidateId"`
}

type LogVote struct {
	LogID      uint      `json:"log_id" gorm:"primary_key; not null; auto_increment"`
	ID         uint      `json:"id"`
	VotedCount uint      `json:"votedCount"`
	CreatedAt  time.Time `json:"created_at"`
}

type InputCandidate struct {
	Name      string `json:"name" validate:"required,min=3,max=50"`
	DOB       string `json:"dob" validate:"required,min=3,max=50"`
	BioLink   string `json:"bioLink" validate:"required,min=5"`
	ImageLink string `json:"imageLink" validate:"required,min=5"`
	Policy    string `json:"policy" validate:"required,min=10"`
}

type InputCheckVote struct {
	NationalID string `json:"nationalId" validate:"required"`
}

type InputVote struct {
	NationalID  string `json:"nationalId" validate:"required"`
	CandidateID uint   `json:"candidateId" validate:"required"`
}

type InputToggleElection struct {
	Enable *bool `json:"enable" validate:"required"`
}

type ResponseElectionResult struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	DOB        string `json:"dob"`
	BioLink    string `json:"bioLink"`
	ImageLink  string `json:"imageLink"`
	Policy     string `json:"policy"`
	VotedCount uint   `json:"votedCount"`
	Percentage string `json:"percentage"`
}

type ResponseElectionCount struct {
	ID         uint `json:"id"`
	VotedCount uint `json:"votedCount"`
}

type ErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}
