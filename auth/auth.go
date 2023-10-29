package auth

func Authenticate(username, password string) bool {
	return username == "user" && password == "pass"
}
