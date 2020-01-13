package show_test

import (
	"context"
	"fmt"
	"testing"

	showtestproto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/show/proto"
	"github.com/ob-vss-ws19/blatt-4-klausklausklaus/show/show"
)

/*
TestCreateShow will be a testcase for adding show into the service.
*/
func TestCreateShow(t *testing.T) {
	service := show.NewShowPool()
	response := showtestproto.CreateShowResponse{}
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &response)

	if err == nil {
		if response.CreateShowId < 0 {
			t.Errorf("Cannot create a show with the id %d", response.CreateShowId)
		} else {
			t.Log("Creating a Show will work.")
		}
	} else {
		fmt.Println(err)
	}
}

/*
TestDeleteShow will be a testcase for deleting a show from the service.
*/
func TestDeleteShow(t *testing.T) {
	service := show.NewShowPool()
	response := showtestproto.CreateShowResponse{}
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &response)
	responseDelete := showtestproto.DeleteShowResponse{}
	err1 := service.DeleteShow(context.TODO(), &showtestproto.DeleteShowRequest{DeleteShowId: response.CreateShowId}, &responseDelete)
	if err == nil && err1 == nil {
		if !responseDelete.Answer {
			t.Errorf("cannot delete the show with the id %d", response.CreateShowId)
		} else {
			t.Log("Deleting a Show will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
TestDeleteShowConnectedMove will be a testcase for deleting will be a testcase for deleting all shows with a specific movieId and cinemaid.

func TestDeleteShowMovie(t *testing.T) {
	service := show.NewShowPool()
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	err1 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	err2 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	err3 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	responseDeleteMovie := showtestproto.DeleteShowMovieResponse{}
	err4 := service.DeleteShowMovie(context.TODO(), &showtestproto.DeleteShowMovieRequest{MovieId: 2}, &responseDeleteMovie)
	responseDeleteCinema := showtestproto.DeleteShowCinemaResponse{}
	err5 := service.DeleteShowCinema(context.TODO(), &showtestproto.DeleteShowCinemaRequest{CinemaId: 1}, &responseDeleteCinema)

	if err == nil && err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		switch {
		case !responseDeleteMovie.Answer:
			t.Errorf("cannot delete shows with the moveid %d", 2)
		case !responseDeleteCinema.Answer:
			t.Errorf("cannot delete shows with the cinemaid %d", 1)
		default:
			t.Log("Deleting a show with the cinemaid and movieid will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		fmt.Println(err4)
		fmt.Println(err5)
	}
}
*/

/*
TestListShow will be a testcase for listing all shows.
*/
func TestListShow(t *testing.T) {
	service := show.NewShowPool()
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	err1 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	err2 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	err3 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	responseList := showtestproto.ListShowResponse{}
	err4 := service.ListShow(context.TODO(), &showtestproto.ListShowRequest{}, &responseList)
	if err == nil && err1 == nil && err2 == nil && err3 == nil && err4 == nil {
		if len(responseList.AllShowsData) != 4 {
			t.Errorf("cannot list all shows")
		} else {
			t.Log("listing all will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		fmt.Println(err4)
	}
}

/*
TestFindShowCinema will be a testcase for finding all shows with a specific movieId and cinemaid.
*/
func TestFindShowCinema(t *testing.T) {
	service := show.NewShowPool()
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	err1 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	err2 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	err3 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	responseFindMovie := showtestproto.FindShowMovieResponse{}
	err4 := service.FindShowMovie(context.TODO(), &showtestproto.FindShowMovieRequest{MovieId: 2}, &responseFindMovie)
	responseFindCinema := showtestproto.FindShowCinemaResponse{}
	err5 := service.FindShowCinema(context.TODO(), &showtestproto.FindShowCinemaRequest{CinemaId: 1}, &responseFindCinema)

	if err == nil && err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		switch {
		case len(responseFindMovie.MovieData) != 2:
			t.Errorf("cannot find all shows with movieid: %d", 2)
		case len(responseFindCinema.CinemaData) != 3:
			t.Errorf("cannot find all shows with cinemaid: %d", 1)
		default:
			t.Log("Finding all shows with a movieid and cinemaid will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		fmt.Println(err4)
		fmt.Println(err5)
	}
}
