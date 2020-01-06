// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: movies.proto

package movieproto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"

	context "context"

	client "github.com/micro/go-micro/client"

	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Movies service

type MoviesService interface {
	// Create a movie.
	CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...client.CallOption) (*CreatedMovieResponse, error)
	// Change a movie.
	ChangeMovie(ctx context.Context, in *ChangeMovieRequest, opts ...client.CallOption) (*ChangeMovieResponse, error)
	// Delete a Movie.
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...client.CallOption) (*DeleteMovieResponse, error)
	// Find will find a movie
	FindMovie(ctx context.Context, in *FindMovieRequest, opts ...client.CallOption) (*FindMovieResponse, error)
	// Stream
	StreamMovie(ctx context.Context, in *StreamMovieRequest, opts ...client.CallOption) (*StreamMovieResponse, error)
}

type moviesService struct {
	c    client.Client
	name string
}

func NewMoviesService(name string, c client.Client) MoviesService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "movieproto"
	}
	return &moviesService{
		c:    c,
		name: name,
	}
}

func (c *moviesService) CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...client.CallOption) (*CreatedMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.CreateMovie", in)
	out := new(CreatedMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) ChangeMovie(ctx context.Context, in *ChangeMovieRequest, opts ...client.CallOption) (*ChangeMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.ChangeMovie", in)
	out := new(ChangeMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...client.CallOption) (*DeleteMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.DeleteMovie", in)
	out := new(DeleteMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) FindMovie(ctx context.Context, in *FindMovieRequest, opts ...client.CallOption) (*FindMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.FindMovie", in)
	out := new(FindMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesService) StreamMovie(ctx context.Context, in *StreamMovieRequest, opts ...client.CallOption) (*StreamMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movies.StreamMovie", in)
	out := new(StreamMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Movies service

type MoviesHandler interface {
	// Create a movie.
	CreateMovie(context.Context, *CreateMovieRequest, *CreatedMovieResponse) error
	// Change a movie.
	ChangeMovie(context.Context, *ChangeMovieRequest, *ChangeMovieResponse) error
	// Delete a Movie.
	DeleteMovie(context.Context, *DeleteMovieRequest, *DeleteMovieResponse) error
	// Find will find a movie
	FindMovie(context.Context, *FindMovieRequest, *FindMovieResponse) error
	// Stream
	StreamMovie(context.Context, *StreamMovieRequest, *StreamMovieResponse) error
}

func RegisterMoviesHandler(s server.Server, hdlr MoviesHandler, opts ...server.HandlerOption) error {
	type movies interface {
		CreateMovie(ctx context.Context, in *CreateMovieRequest, out *CreatedMovieResponse) error
		ChangeMovie(ctx context.Context, in *ChangeMovieRequest, out *ChangeMovieResponse) error
		DeleteMovie(ctx context.Context, in *DeleteMovieRequest, out *DeleteMovieResponse) error
		FindMovie(ctx context.Context, in *FindMovieRequest, out *FindMovieResponse) error
		StreamMovie(ctx context.Context, in *StreamMovieRequest, out *StreamMovieResponse) error
	}
	type Movies struct {
		movies
	}
	h := &moviesHandler{hdlr}
	return s.Handle(s.NewHandler(&Movies{h}, opts...))
}

type moviesHandler struct {
	MoviesHandler
}

func (h *moviesHandler) CreateMovie(ctx context.Context, in *CreateMovieRequest, out *CreatedMovieResponse) error {
	return h.MoviesHandler.CreateMovie(ctx, in, out)
}

func (h *moviesHandler) ChangeMovie(ctx context.Context, in *ChangeMovieRequest, out *ChangeMovieResponse) error {
	return h.MoviesHandler.ChangeMovie(ctx, in, out)
}

func (h *moviesHandler) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, out *DeleteMovieResponse) error {
	return h.MoviesHandler.DeleteMovie(ctx, in, out)
}

func (h *moviesHandler) FindMovie(ctx context.Context, in *FindMovieRequest, out *FindMovieResponse) error {
	return h.MoviesHandler.FindMovie(ctx, in, out)
}

func (h *moviesHandler) StreamMovie(ctx context.Context, in *StreamMovieRequest, out *StreamMovieResponse) error {
	return h.MoviesHandler.StreamMovie(ctx, in, out)
}