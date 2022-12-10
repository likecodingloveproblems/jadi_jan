package models

import (
	"context"
	"testing"
)

func TestCommentCreate(t *testing.T) {
	ctx := context.Background()
	defer tearDown(redisStore, ctx)
	comment := &Comment{FullName: "full name", Description: "some description here"}
	comment.Create(ctx)
	storedComment := getCommentByCount(ctx, 1)
	if comment.FullName != storedComment.FullName && comment.Description != storedComment.Description {
		t.Error("Some thing went wrong, maybe comment is not saved, full name:" +
			comment.FullName + ", description: " + comment.Description)
	}
}

func TestGetComments(t *testing.T) {
	var (
		comment  Comment
		comments Comments
	)
	ctx := context.Background()
	defer tearDown(redisStore, ctx)
	for i := 0; i < 5; i++ {
		comment.FullName = "full name " + string(rune(i))
		comment.Description = "description " + string(rune(i))
		comment.Create(ctx)
	}
	comments.Get(ctx)
	if len(comments.Comments) != 5 {
		t.Error("5 Comments are created, be " + string(rune(len(comments.Comments))) + " comments have gotten.")
	}
}
