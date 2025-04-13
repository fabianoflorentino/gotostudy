package services

import "github.com/fabianoflorentino/gotostudy/repositories"

func GetHealth() (map[string]string, error) {
	return repositories.GetHealth(), nil
}
