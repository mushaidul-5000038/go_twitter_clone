package main

import (
	"fmt"
	gs "go_twitter_clone/services/graph_service"
	ts "go_twitter_clone/services/tweet_service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST(fmt.Sprintf("/%s/post", ts.TweetEndpoint), ts.Post)
	router.GET(fmt.Sprintf("/%s/list", ts.TweetEndpoint), ts.List)
	router.GET(fmt.Sprintf("/%s/list_by_follower", ts.TweetEndpoint), ts.ListByFollower)
	router.POST(fmt.Sprintf("/%s/follow", gs.GraphEndpoint), gs.Follow)
	router.Run()
}
