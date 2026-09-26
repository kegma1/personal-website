package internal

import "os"

func Get_blog_database_path() string {
	if path := os.Getenv("BLOG_DATABASE_PATH"); path != "" {
		return path
	}
	return "./data/db.db"
}
