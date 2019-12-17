package seed

import (
	"context"
	"database/sql"
	"github.com/oreqizer/goiler/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// Seed fills the database with testing data
func Seed(url string) error {
	dbi, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	defer dbi.Close()

	if err := dbi.Ping(); err != nil {
		return err
	}

	ctx := context.Background()

	model := models.Account{
		AuthID:  "TEST",
		Name:    "Yolo",
		Surname: "Swagger",
		Email:   "yolo@swagger.com",
		IsAdmin: true,
	}

	if err := model.Insert(ctx, dbi, boil.Infer()); err != nil {
		return err
	}
	return nil
}
