package student

import (
	"github.com/asadlive84/school/api-gateway/pkg/config"
	"github.com/asadlive84/school/api-gateway/pkg/student/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/student")
	routes.POST("/insert", svc.Insert)
	routes.POST("/check", svc.Login)
	routes.POST("/uploadfile", svc.UploadFile)
	routes.GET("/list", svc.StudentList)
	routes.GET("/clslist", svc.GetClassList)
	routes.GET("/profile/:id", svc.GetStudentProfileById)
	routes.POST("/stdlistbysessionid", svc.GetStudentListBySession)

	return svc

}

func (svc *ServiceClient) Insert(ctx *gin.Context) {
	routes.Insert(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) UploadFile(ctx *gin.Context) {
	routes.UploadFile(ctx, svc.Client)
}
func (svc *ServiceClient) StudentList(ctx *gin.Context) {
	routes.StudentList(ctx, svc.Client)
}
func (svc *ServiceClient) GetClassList(ctx *gin.Context) {
	routes.GetClassList(ctx, svc.Client)
}
func (svc *ServiceClient) GetStudentListBySession(ctx *gin.Context) {
	routes.GetStudentListBySession(ctx, svc.Client)
}

func (svc *ServiceClient) GetStudentProfileById(ctx *gin.Context) {
	routes.GetStudentProfileById(ctx, svc.Client)
}
