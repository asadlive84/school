package routes

import (
	// "context"
	"net/http"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
)

func GetClassList(ctx *gin.Context, d pb.StudentServiceClient) {

	res, err := d.GetClassList(ctx, &pb.GetClassListRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"data": res.Session,
	}

	// Send JSON response
	ctx.JSON(http.StatusOK, data)

}
