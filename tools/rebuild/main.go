package main

import (
	"log"
	"os"
	"strings"
	"time"

	"context"
	"database/sql"
	"personal_website/repositories"

	"github.com/BurntSushi/toml"
	_ "modernc.org/sqlite"
)

var POST_PATH = "./content/posts/"

type postMeta struct {
	Title    string
	Subtitle string
	Language string
	Tags     string
	Created_at time.Time
	Edited_at time.Time
}

func parsePost(postName string) (*postMeta, string, error) {
	postRaw, err := os.ReadFile(POST_PATH + postName)
	if err != nil {
		return &postMeta{}, "", err
	}

	postContent := strings.Split(string(postRaw), "+++")

	meta := postMeta {}

	if _, err := toml.Decode(postContent[0], &meta); err != nil {
		return &postMeta{}, "", err
	}
	return &meta, postContent[1], nil
}

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.New(db)

	postPaths, err := os.ReadDir(POST_PATH)
	if err != nil {
		log.Fatal(err)
	}

	repo.CleanPosts(ctx)

	for _, path := range postPaths {
		metadata, content, err := parsePost(path.Name())
		if err != nil {
			log.Fatal(err)
		}
		repo.CreatePost(ctx, repositories.CreatePostParams{
			Content: content,
			Title: metadata.Title,
			Subtitle: metadata.Subtitle,
			Tags: metadata.Tags,
			Language: metadata.Language,
			CreatedAt: metadata.Created_at,
			EditedAt: metadata.Edited_at,
		})
	}
}
