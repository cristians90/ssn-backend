package post

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"log"
	"ssnbackend/repository/config"
	"ssnbackend/repository/models"
	"time"
)

func InsertPost(post models.PostModel) error {
	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	post.Enabled = true
	post.CreatedAt = time.Now().UTC()
	post.ModifiedAt = time.Time{}
	post.DisabledAt = time.Time{}

	err = db.Save(&post)

	if err != nil {
		return err
	}

	return nil
}

func GetPosts(sinceDate time.Time, limitResults int, offsetResults int) ([]models.PostModelForApi, error) {
	db, err := storm.Open(config.DatabaseFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var postsFromDb []models.PostModel

	err = db.Select(
		q.Lte("CreatedAt", sinceDate),
		q.Eq("Enabled", true)).
		Limit(limitResults).
		Skip(offsetResults).
		//Reverse().
		Find(&postsFromDb)

	var posts []models.PostModelForApi

	if err != nil {
		return posts, err
	}

	for _, item := range postsFromDb {
		posts = append(posts, item.GetPostModelForApi())
	}

	return posts, nil
}
