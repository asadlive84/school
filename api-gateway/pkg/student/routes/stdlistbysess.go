package routes

import (
	// "context"

	"net/http"
	"strconv"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
)

type GetStudentListBySessionRequestBody struct {
	SessionId string `json:"sessionId,omitempty"`
}

func GetStudentListBySession(ctx *gin.Context, d pb.StudentServiceClient) {

	body := GetStudentListBySessionRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ssId, _ := strconv.Atoi(body.SessionId)

	res, err := d.GetStudentListBySession(ctx, &pb.GetStudentListBySessionRequest{
		SessionId: int32(ssId),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"message": res.Student,
	}

	// Send JSON response
	ctx.JSON(http.StatusOK, data)

}
