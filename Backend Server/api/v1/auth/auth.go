package auth

type Keys struct {
	AKey       string `json:"aKey"`
	SKey       string `json:"sKey"`
	DomainName string `json:"domain"`
}

type ProjID struct {
	ProjectID string `json:"projectID"`
}

type Auth struct {
	Signer    *Keys
	ProjectID string
}

var InfoAuth Auth
