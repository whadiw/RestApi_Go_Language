package main

import (
	"go-crud-restapi/controllers/karyawancontroller"
	"go-crud-restapi/database"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	database.ConnectDB()

	api.GET("/api/karyawan", karyawancontroller.GetAll)
	api.GET("/api/karyawan/:id", karyawancontroller.GetById)
	api.POST("/api/karyawan", karyawancontroller.Create)
	api.PUT("/api/karyawan/:id", karyawancontroller.Update)
	api.DELETE("/api/karyawan", karyawancontroller.Delete)

	api.Run()
}
