package auth

type Keys struct {
	AKey string `json:"aKey"`
	SKey string `json:"sKey"`
}

type ProjID struct {
	ProjectID string `json:"projectID"`
}

type Auth struct {
	AuthKeys  *Keys
	ProjectID string
}

var Info Auth
