package models

import (
	"context"
	"strconv"
)

type Comment struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
}

type Comments struct {
	Comments []Comment
}

var commentsCount string = "comments_count"

func getCommentsCount(ctx context.Context) uint64 {
	comments_count := redisStore.db.Get(ctx, commentsCount).Val()
	c, _ := strconv.ParseUint(comments_count, 10, 64)
	return c
}

func incrCommentsCount(ctx context.Context) uint64 {
	c := redisStore.db.Incr(ctx, commentsCount).Val()
	return uint64(c)
}

func getCommentKey(ctx context.Context, counter uint64) (fullNameKey, descriptionKey string) {
	fullNameKey = ":comments:" + string(rune(counter)) + ":full_name"
	descriptionKey = ":comments:" + string(rune(counter)) + ":description"
	return
}

func getNewCommentKey(ctx context.Context) (fullNameKey, descriptionKey string) {
	counter := incrCommentsCount(ctx)
	fullNameKey, descriptionKey = getCommentKey(ctx, counter)
	return
}

func getCommentByCount(ctx context.Context, count uint64) *Comment {
	fullNameKey, descriptionKey := getCommentKey(ctx, count)
	fullName := redisStore.db.Get(ctx, fullNameKey).Val()
	description := redisStore.db.Get(ctx, descriptionKey).Val()
	return &Comment{FullName: fullName, Description: description}
}

func (c *Comment) Create(ctx context.Context) {
	fullNameKey, descriptionKey := getNewCommentKey(ctx)
	redisStore.db.Set(ctx, fullNameKey, c.FullName, 0)
	redisStore.db.Set(ctx, descriptionKey, c.Description, 0)
}

func (c *Comments) Get(ctx context.Context) {
	var (
		comment       *Comment
		commentsCount uint64
		count         uint64
	)
	commentsCount = getCommentsCount(ctx)
	for count = commentsCount; count > 0; count-- {
		comment = getCommentByCount(ctx, count)
		c.Comments = append(c.Comments, *comment)
	}
}
