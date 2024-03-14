package routes

import (
	"fmt"
	"net/http"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
)

func GetStudentProfileById(ctx *gin.Context, d pb.StudentServiceClient) {
	// Extracting query parameters from the URL
	studentId := ctx.Param("id")

	if studentId == "" {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("error"))
		return
	}

	res, err := d.GetStudentProfileById(ctx, &pb.GetStudentProfileByIdRequest{
		StudentId: studentId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"message": res.Student,
		"ad": res.District,
	}

	// Send JSON response
	ctx.JSON(http.StatusOK, data)
}
