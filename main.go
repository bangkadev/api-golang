package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tentangkode/api-golang/controllers/mahasiswacontroller"
	"github.com/tentangkode/api-golang/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/mahasiswa", mahasiswacontroller.Index)
	r.GET("/api/mahasiswa/:id", mahasiswacontroller.Show)
	r.POST("/api/mahasiswa", mahasiswacontroller.Create)
	r.PUT("/api/mahasiswa/:id", mahasiswacontroller.Update)
	r.DELETE("/api/mahasiswa/:id", mahasiswacontroller.Delete)

	r.Run(":8080")
}
