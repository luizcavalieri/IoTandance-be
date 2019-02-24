package auth

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}
