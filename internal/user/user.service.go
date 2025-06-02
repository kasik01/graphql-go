package user

type Service struct {
	repository *Repository
}

func (s *Service) Register(payload RegisterUserRequest) (*RegisterUserResponse, error) {
	user, err := s.repository.FindByUsername(payload.UserName)
	if err == nil && user != nil {
		return nil, &Errors.DuplicatedUsername
	}

	user, err = NewUser(payload)

	if err != nil {
		return nil, err
	}

	user, err = s.repository.Save(user)

	return &RegisterUserResponse{
		UserName: user.Username,
	}, err
}

func (s *Service) FindByUsername(username string) (*User, error) {
	return s.repository.FindByUsername(username)
}

func NewService() *Service {
	repository := NewRepository()
	service := Service{repository}

	return &service
}
