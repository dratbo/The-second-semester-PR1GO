package service

func GenerateToken(username string) (string, error) {
	return "demo-token", nil
}

func VerifyToken(token string) (bool, string) {
	if token == "demo-token" {
		return true, "student"
	}
	return false, ""
}
