package app

import (
	"fmt"
	"log"
	"os"

	"github.com/adhikag24/go-brick/app/controllers"
	"github.com/adhikag24/go-brick/app/repository"
	"github.com/adhikag24/go-brick/app/routes"
	"github.com/adhikag24/go-brick/app/usecase"
)

func StartApp() {
	repo := repository.NewRepo()
	usecase := usecase.NewUC(repo)
	ctrl := controllers.NewControllers(usecase)
	router := routes.NewRouter(ctrl)
	app := router.Router()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if err := app.Run(port); err != nil {
		log.Fatalf("Error when start application: %s", err)
	}
}
