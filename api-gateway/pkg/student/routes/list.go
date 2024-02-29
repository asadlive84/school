package routes

import (
	// "context"
	"net/http"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type StudentListResp struct {
	Id           string               `json:"id,omitempty"`
	Name         string               `json:"name,omitempty"`
	NameBn       string               `json:"name_bn,omitempty"`
	FathersName  string               `json:"fathers_name,omitempty"`
	MothersName  string               `json:"mothers_name,omitempty"`
	Session      string               `json:"session,omitempty"`
	Gender       string               `json:"gender,omitempty"`
	MobileNumber string               `json:"mobile_number,omitempty"`
	ClassName    string               `json:"class_name,omitempty"`
	StdId        string               `json:"stdId,omitempty"`
	Dob          *timestamp.Timestamp `json:"dob,omitempty"`
}

func StudentList(ctx *gin.Context, d pb.StudentServiceClient) {

	res, err := d.StudentList(ctx, &pb.StudentListRequest{})

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
