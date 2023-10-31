package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/tintin007k/rss-agg/internal/database"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := cfg.DB.GetPosts(r.Context(), database.GetPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	log.Printf("Received: %v posts", len(posts))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error in fetching posts")
		return
	}
	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
