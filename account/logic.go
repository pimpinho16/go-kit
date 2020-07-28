package account

import (
	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
	"context"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewServices(rep Repository, logger log.Logger) {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s Service) CreateUser(ctx context.Context, email string, password string) error {
	logger := log.with(s.logger, "method", "CreateUser")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID: id,
		Email: email,
		password: password,
	}

	if err := s.repository.CreateUser(ctx,User); err ≠ nil{
		level.Error(logger).Log("err",err)
		return " ",err
	}

	logger.Log("create user",id)
	return "Success",nil
}

func (s Service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.with(s.logger, "method", "GetUser")

	email, err :=  s.repository.GetUser(ctx,id);

	if err  ≠ nil{
		level.Error(logger).Log("err",err)
		return " ",err
	}

	logger.Log("get user",id)
	return email,nil
}
