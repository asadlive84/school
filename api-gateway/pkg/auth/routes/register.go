package routes

import (
	"context"
	"net/http"

	pb "github.com/asadlive84/school/pb/auth"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func Register(ctx *gin.Context, d pb.AuthServiceClient) {

	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := d.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
		Phone:    body.Phone,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}
