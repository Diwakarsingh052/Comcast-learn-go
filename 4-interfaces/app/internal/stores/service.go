package stores

import "app/internal/models"

type UserRepository interface {
	Create(u models.User) (models.User, bool)
	Update(id int, name string) (models.User, bool)
}

func CreateUser(ur UserRepository, u models.User) (models.User, bool) {
	return ur.Create(u)
}
