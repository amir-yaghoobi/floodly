package user

import (
	"sync"
	"time"

	"github.com/amir-yaghoobi/floodly/reporter"

	"github.com/brianvoe/gofakeit"
)

type UserGenerator struct {
	repository UserRepository
}

func (s UserGenerator) generateNewUser() *User {
	birthDayFrom, _ := time.Parse(time.RFC3339, "1990-01-01T02:03:00+03:30")
	birthDayTo, _ := time.Parse(time.RFC3339, "2010-12-01T02:03:07+03:30")

	registerFrom, _ := time.Parse(time.RFC3339, "2017-09-01T02:03:00+03:30")
	registerTo, _ := time.Parse(time.RFC3339, "2019-01-01T02:03:07+03:30")

	loginFrom, _ := time.Parse(time.RFC3339, "2019-03-01T02:03:00+03:30")
	loginTo, _ := time.Parse(time.RFC3339, "2019-03-25T09:03:07+03:30")

	addressInfo := gofakeit.Address()

	return &User{
		Email:        gofakeit.Email(),
		UserName:     gofakeit.Username(),
		Password:     gofakeit.Password(true, true, true, true, false, 32),
		FirstName:    gofakeit.FirstName(),
		LastName:     gofakeit.LastName(),
		NationalID:   gofakeit.SSN(),
		Picture:      gofakeit.ImageURL(250, 250),
		Gender:       gofakeit.Bool(),
		BirthDay:     gofakeit.DateRange(birthDayFrom, birthDayTo),
		RegisterDate: gofakeit.DateRange(registerFrom, registerTo),
		LastLogin:    gofakeit.DateRange(loginFrom, loginTo),
		LastIP:       gofakeit.IPv4Address(),
		TimeZone:     gofakeit.TimeZone(),
		Country:      gofakeit.Country(),
		Address:      addressInfo.Address,
		PostalCode:   addressInfo.Zip,
		Phone:        gofakeit.Phone(),
		Salary:       float64(gofakeit.Number(1000, 9999)),
	}
}

func (s UserGenerator) Generate(count int, resultC chan<- reporter.Result, done *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		start := time.Now()
		user := s.generateNewUser()
		err := s.repository.Create(user)

		resultC <- reporter.Result{
			Duration: time.Since(start),
			Error:    err,
		}
	}

	done.Done()
}

func NewUserGenerator(repository UserRepository) *UserGenerator {
	return &UserGenerator{repository}
}
