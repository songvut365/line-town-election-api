package model

type Candidate struct {
	ID         uint   `json:"id" gorm:"primary_key; not null; auto_increment"`
	Name       string `json:"name"`
	DOB        string `json:"dob"`
	BioLink    string `json:"bioLink"`
	ImageLink  string `json:"imageLink"`
	Policy     string `json:"policy"`
	VotedCount uint   `json:"votedCount" gorm:"default:0"`
}

type Vote struct {
	ID          uint   `json:"id" gorm:"primary_key; not null; auto_increment"`
	NationalID  string `json:"nationalId"`
	CandidateID uint   `json:"candidateId"`
}

type InputCandidate struct {
	Name      string `json:"name"`
	DOB       string `json:"dob"`
	BioLink   string `json:"bioLink"`
	ImageLink string `json:"imageLink"`
	Policy    string `json:"policy"`
}

type InputCheckVote struct {
	NationalID string `json:"nationalId"`
}

type InputVote struct {
	NationalID  string `json:"nationalId"`
	CandidateID uint   `json:"candidateId"`
}

type InputToggleElection struct {
	Enable bool `json:"enable"`
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
