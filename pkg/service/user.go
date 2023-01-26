package service

import (
	"context"
	"doduykhang/hermes-conversation/internal/db/mysql"
	"doduykhang/hermes-conversation/pkg/dto"
	"log"

	dtos "github.com/dranikpg/dto-mapper"
	"github.com/google/uuid"
)

type User interface {
	CreateUser(request *dto.CreateUser) (*dto.UserDTO, error)
}

type user struct {
	queries *mysql.Queries
	mapper *dtos.Mapper
}

func NewUser(queries *mysql.Queries, mapper *dtos.Mapper) User {
	return &user{
		queries: queries,
		mapper: mapper,
	}
}

func (s *user) CreateUser(request *dto.CreateUser) (*dto.UserDTO, error) {
	var args mysql.CreateUserParams
	err := s.mapper.Map(&args, request)

	if err != nil {
		log.Printf("Error at service.user.CreateUser, %s\n", err)
		return nil, err
	}

	id := uuid.New()
	args.ID = id.String()
	_, err = s.queries.CreateUser(context.Background(), args)
	if err != nil {
		log.Printf("Error at service.user.CreateUser, %s\n", err)
		return nil, err
	}

	var dto dto.UserDTO
	dto.ID = args.ID
	err = s.mapper.Map(&dto, &args)

	if err != nil {
		log.Printf("Error at service.user.CreateUser, %s\n", err)
		return nil, err
	}

	return &dto, nil
}

