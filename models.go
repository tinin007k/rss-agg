package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/tintin007k/rss-agg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	Title       string     `json:"title"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedsToFeed(feeds []database.Feed) []Feed {
	feedList := make([]Feed, len(feeds))
	for i, feed := range feeds {
		feedList[i] = databaseFeedToFeed(feed)
	}
	return feedList
}

func databaseFeedFollowsToFeedFollows(feedFollows database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        feedFollows.ID,
		CreatedAt: feedFollows.CreatedAt,
		UpdatedAt: feedFollows.UpdatedAt,
		UserID:    feedFollows.UserID,
		FeedID:    feedFollows.FeedID,
	}
}

func databaseFeedsFollowToFeedsFollow(feedFollows []database.FeedFollow) []FeedFollows {
	feedList := make([]FeedFollows, len(feedFollows))
	for i, feed := range feedFollows {
		feedList[i] = databaseFeedFollowsToFeedFollows(feed)
	}
	return feedList
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Description: stringSqlNultoString(post.Description),
		PublishedAt: timeSqlNultoTime(post.PublishedAt),
		Url:         post.Url,
		FeedID:      post.FeedID,
		Title:       post.Title,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	postList := make([]Post, len(posts))
	for i, post := range posts {
		postList[i] = databasePostToPost(post)
	}
	return postList
}

func stringSqlNultoString(sqlString sql.NullString) *string {
	if sqlString.Valid == true {
		return &sqlString.String
	}
	return nil
}

func timeSqlNultoTime(sqlTime sql.NullTime) *time.Time {
	if sqlTime.Valid == true {
		return &sqlTime.Time
	}
	return nil
}
