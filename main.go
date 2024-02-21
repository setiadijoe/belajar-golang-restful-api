package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"programmerzamannow/belajar-golang-dependencies-injection/app"
	"programmerzamannow/belajar-golang-dependencies-injection/controller"
	"programmerzamannow/belajar-golang-dependencies-injection/helper"
	"programmerzamannow/belajar-golang-dependencies-injection/middleware"
	"programmerzamannow/belajar-golang-dependencies-injection/repository"
	"programmerzamannow/belajar-golang-dependencies-injection/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
