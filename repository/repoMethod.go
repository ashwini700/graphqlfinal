package repository

import (
	"errors"

	"graphql/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CheckEmail(email string) (models.User, error) {
	var userDetails models.User
	result := r.DB.Where("email = ?", email).First(&userDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("email not found")
	}
	return userDetails, nil
}

func (r *Repo) ViewJobByCid(cid string) ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find job")
	}
	return jobDetails, nil

}

func (r *Repo) ViewAllJob() ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find job")
	}
	return jobDetails, nil

}
func (r *Repo) ViewJobById(id string) (models.Job, error) {
	var jobData models.Job
	result := r.DB.Where("id = ?", id).First(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not find the company")
	}
	return jobData, nil

}

func (r *Repo) CreateJob(JobDetails models.Job) (models.Job, error) {
	result := r.DB.Create(&JobDetails)
	if result.Error != nil {
		return models.Job{}, errors.New("could not create the records")
	}
	return JobDetails, nil

}

func (r *Repo) CreateUser(userDetails models.User) (models.User, error) {
	result := r.DB.Create(&userDetails)
	if result.Error != nil {
		return models.User{}, errors.New("could not create the records")
	}
	return userDetails, nil
}
func (r *Repo) CreateCompany(companyDetails models.Company) (models.Company, error) {
	result := r.DB.Create(&companyDetails)
	if result.Error != nil {
		return models.Company{}, errors.New("could not create the records")
	}
	return companyDetails, nil
}
func (r *Repo) ViewAllCompany() ([]models.Company, error) {
	var companyDetails []models.Company
	result := r.DB.Find(&companyDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find companires")
	}
	return companyDetails, nil
}
func (r *Repo) ViewCompanyByID(cid string) (models.Company, error) {
	var companyData models.Company
	result := r.DB.Where("id = ?", cid).First(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not find the company")
	}
	return companyData, nil
}
