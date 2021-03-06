package reservation_test

import (
	"context"
	"fmt"
	"testing"

	proto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/reservation/proto"
	res "github.com/ob-vss-ws19/blatt-4-klausklausklaus/reservation/reservation"
)

const (
	userID = 23
	showID = 34
)

/*
Test to add a potantialreservation.
*/
func TestAddReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	in := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User: userID,
			Show: showID,
			Seats: []*proto.Seat{
				{Seat: userID},
				{Seat: showID}},
		},
	}
	out := &proto.MakeReservationResponse{}
	err := re.MakeReservation(context.TODO(), in, out)
	if err == nil {
		if out.TmpID < 0 || !out.Works {
			t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
		}
	} else {
		fmt.Println(err)
	}
}

/*
Test to add a potantialreservation accepted it and send the same data again.
Result: First send, stored and accepted second not stored.
*/
func TestAddAcceptReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	in := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  userID,
			Show:  showID,
			Seats: []*proto.Seat{{Seat: userID}, {Seat: showID}},
		},
	}
	out := &proto.MakeReservationResponse{}
	err := re.MakeReservation(context.TODO(), in, out)
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	err1 := re.AcceptReservation(context.TODO(), ina, outa)

	if err == nil && err1 == nil {
		switch {
		case out.TmpID < 0 || !out.Works:
			t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
		case outa.FinalID < 1 && !outa.Taken:
			t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
		default:
			t.Log("test add the same reservation 2 times worked fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
Test Add and accept a Reservation.
Result Potantial good final bad.
*/
func TestAddTroughAwayReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	in := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  userID,
			Show:  showID,
			Seats: []*proto.Seat{{Seat: userID}, {Seat: showID}},
		},
	}
	out := &proto.MakeReservationResponse{}
	err := re.MakeReservation(context.TODO(), in, out)
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: false}
	outa := &proto.AcceptReservationResponse{}
	err1 := re.AcceptReservation(context.TODO(), ina, outa)
	if err == nil && err1 == nil {
		switch {
		case out.TmpID < 0 || !out.Works:
			t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
		case outa.FinalID > 1 && outa.Taken:
			t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
		default:
			t.Log("test add the same reservation 2 times worked fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
Test Add and try to make an error a Reservation.
Result Potantial good final bad.
*/
func TestAddErrorReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	in := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  userID,
			Show:  showID,
			Seats: []*proto.Seat{{Seat: userID}, {Seat: showID}},
		},
	}
	out := &proto.MakeReservationResponse{}
	err := re.MakeReservation(context.TODO(), in, out)
	ina := &proto.AcceptReservationRequest{TmpID: -1, Want: true}
	outa := &proto.AcceptReservationResponse{}
	err1 := re.AcceptReservation(context.TODO(), ina, outa)
	if err == nil && err1 == nil {
		switch {
		case out.TmpID < 0 || !out.Works:
			t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
		case outa.FinalID > 1 && outa.Taken:
			t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
		default:
			t.Log("test add the same reservation 2 times worked fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
Test to add a potantialreservation accepted it and send the same data again.
Result: First send, stored and accepted second not stored.
*/
func TestAddAcceptAddAgainReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	in := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  userID,
			Show:  showID,
			Seats: []*proto.Seat{{Seat: userID}, {Seat: showID}},
		},
	}
	out := &proto.MakeReservationResponse{}
	err := re.MakeReservation(context.TODO(), in, out)
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	err1 := re.AcceptReservation(context.TODO(), ina, outa)
	in2 := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  userID,
			Show:  showID,
			Seats: []*proto.Seat{{Seat: userID}, {Seat: showID}},
		},
	}
	out2 := &proto.MakeReservationResponse{}
	err2 := re.MakeReservation(context.TODO(), in2, out2)

	if err == nil && err1 == nil && err2 == nil {
		switch {
		case out.TmpID < 0 || !out.Works:
			t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
		case out2.TmpID != -1 || out2.Works:
			t.Errorf("second time to send a reservation which should not be executed returned the wrong response: %d and Works: %t", out.TmpID, out.Works)
		case outa.FinalID < 1 && !outa.Taken:
			t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
		default:
			t.Log("test add the same reservation 2 times worked fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}

/*
Test to add a reservation accepted it and delete it.
Result not reservation in the service anymore.
*/
func TestAddAcceptDeleteReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	in := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  userID,
			Show:  showID,
			Seats: []*proto.Seat{{Seat: userID}, {Seat: showID}},
		},
	}
	out := &proto.MakeReservationResponse{}
	err := re.MakeReservation(context.TODO(), in, out)
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	err1 := re.AcceptReservation(context.TODO(), ina, outa)
	din := &proto.DeleteReservationRequest{Id: outa.FinalID}
	dout := &proto.DeleteReservationResponse{}

	err2 := re.DeleteReservation(context.TODO(), din, dout)
	if err == nil && err1 == nil && err2 == nil {
		switch {
		case out.TmpID < 0 || !out.Works:
			t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
		case outa.FinalID < 1 && !outa.Taken:
			t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
		case !dout.Deleted:
			t.Errorf("Cannot delete the reservation with the ID:%d, Taken: %t, Deleted responded: %t", outa.FinalID, outa.Taken, dout.Deleted)
		default:
			t.Log("test add the same reservation 2 times worked fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}

/*
Test to add, accept and find a reservation.
Result: Return reservation.
*/
func TestAddAcceptFindReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	in := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  userID,
			Show:  showID,
			Seats: []*proto.Seat{{Seat: userID}, {Seat: showID}},
		},
	}
	out := &proto.MakeReservationResponse{}
	err := re.MakeReservation(context.TODO(), in, out)
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	err1 := re.AcceptReservation(context.TODO(), ina, outa)
	din := &proto.ShowReservationsRequest{Id: outa.FinalID}
	dout := &proto.ShowReservationsResponse{}

	err2 := re.ShowReservations(context.TODO(), din, dout)
	if err == nil && err1 == nil && err2 == nil {
		switch {
		case out.TmpID < 0 || !out.Works:
			t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
		case outa.FinalID < 1 && !outa.Taken:
			t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
		case dout.Res.ResId < 1 || dout.Res.ResId != outa.FinalID || dout.Res.Show != showID || dout.Res.Show < 1 || dout.Res.User != userID || dout.Res.User < 1 || len(dout.Res.Seats) < 1 || len(dout.Res.Seats) != 2:
			t.Errorf("got the wrong answer back wanted-RES-ID: %d, got: %d", dout.Res.ResId, outa.FinalID)
			t.Errorf("got the wrong answer back wanted-userID: %d, got: %d", dout.Res.User, userID)
			t.Errorf("got the wrong answer back wanted-showID: %d, got: %d", dout.Res.Show, showID)
			t.Errorf("got the wrong answer back wanted-Seats: %d, got: %d", len(dout.Res.Seats), 2)
		default:
			t.Log("test add the same reservation 2 times worked fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}

/*
Test to add a potantialreservation and check if a user has a reservation.
*/
func TestAddCheckHasReservationReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	in := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  userID,
			Show:  showID,
			Seats: []*proto.Seat{{Seat: userID}, {Seat: showID}},
		},
	}
	out := &proto.MakeReservationResponse{}
	err := re.MakeReservation(context.TODO(), in, out)

	fin := &proto.HasReservationsRequest{Res: &proto.Reservation{User: userID}}
	fout := &proto.HasReservationsResponse{}

	err1 := re.HasReservations(context.TODO(), fin, fout)
	if err == nil && err1 == nil {
		switch {
		case out.TmpID < 0 || !out.Works:
			t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
		case fout.Amount < 1 || !fout.Has:
			t.Errorf("got the wrong answer there is a reservation but answer was wrong Amount: %d and Has: %t", fout.Amount, fout.Has)
		default:
			t.Log("worked")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
Test to add a potantialreservation and try to check whether a user has a reservation on an empty service.
*/
func TestAddCheckHasReservationOnEmptyServiceReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	fin := &proto.HasReservationsRequest{Res: &proto.Reservation{User: userID}}
	fout := &proto.HasReservationsResponse{}

	err := re.HasReservations(context.TODO(), fin, fout)

	if err == nil {
		if fout.Amount > 1 || fout.Has {
			t.Errorf("got the wrong answer there is a reservation but answer was wrong Amount: %d and Has: %t", fout.Amount, fout.Has)
		} else {
			t.Log("worked")
		}
	} else {
		fmt.Println(err)
	}
}
