package repositories

func GetHealth() map[string]string {
	return map[string]string{
		"status": "UP",
	}
}
