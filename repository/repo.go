package repository

import (
	"errors"
	"gorm.io/gorm"
	"graphql/models"
)

type Repo struct {
	DB *gorm.DB
}
type UserRepo interface {
	CreateUser(userDetails models.User) (models.User, error)
	CheckEmail(email string) (models.User, error)
	CreateCompany(companyDetails models.Company) (models.Company, error)
	ViewAllCompany() ([]models.Company, error)
	ViewCompanyByID(cid string) (models.Company, error)
	CreateJob(JobDetails models.Job) (models.Job, error)
	ViewJobById(id string) (models.Job, error)
	ViewAllJob() ([]models.Job, error)
	ViewJobByCid(cid string) ([]models.Job, error)
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
