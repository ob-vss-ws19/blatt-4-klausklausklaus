package movies

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	proto "github.com/ob-vss-ws19/blatt-4-klausklausklaus/movies/proto"
)

const (
	maxmoviesid int32 = 98765432
)

type Movie struct {
	name string
}

func (mo *Movie) getName() string {
	return mo.name
}

type MovieHandlerService struct {
	movies map[int32]*Movie
	mutex  *sync.Mutex
}

func (m *MovieHandlerService) getRandomMovieID(length int32) int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potantialID := rand.Int31n(length)
		if !m.containsID(potantialID) {
			return potantialID
		}
	}
}

func (m *MovieHandlerService) containsID(id int32) bool {
	_, inMap := (*m.getMoviesMap())[id]
	return inMap
}

func (m *MovieHandlerService) getMoviesMap() *map[int32]*Movie {
	return &m.movies
}

func CreateNewMoviesHandlerInstance() *MovieHandlerService {
	return &MovieHandlerService{
		movies: make(map[int32]*Movie),
		mutex:  &sync.Mutex{},
	}
}

func (m *MovieHandlerService) appendMovie(id int32, movie *Movie) bool {
	if id > 0 && movie != nil {
		m.mutex.Lock()
		(*m.getMoviesMap())[id] = movie
		defer m.mutex.Unlock()
		return true
	}
	return false
}

func (m *MovieHandlerService) CreateMovie(context context.Context, in *proto.CreateMovieRequest, out *proto.CreatedMovieResponse) error {
	if in.GetName() != "" {
		mid := m.getRandomMovieID(maxmoviesid)
		if m.appendMovie(mid, &Movie{name: in.GetName()}) {
			out.Movie = &proto.Movie{Id: mid, Name: in.GetName()}
			return nil
		}
	}
	return fmt.Errorf("cannot create movie with an empty name")
}

func (m *MovieHandlerService) change(id int32, pname string) bool {
	if m.containsID(id) {
		m.mutex.Lock()
		(*m.getMoviesMap())[id] = &Movie{name: pname}
		m.mutex.Unlock()
		return true
	}
	return false
}

func (m *MovieHandlerService) StreamMovie(ctx context.Context, in *proto.StreamMovieRequest, out *proto.StreamMovieResponse) error {
	if len(*m.getMoviesMap()) > 0 {
		movies := []*proto.Movie{}
		for k, v := range *m.getMoviesMap() {
			movies = append(movies, &proto.Movie{Id: k, Name: v.getName()})
		}
		out.Movies = movies
	}
	return fmt.Errorf("there are currently no movies store (Advice: find some customers)")
}

func (m *MovieHandlerService) ChangeMovie(ctx context.Context, in *proto.ChangeMovieRequest, out *proto.ChangeMovieResponse) error {
	if in.Movie.Id > 0 && in.Movie.Name != "" {
		if m.change(in.Movie.Id, in.Movie.Name) {
			out.Changed = true
			return nil
		}
	}
	return fmt.Errorf("cannot change the movie. The movie id or the name are not ok. See: %d %s", in.Movie.Id, in.Movie.Name)
}
func (m *MovieHandlerService) DeleteMovie(ctx context.Context, in *proto.DeleteMovieRequest, out *proto.DeleteMovieResponse) error {
	if m.containsID(in.Id) {
		m.mutex.Lock()
		delete(m.movies, in.Id)
		out.Deleted = true
		m.mutex.Unlock()
		return nil
	}
	return fmt.Errorf("cannot delete movie with the id %d", in.Id)
}

func (m *MovieHandlerService) Find(value interface{}) interface{} {
	switch tp := value.(type) {
	case int32:
		if m.containsID(tp) {
			return (*m.getMoviesMap())[tp].getName()
		}
		return ""
	case string:
		for k, v := range *m.getMoviesMap() {
			if v.getName() == tp {
				return k
			}
		}
		return 0
	default:
		return nil
	}
}

func (m *MovieHandlerService) FindMovie(ctx context.Context, in *proto.FindMovieRequest, out *proto.FindMovieResponse) error {
	switch {
	case m.containsID(in.Movie.Id):
		out.Movie.Name = m.Find(in.Movie.Id).(string)
		out.Movie.Id = in.Movie.Id
		return nil
	case in.Movie.Name != "":
		out.Movie.Name = in.Movie.Name
		out.Movie.Id = m.Find(in.Movie.Name).(int32)
		return nil
	default:
		return fmt.Errorf("cannot find a movie with the given id %d or the name %s", in.Movie.Id, in.Movie.Name)
	}
}
