package postgres

import (
	"redis/models"
	"strconv"
	"strings"
)

func (p *Postgres) GetUser(key string) (models.User, error) {
	var user models.User

	id, err := strconv.Atoi(strings.Split(key, ":")[1])
	if err != nil {
		return user, err
	}

	// result := p.db.First(&user, id)
	// return user, result.Error
	user = models.User{
		Id:   id,
		Name: "Alex",
		Age:  12,
	}
	return user, nil
}
