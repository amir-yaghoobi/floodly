package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/amir-yaghoobi/floodly/migration"
	"github.com/amir-yaghoobi/floodly/user"
)

type Repository struct {
	db *sql.DB
}

func (r Repository) DropTable() error {
	tbName := migration.TableName(user.User{})
	_, err := r.db.Exec("DROP TABLE IF EXISTS " + tbName)
	return err
}

func (r Repository) Migrate(hardReset bool) error {
	if hardReset {
		err := r.DropTable()
		if err != nil {
			return err
		}
	}

	createStatement := migration.GetSchemaForCrate(user.User{})

	_, err := r.db.Exec(createStatement)
	return err
}

func (r Repository) Create(u *user.User) error {
	tbName := migration.TableName(*u)
	columns := migration.GetColumnNames(*u)
	columnNames := strings.Join(columns, " , ")

	valuesHolder := make([]string, len(columns))
	for i := range columns {
		valuesHolder[i] = "?"
	}

	values := strings.Join(valuesHolder, ",")

	insertStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tbName, columnNames, values)

	statement, err := r.db.Prepare(insertStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		u.Email,
		u.UserName,
		u.Password,
		u.FirstName,
		u.LastName,
		u.NationalID,
		u.Picture,
		u.Gender,
		u.BirthDay,
		u.RegisterDate,
		u.LastLogin,
		u.LastIP,
		u.TimeZone,
		u.Country,
		u.Address,
		u.PostalCode,
		u.Phone,
		u.Salary,
	)
	return err
}

func NewCrateRepository(db *sql.DB) *Repository {
	return &Repository{db}
}
