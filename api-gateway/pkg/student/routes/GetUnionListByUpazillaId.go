package routes

import (
	"fmt"
	"net/http"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
)

func GetUnionListByUpazillaId(ctx *gin.Context, d pb.StudentServiceClient) {
	// Extracting query parameters from the URL
	upazilaId := ctx.Param("id")

	if upazilaId == "" {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("error"))
		return
	}

	res, err := d.GetUnionListByUpazillaId(ctx, &pb.GetUnionListByUpazillaIdRequest{
		UpazillaId: upazilaId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"message": res.UnionInfo,
	}

	// Send JSON response
	ctx.JSON(http.StatusOK, data)
}
