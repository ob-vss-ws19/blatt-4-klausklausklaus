package main

import (
	"fmt"

	reservationproto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/reservation/proto"

	micro "github.com/micro/go-micro"
	showproto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/show/proto"
	show "github.com/ob-vss-ws19/blatt-4-klausklausklaus/show/show"
)

const serviceName = "show"

func main() {
	service := micro.NewService(
		micro.Name(serviceName),
	)
	service.Init()

	newService := show.NewShowPool()
	newService.AddDependency(&show.SServiceDependency{
		ReservationService: func() reservationproto.ReservationService {
			return reservationproto.NewReservationService("reservation", service.Client())
		},
	})

	err1 := showproto.RegisterShowHandler(service.Server(), newService)

	// Run the server
	if err1 == nil {
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err1)
	}
}
