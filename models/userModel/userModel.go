package userModel

type User struct {
	ID   int    `json:"id"`
	PW   string `json:"pw"`
	Name string `json:"name"`
}
