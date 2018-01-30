package types

type ListVersion struct {
	IdVersion        int       `json:"id_version"`
	CurrentVersion   bool      `json:"current_version"`
	Subject          string    `json:"subject"`
}
