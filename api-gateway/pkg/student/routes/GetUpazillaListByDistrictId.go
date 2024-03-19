package routes

import (
	"fmt"
	"net/http"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
)

func GetUpazillaListByDistrictId(ctx *gin.Context, d pb.StudentServiceClient) {
	// Extracting query parameters from the URL
	districtId := ctx.Param("id")

	if districtId == "" {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("error"))
		return
	}

	res, err := d.GetUpazillaListByDistrictId(ctx, &pb.GetUpazillaListByDistrictIdRequest{
		DictrictId: districtId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"message": res.Upazilla,
	}

	// Send JSON response
	ctx.JSON(http.StatusOK, data)
}
