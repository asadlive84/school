package routes

import (
	"fmt"
	"net/http"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
)

func GetVillageOrRoadListByUnionId(ctx *gin.Context, d pb.StudentServiceClient) {
	// Extracting query parameters from the URL
	uiononId := ctx.Param("id")

	if uiononId == "" {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("error"))
		return
	}

	res, err := d.GetVillageOrRoadListByUnionId(ctx, &pb.GetVillageOrRoadListByUnionIdRequest{
		UiononId: uiononId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"message": res.VillageOrRoad,
	}

	// Send JSON response
	ctx.JSON(http.StatusOK, data)
}
