package cinemahall

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	cinemaproto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/cinemahall/proto"
	showproto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/show/proto"
)

const (
	maxcinemaid = 432342
	one         = 1
)

type CinemaDependency struct {
	ShowService func() showproto.ShowService
}

type CinemaPool struct {
	cinemamap  map[int32]*cinema
	mutex      *sync.Mutex
	dependency *CinemaDependency
}

type seats struct {
	row    int
	column int
}

type cinema struct {
	name    string
	seatmap map[*seats]bool
	row     int32
	colum   int32
}

func (handler *CinemaPool) AddDependency(dep *CinemaDependency) {
	handler.dependency = dep
}

func NewCinemaPool() *CinemaPool {
	newcinema := make(map[int32]*cinema)

	return &CinemaPool{
		mutex:     &sync.Mutex{},
		cinemamap: newcinema,
	}
}

func (handler *CinemaPool) getRandomCinemaID() int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potentialID := int32(rand.Intn(maxcinemaid))
		if !handler.containscinema(potentialID) {
			return potentialID
		}
	}
}

func (handler *CinemaPool) containscinema(id int32) bool {
	_, containsinmap := handler.cinemamap[id]
	return containsinmap
}

func (handler *CinemaPool) containsandreturncinema(id int32) (*cinema, bool) {
	currentcinema, mapcontainscinema := handler.cinemamap[id]
	return currentcinema, mapcontainscinema
}

func (currentcinema *cinema) containsSeatMap(row int32, column int32) bool {
	x := false
	for seat := range currentcinema.seatmap {
		if int32(seat.row) == row && int32(seat.column) == column {
			x = true
		}
	}
	return x
}

func (currentcinema *cinema) getSeat(row int32, column int32) *seats {
	for seat := range currentcinema.seatmap {
		if int32(seat.row) == row && int32(seat.column) == column {
			return seat
		}
	}
	return nil
}

func (handler *CinemaPool) Create(ctx context.Context, request *cinemaproto.CreateCinemaRequest, response *cinemaproto.CreateCinemaResponse) error {
	if len(request.Name) == 0 || request.Row == 0 || request.Column == 0 {
		return errors.New("cinema service - create | Cannot create a cinema with an empty name, zero rows or zero columns")
	}

	newseatmap := map[*seats]bool{}
	for i := int32(one); i <= request.Row; i++ {
		for j := int32(one); j <= request.Column; j++ {
			newseatmap[&seats{row: int(i), column: int(j)}] = false
		}
	}

	createid := handler.getRandomCinemaID()
	handler.mutex.Lock()
	handler.cinemamap[createid] = &cinema{name: request.Name, seatmap: newseatmap, row: request.Row, colum: request.Column}
	handler.mutex.Unlock()
	response.Name = request.Name
	response.Id = createid
	return nil
}

func (handler *CinemaPool) GetSizeOfCinema(ctx context.Context, in *cinemaproto.SizeRequest, out *cinemaproto.SizeResponse) error {
	if handler.containscinema(in.Id) {
		out.Row = handler.cinemamap[in.Id].row
		out.Column = handler.cinemamap[in.Id].colum
		return nil
	}
	return fmt.Errorf("cannot get the size of the cinema by this id %d", in.Id)
}

func (handler *CinemaPool) DeleteShows(ctx context.Context, cinemaid int32) {
	if handler.containscinema(cinemaid) {
		service := handler.dependency.ShowService()
		in := &showproto.DeleteShowCinemaRequest{CinemaId: cinemaid}
		res, err := service.DeleteShowCinema(ctx, in)
		if err != nil && res.Answer {
			fmt.Printf("All shows are deleted for the Cinema %d", cinemaid)
		}
	}
}

func (handler *CinemaPool) Delete(ctx context.Context, request *cinemaproto.DeleteCinemaRequest, response *cinemaproto.DeleteCinemaResponse) error {
	handler.mutex.Lock()
	mapcontainscinema := handler.containscinema(request.Id)
	if !mapcontainscinema {
		handler.mutex.Unlock()
		response.Answer = false
		return fmt.Errorf("cinema service - delete | Cannot delete cinema with id: %d", request.Id)
	}
	delete(handler.cinemamap, request.Id)
	handler.mutex.Unlock()
	response.Answer = true
	return nil
}

func (handler *CinemaPool) Reservation(ctx context.Context, request *cinemaproto.ReservationRequest, response *cinemaproto.ReservationResponse) error {
	handler.mutex.Lock()
	currentcinema, mapcontainscinema := handler.containsandreturncinema(request.Id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		response.Answer = false
		return errors.New("cinema service - reservation | Cannot execute reservation because cinema doesnt exist")
	}
	for _, curreservation := range request.Seatreservation {
		if currentcinema.containsSeatMap(curreservation.Row, curreservation.Column) {
			currentcinema.seatmap[currentcinema.getSeat(curreservation.Row, curreservation.Column)] = true
		}
	}
	handler.cinemamap[request.Id] = currentcinema
	handler.mutex.Unlock()
	response.Answer = true
	return nil
}

func (handler *CinemaPool) Storno(ctx context.Context, request *cinemaproto.StornoRequest, response *cinemaproto.StornoResponse) error {
	handler.mutex.Lock()
	currentcinema, mapcontainscinema := handler.containsandreturncinema(request.Id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		response.Answer = false
		return errors.New("cinema service - storno | Cannot execute storno because cinema doesnt exist")
	}
	for _, curreservation := range request.Seatstorno {
		if currentcinema.containsSeatMap(curreservation.Row, curreservation.Column) {
			currentcinema.seatmap[currentcinema.getSeat(curreservation.Row, curreservation.Column)] = false
		}
	}
	handler.cinemamap[request.Id] = currentcinema
	handler.mutex.Unlock()

	response.Answer = true
	return nil
}

func (handler *CinemaPool) CheckSeats(ctx context.Context, request *cinemaproto.CheckSeatsRequest, response *cinemaproto.CheckSeatsResponse) error {
	handler.mutex.Lock()
	currentcinema, mapcontainscinema := handler.containsandreturncinema(request.Id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		response.Answer = false
		return errors.New("cinema service - checkSeats | Cannot execute storno because cinema doesnt exist")
	}

	isnotreserved := true

	for _, curseat := range request.Seatcheck {
		if currentcinema.containsSeatMap(curseat.Row, curseat.Column) {
			value := currentcinema.seatmap[currentcinema.getSeat(curseat.Row, curseat.Column)]
			if value {
				isnotreserved = false
			}
		}
	}

	handler.mutex.Unlock()
	response.Answer = isnotreserved
	return nil
}

func (handler *CinemaPool) FreeSeats(ctx context.Context, request *cinemaproto.FreeSeatsRequest, response *cinemaproto.FreeSeatsResponse) error {
	handler.mutex.Lock()
	currentcinema, mapcontainscinema := handler.containsandreturncinema(request.Id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		response.Freeseats = nil
		return errors.New("cinema service - freeSeats | Cannot execute storno because cinema doesnt exist")
	}

	retseatmap := map[*seats]bool{}
	for curseat, checkseat := range currentcinema.seatmap {
		if !checkseat {
			retseatmap[curseat] = checkseat
		}
	}
	handler.mutex.Unlock()

	response.Freeseats = makeSendable(retseatmap)
	return nil
}

func makeSendable(m map[*seats]bool) []*cinemaproto.SeatMessage {
	x := []*cinemaproto.SeatMessage{}
	if len(m) > 0 {
		for k := range m {
			x = append(x, &cinemaproto.SeatMessage{Row: int32(k.row), Column: int32(k.column)})
		}
	}
	return x
}
