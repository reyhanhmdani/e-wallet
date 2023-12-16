package util

func CreateCriteria(username, email string) map[string]interface{} {
	return map[string]interface{}{"username": username, "email": email}
}
