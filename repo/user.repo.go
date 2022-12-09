package repo

import (
	"fp4/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByUserID(userID string) (models.User, error)
}

type userRepo struct {
	connection *gorm.DB
}

func NewUserRepo(connection *gorm.DB) UserRepository {
	return &userRepo{
		connection: connection,
	}
}

func (c *userRepo) InsertUser(user models.User) (models.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	c.connection.Save(&user)
	return user, nil
}

func (c *userRepo) UpdateUser(user models.User) (models.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		c.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	c.connection.Save(&user)
	return user, nil
}

func (c *userRepo) FindByEmail(email string) (models.User, error) {
	var user models.User
	res := c.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (c *userRepo) FindByUserID(userID string) (models.User, error) {
	var user models.User
	res := c.connection.Where("id = ?", userID).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
