package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Post(t *testing.T) {
	author := "author"
	t.Run("Given valid url, when creating a post link, then it should return a post link", func(t *testing.T) {
		title := "title"
		link := "https://google.com/path/to/post"
		post, err := NewPost(title, link, "", author)
		require.NoError(t, err)
		assert.Equal(t, title, post.Title)
		assert.Equal(t, link, post.Link.String())
		assert.IsType(t, post.Id, uuid.UUID{})
	})

	t.Run("Given a url without protocol, when creating a post link, then it should return a post link", func(t *testing.T) {
		title := "title"
		link := "/path/to/post"
		path := "martinfowler.com"
		post, err := NewPost(title, link, path, author)
		require.NoError(t, err)
		assert.Equal(t, title, post.Title)
		assert.Equal(t, "https://"+path+link, post.Link.String())
	})

	t.Run("Given invalid url, when creating a post link, then it should return an error", func(t *testing.T) {
		title := "title"
		link := "invalid url"
		_, err := NewPost(title, link, "", author)
		require.Error(t, err)
	})
}
