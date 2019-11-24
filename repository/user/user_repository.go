package user

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"log"
	"ssnbackend/repository/config"
	"ssnbackend/repository/models"
	"time"
)

func InsertUser(user models.UserModel) error {
	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	user.Enabled = true
	user.CreatedAt = time.Now().UTC()
	user.ModifiedAt = time.Time{}
	user.DisabledAt = time.Time{}

	err = db.Save(&user)

	if err != nil {
		return err
	}

	return nil
}

func SetAvatar(binaryImage []byte, contentType string, userID uint64) error {
	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	avatar := models.UserAvatarModel{}

	err = db.Select(q.Eq("UserID", userID)).First(&avatar)

	if err == nil {
		avatar.BinaryImage = binaryImage
		avatar.BinaryContentType = contentType
		avatar.ModifiedAt = time.Now().UTC()
		err = db.Update(&avatar)
	} else {
		avatar.BinaryImage = binaryImage
		avatar.BinaryContentType = contentType
		avatar.UserID = userID
		avatar.Enabled = true
		avatar.CreatedAt = time.Now().UTC()
		avatar.ModifiedAt = time.Time{}
		avatar.DisabledAt = time.Time{}
		err = db.Save(&avatar)
	}

	if err != nil {
		return err
	}

	return nil
}

func GetAvatar(userID uint64) (models.UserAvatarModel, error) {
	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var avatar models.UserAvatarModel

	err = db.Select(q.Eq("UserID", userID)).First(&avatar)

	if err != nil {
		return models.UserAvatarModel{}, err
	}

	return avatar, nil
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
