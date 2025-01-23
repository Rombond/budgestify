package user

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
	Hash  []byte `json:"-"`
}

type UserForm struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
	Hash  string `json:"hash"`
}
