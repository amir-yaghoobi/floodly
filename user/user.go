package user

import "time"

type User struct {
	Email        string    `crate:"column:email"`
	UserName     string    `crate:"column:user_name"`
	Password     string    `crate:"column:password"`
	FirstName    string    `crate:"column:first_name"`
	LastName     string    `crate:"column:last_name"`
	NationalID   string    `crate:"column:national_id"`
	Picture      string    `crate:"column:picture"`
	Gender       bool      `crate:"column:gender"`
	BirthDay     time.Time `crate:"column:birthday"`
	RegisterDate time.Time `crate:"column:register_date"`
	LastLogin    time.Time `crate:"column:last_login"`
	LastIP       string    `crate:"column:last_ip;type:ip"`
	TimeZone     string    `crate:"column:time_zone"`
	Country      string    `crate:"column:country"`
	Address      string    `crate:"column:address"`
	PostalCode   string    `crate:"column:postal_code"`
	Phone        string    `crate:"column:phone"`
	Salary       float64   `crate:"column:salary"`
}
