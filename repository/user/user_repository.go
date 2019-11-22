package user

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"log"
	"ssn-backend/repository/config"
	"ssn-backend/repository/models"
	"time"
)

func InsertUser(user models.UserModel) error {
	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	user.Enabled = true
	user.CreatedAt = time.Now()
	user.ModifiedAt = time.Time{}
	user.DisabledAt = time.Time{}

	err = db.Save(&user)

	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (models.UserModel, error) {
	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var userInDb models.UserModel

	err = db.Select(q.Eq("Username", username)).First(&userInDb)

	if err != nil {
		return models.UserModel{}, err
	}

	return userInDb, nil
}

func GetUserById(userId uint64) (models.UserModel, error) {
	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var userInDb models.UserModel

	err = db.Select(q.Eq("ID", userId)).First(&userInDb)

	if err != nil {
		return models.UserModel{}, err
	}

	return userInDb, nil
}

func UpdateUser(user models.UserModel) error {
	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var userInDb models.UserModel

	err = db.Select(q.Eq("ID", user.ID), q.Eq("Enabled", true)).First(&userInDb)

	//err = db.One("ID", user.ID, &userInDb)

	user.Enabled = true
	user.ModifiedAt = time.Now()
	user.CreatedAt = userInDb.CreatedAt
	user.DisabledAt = userInDb.DisabledAt

	err = db.Update(&user)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(idUser uint64) error {

	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var userInDb models.UserModel
	err = db.One("ID", idUser, &userInDb)

	userInDb.Enabled = false
	userInDb.DisabledAt = time.Now()

	err = db.Update(&userInDb)

	if err != nil {
		return err
	}

	return nil
}
