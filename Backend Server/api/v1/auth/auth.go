package auth

type Login struct {
	Login      string
	Password   string
	DomainName string
}

type ProjID struct {
	ProjectID string
}

type Auth struct {
	Auth      *Login
	ProjectID string
}

var InfoAuth Auth
