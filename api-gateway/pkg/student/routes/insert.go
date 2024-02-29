package routes

import (
	"context"
	"net/http"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type InsertRequestBody struct {
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

func Insert(ctx *gin.Context, d pb.StudentServiceClient) {

	body := InsertRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := d.Insert(context.Background(), &pb.InsertRequest{
		Student: &pb.Student{
			Id:           body.Id,
			Name:         body.Name,
			NameBn:       body.NameBn,
			StdId:        body.StdId,
			Dob:          body.Dob,
			FathersName:  body.FathersName,
			MothersName:  body.MothersName,
			Session:      body.Session,
			Gender:       body.Gender,
			MobileNumber: body.MobileNumber,
			ClassName:    body.ClassName,
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}
