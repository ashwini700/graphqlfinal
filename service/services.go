package service

import (
	"errors"
	"graphql/authentication"
	"graphql/graph/model"
	"graphql/repository"
)

type UserService interface {
	UserSignup(userData model.NewUser) (*model.User, error)
	LoginUser(userData model.LoginUser) (*model.Login, error)
	CreateCompany(companyDetails model.NewCompnay) (*model.Company, error)
	ViewAllCompanies() ([]*model.Company, error)
	ViewCompanyById(cid string) (*model.Company, error)
	CreateJob(jobData model.NewJob) (*model.Job, error)
	ViewJobByID(id string) (*model.Job, error)
	ViewAllJob() ([]*model.Job, error)
	ViewJobByCid(cid string) ([]*model.Job, error)
}
type Service struct {
	userRepo repository.UserRepo
	a        *authentication.Auth
}

func NewService(a *authentication.Auth, userRepo repository.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be nil")
	}
	return &Service{
		userRepo: userRepo,
		a:        a,
	}, nil

}
