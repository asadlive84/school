package routes

import (
	// "context"

	"net/http"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
)

func GetDistrict(ctx *gin.Context, d pb.StudentServiceClient) {

	res, err := d.GetDistrictList(ctx, &pb.GetDistrictListRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"message": res.District,
	}

	// Send JSON response
	ctx.JSON(http.StatusOK, data)

}
