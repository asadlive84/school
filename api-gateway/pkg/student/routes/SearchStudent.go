package routes

import (
	// "context"
	"context"
	"fmt"
	"net/http"
	"strconv"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
)

type SearchData struct {
	District     string   `json:"district"`
	Upazillas    []string `json:"upazillas"`
	Unions       []string `json:"unions"`
	VillageRoads []string `json:"villageRoads"`
	SessionIds   []string `json:"sessionIds"`
}

func SearchStudent(ctx *gin.Context, d pb.StudentServiceClient) {

	body := SearchData{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	distId, err := strconv.Atoi(body.District)
	if err != nil {

	}

	upazillIds := make([]int32, 0, 1)
	unionIds := make([]int32, 0, 1)
	villageRoadIds := make([]int32, 0, 1)
	sessionIds := make([]int32, 0, 1)

	for _, up := range body.Upazillas {
		upId, err := strconv.Atoi(up)
		if err != nil {
			continue
		}
		upazillIds = append(upazillIds, int32(upId))
	}

	for _, un := range body.Unions {
		unId, err := strconv.Atoi(un)
		if err != nil {
			continue
		}

		unionIds = append(unionIds, int32(unId))
	}

	for _, vill := range body.VillageRoads {
		villid, err := strconv.Atoi(vill)
		if err != nil {
			continue
		}

		villageRoadIds = append(villageRoadIds, int32(villid))
	}

	fmt.Println("========================================")
	fmt.Printf("sessionId %+v\n", body.SessionIds)
	fmt.Println("========================================")

	for _, sessionId := range body.SessionIds {
		sessionId, err := strconv.Atoi(sessionId)
		if err != nil {
			continue
		}

		sessionIds = append(sessionIds, int32(sessionId))
	}

	res, err := d.SearchStudent(context.Background(), &pb.SearchStudentRequest{
		DistrictId:      int32(distId),
		UpazillaId:      upazillIds,
		UnonId:          unionIds,
		VillageOrRoadId: villageRoadIds,
		SessionId:       sessionIds,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"message": res,
	}

	// Send JSON response
	ctx.JSON(http.StatusOK, data)

}
