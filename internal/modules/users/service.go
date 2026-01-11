package users

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, dto *CreateUserDTO) (*UserResponse, error) {
	user, err := s.repo.Create(ctx, dto)
	if err != nil {
		return nil, err
	}
	return ToResponse(user), nil
}

func (s *Service) FindAll(ctx context.Context) ([]*UserResponse, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return ToResponseList(users), nil
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
