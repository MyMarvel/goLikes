package routes

import (
	"fmt"
	"net/http"

	"likes_handler/controllers"

	"github.com/gin-gonic/gin"
)

var controllersFactory controllers.IFactory

func InitControllersFactory(initFactory controllers.IFactory) {
	controllersFactory = initFactory
}

type doLikeParams struct {
	TargetAccount  string `json:"targetAccount" binding:"required"`
	CurrentAccount string `json:"currentAccount" binding:"required"`
}

type getLikeCountParams struct {
	TargetAccount string `json:"targetAccount" binding:"required"`
}

func GenerateRoutes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1.0")
	{
		eg := v1.Group("/likes")
		{
			eg.POST("/doLike", doLike)
			eg.POST("/getLikeCount", getLikeCount)
		}
	}
	return r
}

// @Summary Adds a like to an account
// @Description Increments likes counter for a person in case the user didn't like them yet
// @Accept  json
// @Produce json
// @Param   Body body doLikeParams true "Params:<ul><li><i>TargetAccount</i>: The account name we like.</li><li><i>CurrentAccount</i>: Our account (the person who likes)</li></ul>" SchemaExample(asdfsdf:asdfff)
// @Success 200 string {object} "Example: <code>{'likes': likesCount}</code>"
// @Failure 208 string {object} "Example: <code>{'warning': 'Already liked'}</code>"
// @Failure 400 string {object} "Example: <code>{'error': 'Error description'}</code>"
// @Tags    Likes
// @Router  /likes/doLike [post]
func doLike(g *gin.Context) {
	var params doLikeParams
	if err := g.ShouldBindJSON(&params); err == nil {
		var acc controllers.IAccount = controllersFactory.NewIAccount(params.CurrentAccount)
		isAccountAlreadyLikes, err := acc.IsAccountLikes(params.TargetAccount)
		if err == nil {
			if !isAccountAlreadyLikes {
				likesCount, err := acc.IncrementLikeCounter(params.TargetAccount)
				if err == nil {
					printLikesCount(likesCount, g)
				} else {
					handleError(err, g)
				}
			} else {
				g.JSON(http.StatusAlreadyReported, gin.H{"warning": "Already liked"})
			}
		} else {
			handleError(err, g)
		}
	} else {
		handleError(err, g)
	}
}

func handleError(e error, g *gin.Context) {
	fmt.Println(http.StatusBadRequest, e.Error())
	g.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
}

func printLikesCount(count int, g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"likes": fmt.Sprintf("%d", count)})
}

// @Summary Get likes for an account
// @Description Returns amount of likes collected by a certain account
// @Accept  json
// @Produce json
// @Param   Body body getLikeCountParams true "Params:<ul><li><i>TargetAccount</i>: The account name we like</li><ul>"
// @Success 200 string {object} "Example: <code>{'likes': likesCount}</code>"
// @Failure 400 string {object} "Example: <code>{'error': 'Error description'}</code>"
// @Tags    Likes
// @Router  /likes/getLikeCount [post]
func getLikeCount(g *gin.Context) {
	var params getLikeCountParams
	if err := g.ShouldBindJSON(&params); err == nil {
		var acc controllers.IAccount = controllersFactory.NewIAccount(params.TargetAccount)
		if likesCount, err := acc.GetLikesCount(params.TargetAccount); err == nil {
			printLikesCount(likesCount, g)
		} else {
			handleError(err, g)
		}
	} else {
		handleError(err, g)
	}
}
