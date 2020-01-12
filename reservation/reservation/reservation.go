package reservation

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	protocin "github.com/ob-vss-ws19/blatt-4-klausklausklaus/cinemahall/proto"
	proto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/reservation/proto"
	showproto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/show/proto"
)

/*
containsID will check whether generated ID is already set or not.
@id will be a int32 id to check for.
*/
func (r *ReservatServiceHandler) containsID(id int32) bool {
	_, inMap := (*r.getReservationsMap())[id]
	return inMap
}

/*
containsPotantialReservations will check whether a given id belongs to the unaccepted reservationsl.
*/
func (r *ReservatServiceHandler) containsPotantialReservations(id int32) bool {
	_, inMap := (*r.getPotantialReservationsMap())[id]
	return inMap
}

/*
Get the Map with all potantial reservations.
*/
func (r *ReservatServiceHandler) getPotantialReservationsMap() *map[int32]*Reservation {
	return &r.unaccepted
}

/*
Function to produce a random reservationsID.
@param length will be the length of the number.
@param seed will be a seed in order to produce "really" random numbers.
*/
func (r *ReservatServiceHandler) getRandomReservationsID(length int32) int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potantialID := rand.Int31n(length)
		if !r.containsID(potantialID) && !r.containsPotantialReservations(potantialID) {
			return potantialID
		}
	}
}

/*
getReservationsMap will return a pointer to the current reservations map in order to work in that.
*/
func (r *ReservatServiceHandler) getReservationsMap() *map[int32]*Reservation {
	return &r.reservations
}

/*
Reservation will be the representation of a reservation.
*/
type Reservation struct {
	UserID int32
	ShowID int32
	Seats  []int32
}

/*
ReservationsDependency all dependencys for the reservation service.
*/
type ReservationsDependency struct {
	Cinemahall func() protocin.CinemaService
	Show       func() showproto.ShowService
}

type ReservatServiceHandler struct {
	reservations map[int32]*Reservation
	unaccepted   map[int32]*Reservation
	dependencies *ReservationsDependency
	mutex        *sync.Mutex
}

func CreateNewReservationHandlerInstance() *ReservatServiceHandler {
	return &ReservatServiceHandler{
		reservations: map[int32]*Reservation{},
		unaccepted:   map[int32]*Reservation{},
		mutex:        &sync.Mutex{},
	}
}

func (r *ReservatServiceHandler) AddDependencyRes(dep *ReservationsDependency) {
	r.dependencies = dep
}

func convertSeats(seats []*proto.Seat) []int32 {
	newSeats := []int32{}
	if len(seats) > 0 {
		for _, i := range seats {
			newSeats = append(newSeats, i.Seat)
		}
	}
	return newSeats
}

func (r *ReservatServiceHandler) checkIfSeatIsTaken(seat int32) bool {
	for _, v := range *r.getReservationsMap() {
		for _, resSeat := range v.Seats {
			if resSeat == seat {
				return true
			}
		}
	}
	return false
}

func (r *ReservatServiceHandler) checkIfSeatsStillFree(seats *[]int32) bool {
	for _, seat := range *seats {
		if taken := r.checkIfSeatIsTaken(seat); taken {
			return false
		}
	}
	return true
}

func (r *ReservatServiceHandler) addInReservationsMap(id int32, reserve Reservation) bool {
	if reserve.UserID > 0 && reserve.ShowID > 0 && len(reserve.Seats) > 0 && r.checkIfSeatsStillFree(&reserve.Seats) {
		(*r.getPotantialReservationsMap())[id] = &reserve
		return true
	}
	fmt.Println(fmt.Errorf("cannot add a potantial reservation because the given reservation or id were no as expected userID: %d, showID: %d, Seats-len: %d, ReservationID: %d", reserve.UserID, reserve.ShowID, len(reserve.Seats), id))
	return false
}

func (r *ReservatServiceHandler) makePotentialReservation(seats []*proto.Seat, userid, showid int32) (bool, int32) {
	seat := convertSeats(seats)
	id := r.getRandomReservationsID(91734628)
	if added := r.addInReservationsMap(id, Reservation{
		UserID: userid,
		ShowID: showid,
		Seats:  seat,
	}); added {
		return true, id
	}
	return false, -1
}

func (r *ReservatServiceHandler) MakeReservation(ctx context.Context, in *proto.MakeReservationRequest, out *proto.MakeReservationResponse) error {
	if len(in.Res.Seats) > 0 && in.Res.Show > 0 && in.Res.User > 0 {
		r.mutex.Lock()
		if b, id := r.makePotentialReservation(in.Res.Seats, in.Res.User, in.Res.Show); b {
			r.mutex.Unlock()
			out.TmpID = id
			out.Works = true
			return nil
		}
		r.mutex.Unlock()
		out.TmpID = -1
		out.Works = false
		return nil
	}
	return fmt.Errorf("cannot create a reservation with an list of size: %d", len(in.Res.Seats))
}

func (r *ReservatServiceHandler) swapValuesBetweenMaps(id int32) bool {
	if r.containsPotantialReservations(id) {
		(*r.getReservationsMap())[id] = (*r.getPotantialReservationsMap())[id]
		delete(*r.getPotantialReservationsMap(), id)
		return true
	}
	return false
}

func (r *ReservatServiceHandler) convertSeatsToCinemaHallSeats(ctx context.Context, seat int32) (int32, int32) {
	service := r.dependencies.Cinemahall()
	in := &protocin.SizeRequest{}
	resp, err := service.GetSizeOfCinema(ctx, in)
	if err == nil && resp != nil {
		reihe := seat / resp.Column
		spalte := seat % resp.Column
		if spalte == 0 {
			spalte = resp.Column
		}
		return reihe, spalte
	}
	return int32(-1), int32(-1)
}

func (r *ReservatServiceHandler) makeSeatsCinemaHallSeats(context context.Context, id int32, cinema *protocin.SizeResponse) *[]*protocin.SeatMessage {
	seats := []*protocin.SeatMessage{}
	if r.containsPotantialReservations(id) {
		for _, v := range (*r.getPotantialReservationsMap())[id].Seats {
			row, column := r.convertSeatsToCinemaHallSeats(context, v)
			seats = append(seats, &protocin.SeatMessage{Row: row, Column: column})
		}
		return &seats
	}
	return &seats
}

func (r *ReservatServiceHandler) AcceptReservation(ctx context.Context, in *proto.AcceptReservationRequest, out *proto.AcceptReservationResponse) error {
	r.mutex.Lock()
	switch {
	case in.TmpID > 0 && in.Want && r.containsPotantialReservations(in.TmpID) && r.checkIfSeatsStillFree(&(*r.getPotantialReservationsMap())[in.TmpID].Seats):
		moviehandler := r.dependencies.Show()
		service := r.dependencies.Cinemahall()
		var (
			cinemaid int32
			seats    *[]*protocin.SeatMessage
			cinin    *protocin.ReservationRequest
		)
		response, err := moviehandler.ListShow(ctx, &showproto.ListShowRequest{})
		if err != nil {
			fmt.Println(response)
			fmt.Println(err)
			fmt.Println("there is an error while accepting a reservation")
		}
		for k, v := range response.ShowId {
			if v == (*r.getPotantialReservationsMap())[in.TmpID].ShowID {
				cinemaid = response.AllShowsData[k].CinemaId
			}
		}
		if swapped := r.swapValuesBetweenMaps(in.TmpID); !swapped {
			r.mutex.Unlock()
			return fmt.Errorf("cannot make the potantial reservation a actuall reservation id: %d (invalid id)", in.TmpID)
		}
		if cinemaid > 0 {
			inSize := &protocin.SizeRequest{Id: cinemaid}
			size, err := service.GetSizeOfCinema(ctx, inSize)
			if err == nil {
				seats = r.makeSeatsCinemaHallSeats(ctx, in.TmpID, size)
				cinin = &protocin.ReservationRequest{Id: cinemaid, Seatreservation: *seats}
			} else {
				fmt.Printf("cannot find a Cinema with the given id %d, Error: %e \n", cinemaid, err)
			}
		}
		responseReservationRequest, err2 := service.Reservation(ctx, cinin)
		if cinemaid > 0 && err2 == nil && responseReservationRequest.Answer {
			out.FinalID = in.TmpID
			out.Taken = true
			r.mutex.Unlock()
			return nil
		}
		r.mutex.Unlock()
		return fmt.Errorf("cannot accept or delete a reservation with the id: %d (invalid id)", in.TmpID)
	case r.containsPotantialReservations(in.TmpID) && !in.Want:
		r.mutex.Lock()
		delete(*r.getPotantialReservationsMap(), in.TmpID)
		r.mutex.Unlock()
		out.FinalID = -1
		out.Taken = false
		r.mutex.Unlock()
		return nil
	default:
		out.FinalID = -1
		out.Taken = false
		r.mutex.Unlock()
		return fmt.Errorf("cannot accept the reservation because the requested seats are already taken")
	}
}

func (r *ReservatServiceHandler) rdelete(id int32) bool {
	if r.containsID(id) {
		r.mutex.Lock()
		delete(*r.getReservationsMap(), id)
		r.mutex.Unlock()
		return true
	}
	return false
}

func (r *ReservatServiceHandler) DeleteReservationByShowID(ctx context.Context, showID int32) {
	if showID > -1 {
		for resID, reservation := range *r.getReservationsMap() {
			if reservation.ShowID == showID {
				r.rdelete(resID)
			}
		}
	}
}

func (r *ReservatServiceHandler) DeleteReservation(ctx context.Context, in *proto.DeleteReservationRequest, out *proto.DeleteReservationResponse) error {
	switch {
	case r.containsID(in.Id):
		r.rdelete(in.Id)
		out.Deleted = true
		return nil
	case in.ShowId > -1:
		r.DeleteReservationByShowID(ctx, in.ShowId)
		out.Deleted = true
		return nil
	default:
		out.Deleted = false
		return fmt.Errorf("cannot find a entry with the id %d --> cannot delete this", in.Id)
	}
}

func (r *ReservatServiceHandler) changeReservation(id int32, in proto.Reservation) bool {
	if r.containsID(id) {
		r.mutex.Lock()
		(*r.getReservationsMap())[id] = &Reservation{
			UserID: in.User,
			ShowID: in.Show,
			Seats:  convertSeats(in.Seats),
		}
		r.mutex.Unlock()
		return true
	}
	return false
}

func (r *ReservatServiceHandler) ChangeReservation(ctx context.Context, in *proto.ChangeReservationRequest, out *proto.ChangeReservationResponse) error {
	if r.containsID(in.Res.ResId) {
		if changed := r.changeReservation(in.Res.ResId, *in.Res); changed {
			out.Changed = true
			return nil
		}
	}
	out.Changed = false
	return fmt.Errorf("cannot change the reservation with the id: %d", in.Res.ResId)
}

func (r *ReservatServiceHandler) FindSingleReservation(id int32) *Reservation {
	if (id) > 0 && r.containsID(id) {
		for k, v := range *r.getReservationsMap() {
			if k == id {
				return v
			}
		}
	}
	return nil
}

func makeSeatToProtoSeat(localseats []int32) []*proto.Seat {
	convertedSeats := []*proto.Seat{}
	if len(localseats) > 0 {
		for _, lSeat := range localseats {
			convertedSeats = append(convertedSeats, &proto.Seat{
				Seat: lSeat,
			})
		}
	}
	return convertedSeats
}

func (r *ReservatServiceHandler) ShowReservations(ctx context.Context, in *proto.ShowReservationsRequest, out *proto.ShowReservationsResponse) error {
	if (in.Id) > 0 && r.containsID(in.Id) {
		reservation := r.FindSingleReservation(in.Id)
		out.Res = &proto.Reservation{
			ResId: in.Id,
			User:  reservation.UserID,
			Show:  reservation.ShowID,
			Seats: makeSeatToProtoSeat(reservation.Seats),
		}
		return nil
	}
	return fmt.Errorf("cannot find a user with this id: %d", in.Id)
}

func (r *ReservatServiceHandler) prepareStream() []*proto.Reservation {
	if len(r.reservations) > 0 {
		res := []*proto.Reservation{}
		for k, v := range r.reservations {
			res = append(res, &proto.Reservation{
				ResId: k,
				User:  v.UserID,
				Show:  v.ShowID,
				Seats: makeSeatToProtoSeat(v.Seats),
			})
		}
		return res
	}
	return nil
}

func (r *ReservatServiceHandler) StreamUsersReservations(ctx context.Context, in *proto.StreamUsersReservationsRequest, out *proto.StreamUsersReservationsResponse) error {
	if reservations := r.prepareStream(); len(reservations) > 0 {
		out.Reservations = reservations
		return nil
	}
	out.Reservations = []*proto.Reservation{}
	return fmt.Errorf("there is noting yet we could stream")
}

func (r *ReservatServiceHandler) FindUserIfReservation(id int32) (bool, int32, bool) {
	if id > 0 {
		r.mutex.Lock()
		entrys := 0
		potantial := false
		for _, v := range *r.getPotantialReservationsMap() {
			if v.UserID == id {
				entrys++
				potantial = true
			}
		}
		for _, v2 := range *r.getReservationsMap() {
			if v2.UserID == id {
				entrys++
			}
		}
		r.mutex.Unlock()
		return (entrys > 0), int32(entrys), potantial
	}
	return false, -1, false
}

func (r *ReservatServiceHandler) HasReservations(ctx context.Context, in *proto.HasReservationsRequest, out *proto.HasReservationsResponse) error {
	has, howmuch, _ := r.FindUserIfReservation(in.Res.User)
	switch {
	case has && howmuch > 0:
		out.Has = has
		out.Amount = howmuch
		return nil
	case !has && howmuch < 1:
		out.Has = false
		out.Amount = -1
		return nil
	default:
		return fmt.Errorf("an error occurred cannot make sure wether the user is unknown or got a wrong combination of has entry and the amout like (true and -1)")
	}
}
