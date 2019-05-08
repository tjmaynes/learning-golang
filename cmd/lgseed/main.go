package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/tjmaynes/learning-golang/db"
	"github.com/tjmaynes/learning-golang/post"
)

func main() {
	var dbSource = flag.String("DB_SOURCE", os.Getenv("DB_SOURCE"), "Database url connection string.")
	var dbType = flag.String("DB_TYPE", os.Getenv("DB_TYPE"), "Database Type, such as postgres, mysql, etc.")

	flag.Parse()

	dbConn, err := db.ConnectDB(*dbSource, *dbType)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	postRepository := post.NewPostRepository(dbConn)
	ctx := context.Background()

	const jsonStream = `
	[
		{"Title": "My next great blog post", "Content": "Nothing to say right now..."},
		{"Title": "Great Expectations", "Content": "Love this book"}
	]
	`

	var posts []post.Post
	err = json.Unmarshal([]byte(jsonStream), &posts)
	if err != nil {
		panic(err)
	}

	var ids []int64
	for _, post := range posts {
		post, err := postRepository.AddPost(ctx, &post)
		if err != nil {
			panic(err)
		}
		ids = append(ids, post.ID)
	}

	var newPosts []*post.Post
	for _, id := range ids {
		newPost, err := postRepository.GetByPostID(ctx, id)
		if err != nil {
			panic(err)
		}
		newPosts = append(newPosts, newPost)
	}

	json, _ := json.MarshalIndent(&newPosts, "", "   ")
	fmt.Printf("ADDED:\n%s\n", json)
}