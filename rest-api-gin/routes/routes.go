package routes

import (
	"go-rest-api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.GET("/alunos/:cpf", controllers.BuscaAlunoPorCPF)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PUT("/alunos/:id", controllers.EditaAluno)

	r.Run()
}
