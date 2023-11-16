package tweetservice

import (
	"go_twitter_clone/models"
	"time"
)

type ByTime []models.Tweet

func (b ByTime) Len() int          { return len(b) }
func (b ByTime) Swap(i int, j int) { b[i], b[j] = b[j], b[i] }
func (b ByTime) Less(i int, j int) bool {
	date1, _ := time.Parse("2006-01-02", b[i].CreatedAt)
	date2, _ := time.Parse("2006-01-02", b[j].CreatedAt)
	return date1.After(date2)
}
