package model

type Candidate struct {
	ID         uint   `gorm:"autoIncrement;primaryKey" json:"id"`
	Name       string `json:"name"`
	DOB        string `json:"dob"`
	BioLink    string `json:"bioLink"`
	ImageLink  string `json:"imageLink"`
	Policy     string `json:"policy"`
	VotedCount uint   `json:"voutedCount"`
}

type Vote struct {
	ID          uint `gorm:"autoIncrement;primaryKey"`
	NationalID  uint `json:"nationalId"`
	CandidateID uint `json:"candidateId"`
}

type System struct {
	Status bool `json:"status"`
}

type InputCandidate struct {
	Name      string `json:"name"`
	DOB       string `json:"dob"`
	BioLink   string `json:"bioLink"`
	ImageLink string `json:"imageLink"`
	Policy    string `json:"policy"`
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
	ID          uint `json:"id"`
	VoutedCount uint `json:"voutedCount"`
}
