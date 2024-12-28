package main

import (
	"go-SOLID-principle/config"
	"go-SOLID-principle/routes"
)

func main() {
	//inisialiasai Gin
	init := config.Init()
	
	app := routes.Init(init)


	// routes.Routes(router, db)

	//membuat route
	// router.GET("/persons", inDB.GetAllPerson)
	// router.GET("/person/:id", inDB.GetDetailPerson)
	// router.POST("/person", inDB.CreatePerson)
	// router.PUT("/person", inDB.UpdatePerson)
	// router.DELETE("/person/:id", inDB.DeletePerson)
	

	//mulai server dengan port 3000
	app.Run(":4050")
}