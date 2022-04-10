// +build integration

package db

import (
	"testing"
 
	"github.com/stretchr/testify/assert"
	"context"
	"github.com/Kylescottw/pulse-api/internal/comment"
)


func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug: "slug",
			Author: "author",
			Body: "body",
		})

		assert.NoError(t, err )
		
		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)
 
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug: "new-slug",
			Author: "tester",
			Body: "body",
		})
		assert.NoError(t, err )

		err = db.DeleteComment(context.Background(), cmt.ID)

		assert.NoError(t, err )

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)

	})
}