package service

import (
	"errors"
	"graphql/graph/model"
	"graphql/models"
	pkgs "graphql/package"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) LoginUser(userData model.LoginUser) (*model.Login, error) {
	userDetails, err := s.userRepo.CheckEmail(userData.Email)
	if err != nil {
		return nil, errors.New("email not found")
	}

	err = pkgs.CheckHashedPassword(userData.Password, userDetails.HashPassword)
	if err != nil {
		log.Info().Err(err).Send()
		return nil, errors.New("entered password is not wrong")
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "job portal project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token, err := s.a.GenerateToken(claims)
	if err != nil {
		return nil, err
	}

	return &model.Login{
		Name:  userDetails.Name,
		Email: userData.Email,
		Token: token,
	}, nil
}

func (s *Service) ViewJobByCid(cid string) ([]*model.Job, error) {
	jobDetails, err := s.userRepo.ViewJobByCid(cid)
	if err != nil {
		return nil, err
	}
	var jobDatas []*model.Job
	for _, v := range jobDetails {
		jobData := &model.Job{
			ID:     strconv.FormatUint(uint64(v.ID), 10),
			Cid:    strconv.FormatUint(uint64(v.Cid), 10),
			Role:   v.Role,
			Salary: v.Salary,
		}
		jobDatas = append(jobDatas, jobData)
	}
	return jobDatas, nil

}

func (s *Service) ViewAllJob() ([]*model.Job, error) {
	jobDetails, err := s.userRepo.ViewAllJob()
	if err != nil {
		return nil, err
	}
	var jobDatas []*model.Job

	for _, v := range jobDetails {
		jobData := &model.Job{
			ID:     strconv.FormatUint(uint64(v.ID), 10),
			Cid:    strconv.FormatUint(uint64(v.Cid), 10),
			Role:   v.Role,
			Salary: v.Salary,
		}
		jobDatas = append(jobDatas, jobData)
	}
	return jobDatas, nil

}

func (s *Service) ViewJobByID(id string) (*model.Job, error) {
	jobData, err := s.userRepo.ViewJobById(id)
	if err != nil {
		return &model.Job{}, err
	}
	return &model.Job{
		ID:     strconv.FormatUint(uint64(jobData.ID), 10),
		Cid:    strconv.FormatUint(uint64(jobData.Cid), 10),
		Role:   jobData.Role,
		Salary: jobData.Salary,
	}, nil

}

func (s *Service) CreateJob(JobDetails model.NewJob) (*model.Job, error) {
	uid, err := strconv.ParseUint(JobDetails.Cid, 10, 32)
	if err != nil {
		return nil, err
	}
	jd := models.Job{
		Cid:    uint(uid),
		Role:   JobDetails.Role,
		Salary: JobDetails.Salary,
	}
	cd, err := s.userRepo.CreateJob(jd)
	if err != nil {
		return nil, err
	}
	id := strconv.FormatUint(uint64(cd.ID), 10)
	return &model.Job{
		ID:     id,
		Cid:    strconv.FormatUint(uint64(cd.Cid), 10),
		Role:   jd.Role,
		Salary: jd.Salary,
	}, nil

}

func (s *Service) ViewCompanyById(cid string) (*model.Company, error) {
	companyData, err := s.userRepo.ViewCompanyByID(cid)
	if err != nil {
		return &model.Company{}, err
	}
	return &model.Company{
		ID:        strconv.FormatUint(uint64(companyData.ID), 10),
		Name:      companyData.Name,
		Location:  companyData.Location,
		CreatedAt: companyData.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: companyData.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *Service) ViewAllCompanies() ([]*model.Company, error) {
	companyDetails, err := s.userRepo.ViewAllCompany()
	if err != nil {
		return nil, err
	}
	var companyDatas []*model.Company

	for _, v := range companyDetails {
		companyData := &model.Company{
			ID:       strconv.FormatUint(uint64(v.ID), 10),
			Name:     v.Name,
			Location: v.Location,
		}
		companyDatas = append(companyDatas, companyData)
	}
	return companyDatas, nil

}

func (s *Service) CreateCompany(companyDetails model.NewCompnay) (*model.Company, error) {
	cd := models.Company{
		Name:     companyDetails.Name,
		Location: companyDetails.Location,
	}
	cd, err := s.userRepo.CreateCompany(cd)
	if err != nil {
		return nil, err
	}
	cid := strconv.FormatUint(uint64(cd.ID), 10)
	return &model.Company{
		ID:        cid,
		Name:      cd.Name,
		Location:  cd.Location,
		CreatedAt: cd.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: cd.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *Service) UserSignup(userData model.NewUser) (*model.User, error) {
	hashedPassword, err := pkgs.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}

	userDetails := models.User{
		Name:         userData.Name,
		Email:        userData.Email,
		HashPassword: hashedPassword,
	}

	userDetails, err = s.userRepo.CreateUser(userDetails)
	if err != nil {
		return nil, err
	}

	uid := strconv.FormatUint(uint64(userDetails.ID), 10)

	return &model.User{
		ID:        uid,
		Name:      userDetails.Name,
		Email:     userDetails.Email,
		CreatedAt: userDetails.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userDetails.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
