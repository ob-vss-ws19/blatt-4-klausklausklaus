package cinemahall_test

import (
	"context"
	"fmt"
	"github.com/ob-vss-ws19/blatt-4-klausklausklaus/cinemahall/cinemahall"
	cinemaprototest "github.com/ob-vss-ws19/blatt-4-klausklausklaus/cinemahall/proto"
	"testing"
)

const (
	RowValue    = 5
	ColumnValue = 5
	one         = 1
)

func TestDelete(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()

	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: RowValue, Column: ColumnValue}, &response)
	responseDelete := cinemaprototest.DeleteCinemaResponse{}
	err1 := service.Delete(context.TODO(), &cinemaprototest.DeleteCinemaRequest{Id: response.Id}, &responseDelete)
	if err == nil && err1 == nil {
		if !responseDelete.Answer {
			t.Errorf("Cannot delete the cinema with the namide %d", 1)
		} else {
			t.Log("Deleting a Cinema will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

func TestCreate(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: RowValue, Column: ColumnValue}, &response)
	if err == nil {
		switch {
		case response.Name != "C1":
			t.Errorf("Cannot create a cinema with the name %s", TestName)
		case response.Id < 0:
			t.Fatal("Cannot create a cinema with a proper ID")
		default:
			t.Log("Creating a Cinema will work.")
		}
	} else {
		fmt.Println(err)
	}
}

func TestReservation(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: RowValue, Column: RowValue}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	row := int32(one)
	col := int32(one)
	x = append(x, &cinemaprototest.SeatMessage{Row: row, Column: col})
	err1 := service.Reservation(context.TODO(), &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)

	if err == nil && err1 == nil {
		if !responseReservation.Answer {
			t.Error("Reservation failed")
		} else {
			t.Log("Reservation for a seat will work in a cinema .")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

func TestStorno(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: RowValue, Column: ColumnValue}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	row := int32(one)
	col := int32(one)
	x = append(x, &cinemaprototest.SeatMessage{Row: row, Column: col})
	err1 := service.Reservation(context.TODO(), &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)
	responseStorno := cinemaprototest.StornoResponse{}
	err2 := service.Storno(context.TODO(), &cinemaprototest.StornoRequest{Id: response.Id, Seatstorno: x}, &responseStorno)

	if err == nil && err1 == nil && err2 == nil {
		if !responseStorno.Answer {
			t.Error("Storno failed")
		} else {
			t.Log("Storno will work in a cinema.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}

func TestCheckSeats(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: RowValue, Column: ColumnValue}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	row := int32(one)
	col := int32(one)
	x = append(x, &cinemaprototest.SeatMessage{Row: row, Column: col})
	err1 := service.Reservation(context.TODO(), &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)
	responseCheckSeats := cinemaprototest.CheckSeatsResponse{}
	err2 := service.CheckSeats(context.TODO(), &cinemaprototest.CheckSeatsRequest{Id: response.Id, Seatcheck: x}, &responseCheckSeats)

	if err == nil && err1 == nil && err2 == nil {
		if responseCheckSeats.Answer {
			t.Error("CheckSeats failed")
		} else {
			t.Log("CheckSeats will work in a cinema .")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}
