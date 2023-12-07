package model

type Authenticator struct {
	ID            int     `json:"user_id,omitempty"`
	Username      string  `json:"username,omitempty"`
	Password      string  `json:"password,omitempty"`
	Email         string  `json:"email,omitempty"`
	Role          string  `json:"role,omitempty"`
}