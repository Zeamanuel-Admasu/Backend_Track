package concurrency

import (
	"fmt"
	"library_management/models"
	"library_management/services"
	"time"
)

func StartReservationWorker(lib *services.Library) {
	go func() {
		for req := range lib.ReservationChannel() {
			go func(req models.ReservationRequest) {
				time.Sleep(10 * time.Second)

				lib.Mutex().Lock()
				defer lib.Mutex().Unlock()

				book, exists := lib.Books()[req.BookID]
				if !exists {
					return
				}
				if book.Status == "Reserved" && book.ReservedBy == req.MemberID {
					book.Status = "Available"
					book.ReservedBy = 0
					fmt.Printf("Reservation for Book ID %d by Member %d expired\n", req.BookID, req.MemberID)
				}
			}(req)
		}
	}()
}
