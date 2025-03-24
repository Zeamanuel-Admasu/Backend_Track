package main

import (
	"library_management/concurrency"
	"library_management/controllers"
	"library_management/services"
)

func main() {
	library := services.NewLibrary()
	concurrency.StartReservationWorker(library)
	controllers.StartLibrarySystem(library)
}
