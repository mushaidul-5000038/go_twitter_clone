package tweetservice

import (
	"encoding/json"
	"errors"
	"go_twitter_clone/models"
	"log"
	"net/http"
	"sort"

	gs "go_twitter_clone/services/graph_service"

	"github.com/gin-gonic/gin"
)

const (
	TweetEndpoint = "tweets"
)

var Tweets = []models.Tweet{
	{Content: "https://twitter.com/", CreatedAt: "2023-01-23", CreatedBy: 23},
}

func Post(ctx *gin.Context) {
	payloadJSON, err := ctx.GetRawData()
	newTweet := models.Tweet{}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(payloadJSON, &newTweet)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: make a separate payload type and handle the creation time logic here

	isCorrectLength := verifyContentLength(newTweet.Content)

	if !isCorrectLength {
		log.Fatal(errors.New("incorrent content length"))
	}

	Tweets = append(Tweets, newTweet)
	ctx.IndentedJSON(http.StatusCreated, Tweets)
}

func List(ctx *gin.Context) {
	myTweets := getAllPersonalTweets()
	sort.Sort(ByTime(myTweets))
	ctx.IndentedJSON(http.StatusOK, Tweets)
}

func ListByFollower(ctx *gin.Context) {
	myFollowerIDs := gs.GetFollowers()
	myFollowerTweets := getAllTweetsFromFollowers(myFollowerIDs)
	sort.Sort(ByTime(myFollowerTweets))
	ctx.IndentedJSON(http.StatusOK, Tweets)
}

func getAllPersonalTweets() []models.Tweet {
	filteredTweets := []models.Tweet{}
	for _, tweet := range Tweets {
		if tweet.CreatedBy == gs.CurrentUserID {
			filteredTweets = append(filteredTweets, tweet)
		}
	}
	return filteredTweets
}

func getAllTweetsFromFollowers(followerIDs []uint64) []models.Tweet {
	filteredTweets := []models.Tweet{}
	followerMap := getFollowerMap(followerIDs)
	for _, tweet := range Tweets {
		if followerMap[tweet.CreatedBy] {
			filteredTweets = append(filteredTweets, tweet)
		}
	}
	return filteredTweets
}

func getFollowerMap(followerIDs []uint64) map[uint64]bool {
	followerMap := make(map[uint64]bool)
	for _, followerID := range followerIDs {
		followerMap[followerID] = true
	}
	return followerMap
}

func verifyContentLength(content string) bool {
	// TODO : add length constraint logic herer
	return true
}
