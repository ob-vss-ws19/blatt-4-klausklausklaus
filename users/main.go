package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	protores "github.com/ob-vss-ws19/blatt-4-klausklausklaus/reservation/proto"
	proto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/users/proto"
	us "github.com/ob-vss-ws19/blatt-4-klausklausklaus/users/users"
)

/*
Main Function to start a new users service.
*/
func main() {
	service := micro.NewService(micro.Name("users"))
	service.Init()
	newUserService := us.CreateNewUserHandleInstance()
	newUserService.AddDependency(&us.Dependencies{
		ResService: func() protores.ReservationService {
			return protores.NewReservationService("reservation", service.Client())
		},
	})
	err1 := proto.RegisterUsersHandler(service.Server(), newUserService)

	if err1 == nil {
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err1)
	}
}
