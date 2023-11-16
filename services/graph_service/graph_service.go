package followservice

import (
	"encoding/json"
	"go_twitter_clone/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	GraphEndpoint = "graph"
)

var Graphs = []models.Graph{}

type GraphPayload struct {
	FolloweeID uint64 `json:"followee_id"`
}

const CurrentUserID uint64 = 23 // This is the actual user id who is making the request

func Follow(ctx *gin.Context) {
	payloadJSON, err := ctx.GetRawData()
	payload := GraphPayload{}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(payloadJSON, &payload)
	if err != nil {
		log.Fatal(err)
	}
	newGraph := models.Graph{
		UserID:     CurrentUserID,
		FolloweeID: payload.FolloweeID,
		CreatedAt:  "2023-02-04", // TODO: take the current time and assign it
	}

	Graphs = append(Graphs, newGraph)
	ctx.IndentedJSON(http.StatusCreated, Graphs)
}

func GetFollowers() []uint64 {
	followerIDs := []uint64{}
	for _, graph := range Graphs {
		if graph.FolloweeID == CurrentUserID {
			followerIDs = append(followerIDs, graph.UserID)
		}
	}
	return followerIDs
}
