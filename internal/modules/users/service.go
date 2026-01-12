package users

import (
	"context"

	"github.com/eren_dev/go_server/internal/shared/pagination"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type PaginatedUsers struct {
	Data       []*UserResponse       `json:"data"`
	Pagination pagination.Response `json:"pagination"`
}

func (s *Service) Create(ctx context.Context, dto *CreateUserDTO) (*UserResponse, error) {
	user, err := s.repo.Create(ctx, dto)
	if err != nil {
		return nil, err
	}
	return ToResponse(user), nil
}

func (s *Service) FindAll(ctx context.Context, params pagination.Params) (*PaginatedUsers, error) {
	users, total, err := s.repo.FindAll(ctx, params)
	if err != nil {
		return nil, err
	}

	return &PaginatedUsers{
		Data:       ToResponseList(users),
		Pagination: pagination.NewResponse(params, total),
	}, nil
}

func (s *Service) FindByID(ctx context.Context, id string) (*UserResponse, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToResponse(user), nil
}

func (s *Service) Update(ctx context.Context, id string, dto *UpdateUserDTO) (*UserResponse, error) {
	user, err := s.repo.Update(ctx, id, dto)
	if err != nil {
		return nil, err
	}
	return ToResponse(user), nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
