package lobsters

import (
	"fmt"
	"time"
)

// User represents a user
type User struct {
	Username    string    `json:"username"`
	CreatedAt   time.Time `json:"created_at"`
	IsAdmin     bool      `json:"is_admin"`
	IsModerator bool      `json:"is_moderator"`
	AvatarURL   string    `json:"avatar_url"`
}

// Comment represents a comment
type Comment struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ID          string    `json:"short_id"`
	IsDeleted   bool      `json:"is_deleted"`
	IsModerated bool      `json:"is_moderated"`
	Score       int       `json:"score"`
	Comment     string    `json:"comment"`
	URL         string    `json:"url"`
	IndentLevel int       `json:"indent_level"`
	User        *User     `json:"commenting_user"`
}

// Story represents a story
type Story struct {
	CreatedAt    time.Time  `json:"created_at"`
	URL          string     `json:"url"`
	Title        string     `json:"title"`
	ID           string     `json:"short_id"`
	Score        int        `json:"score"`
	CommentCount int        `json:"comment_count"`
	Description  string     `json:"description"`
	CommentsURL  string     `json:"comments_url"`
	Submitter    *User      `json:"submitter_user"`
	Tags         []string   `json:"tags"`
	Comments     []*Comment `json:"comments"`
}

// StoriesService communicates with the stories
// related endpoints in the Lobsters API
type StoriesService interface {
	Get(id string) (*Story, error)
	All(path string) ([]*Story, error)
}

// storiesService implements StoriesService.
type storiesService struct {
	client *Client
}

// Story is a convenience method proxying Stories.Get
func (c *Client) Story(id string) (*Story, error) {
	return c.Stories.Get(id)
}

// Hottest is a convenience method proxying Stories.All("hottest.json")
func (c *Client) Hottest() ([]*Story, error) {
	return c.Stories.All("hottest.json")
}

// Newest is a convenience method proxying Stories.All("newest.json")
func (c *Client) Newest() ([]*Story, error) {
	return c.Stories.All("newest.json")
}

// Get retrieves a story with the given id
func (s *storiesService) Get(id string) (*Story, error) {
	req, err := s.client.NewRequest(fmt.Sprintf("s/%v.json", id))
	if err != nil {
		return nil, err
	}

	var story Story
	_, err = s.client.Do(req, &story)
	if err != nil {
		return nil, err
	}

	return &story, nil
}

// All retrieves the stories for the given path
func (s *storiesService) All(path string) ([]*Story, error) {
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, err
	}

	var stories []*Story
	_, err = s.client.Do(req, &stories)
	if err != nil {
		return nil, err
	}

	return stories, nil
}
