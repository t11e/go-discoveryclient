package discoveryclient

type ExportQuery struct {
	Scores           bool        `json:"scores,omitempty"`
	Sources          bool        `json:"sources,omitempty"`
	Criteria         []Criterion `json:"criteria,omitempty"`
	ExactMatchesOnly bool        `json:"exactMatchesOnly,omitempty"`
	Values           []string    `json:"values,omitempty"`
}
