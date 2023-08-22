package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_PostLing(t *testing.T) {
	t.Run("Given valid url, when creating a post link, then it should return a post link", func(t *testing.T) {
		title := "title"
		link := "https://google.com/path/to/post"
		postLink, err := NewPostLink(title, link, "")
		require.NoError(t, err)
		assert.Equal(t, title, postLink.Title)
		assert.Equal(t, link, postLink.Link.String())
	})

	t.Run("Given a url without protocol, when creating a post link, then it should return a post link", func(t *testing.T) {
		title := "title"
		link := "/path/to/post"
		path := "martinfowler.com"
		postLink, err := NewPostLink(title, link, path)
		require.NoError(t, err)
		assert.Equal(t, title, postLink.Title)
		assert.Equal(t, "https://"+path+link, postLink.Link.String())
	})

	t.Run("Given invalid url, when creating a post link, then it should return an error", func(t *testing.T) {
		title := "title"
		link := "invalid url"
		_, err := NewPostLink(title, link, "")
		require.Error(t, err)
	})
}
