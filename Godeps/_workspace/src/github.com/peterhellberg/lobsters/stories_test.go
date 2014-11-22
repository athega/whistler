package lobsters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStory_jyloq8(t *testing.T) {
	ts, c := testServerAndClientByFixture("jyloq8")
	defer ts.Close()

	story, err := c.Story("jyloq8")

	assert.Nil(t, err)
	assert.Equal(t, "jcs", story.Submitter.Username)
	assert.Equal(t, 8, story.Score)
	assert.Equal(t, "https://ello.co/gb/post/knOWk-qeTqfSpJ6f8-arCQ", story.URL)
	assert.Equal(t, 3, story.Comments[2].IndentLevel)
}

func TestHottest(t *testing.T) {
	ts, c := testServerAndClientByFixture("hottest")
	defer ts.Close()

	stories, err := c.Hottest()

	assert.Nil(t, err)
	assert.Equal(t, "hmarr", stories[2].Submitter.Username)
	assert.Equal(t, "sevan", stories[6].Submitter.Username)
	assert.Equal(t, 5, stories[5].Score)
	assert.Equal(t, "http://fetching.io/", stories[10].URL)
	assert.Equal(t, 0, stories[7].CommentCount)
}

func TestNewest(t *testing.T) {
	ts, c := testServerAndClientByFixture("newest")
	defer ts.Close()

	stories, err := c.Newest()

	assert.Nil(t, err)
	assert.Equal(t, "cmeiklejohn", stories[3].Submitter.Username)
	assert.Equal(t, 2, stories[4].Score)
	assert.Equal(t, "https://www.youtube.com/watch?v=8_z9-iRiSZw", stories[5].URL)
	assert.Equal(t, 6, stories[1].CommentCount)
}
