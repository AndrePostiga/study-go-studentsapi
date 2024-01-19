package api

import (
	"github.com/andrepostiga/api-go-gin/controllers"
	domainservices "github.com/andrepostiga/api-go-gin/domain/domainServices"
	"github.com/andrepostiga/api-go-gin/infrastructure/database"
	"github.com/andrepostiga/api-go-gin/infrastructure/database/repositories"
	externalservices "github.com/andrepostiga/api-go-gin/infrastructure/externalServices"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Router *gin.Engine
}

func StartWebApi() *Routes {
	r := Routes{gin.Default()}
	DB := database.ConnectPostgres()

	registryRequestValidators()

	// clients
	cepClient := externalservices.NewCepClient("https://brasilapi.com.br/api")

	// Repositories
	studentRepository := repositories.NewStudentRepository(DB)

	// Domain Services
	studentDomainService := domainservices.NewCreateStudentService(studentRepository, cepClient)

	// Controllers
	studentCotroller := controllers.NewStudentController(studentRepository, studentDomainService)

	routerGroup := r.Router.Group("/v1")
	{
		studentCotroller.SetupStudentController(routerGroup)
		controllers.SetupPingController(routerGroup)
	}

	return &r
}
