package queries

import (
	"github.com/abhishek-bhangalia-busy/banking-api/db"
	"github.com/abhishek-bhangalia-busy/banking-api/models"
)

func CreateUser(user *models.User) (uint64, error) {
	_, insertErr := db.DB.Model(user).Returning("id").Insert()

	if insertErr != nil {
		return 0, insertErr
	}
	return user.ID, nil
}

func SelectUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := db.DB.Model(user).Where("email = ?", email).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}


func SelectUserByID(id uint64) (*models.User, error) {
	user := new(models.User)
	err := db.DB.Model(user).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}