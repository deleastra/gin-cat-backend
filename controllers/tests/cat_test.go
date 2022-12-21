package controllers

import (
	"context"
	"reflect"
	"testing"

	"example.com/golang-restfulapi/controllers"
	"example.com/golang-restfulapi/models"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCatsController_GetCats(t *testing.T) {
	// Set up test data.
	cat1 := models.Cats{ID: 1, Name: "Fluffy", Image: "http://example.com/fluffy.jpg"}
	cat2 := models.Cats{ID: 2, Name: "Whiskers", Image: "http://example.com/whiskers.jpg"}
	expectedCats := []models.Cats{cat1, cat2}

	// Set up mock database.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Convert the *sql.DB value to a *gorm.DB value.
	gormDB, err := gorm.Open(sqlite.Open("cat.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a gorm database connection", err)
	}

	// Expect a SELECT query to be executed, and return a slice of rows as the result.
	rows := sqlmock.NewRows([]string{"id", "name", "image"}).
		AddRow(1, "Fluffy", "http://example.com/fluffy.jpg").
		AddRow(2, "Whiskers", "http://example.com/whiskers.jpg")
	mock.ExpectQuery("SELECT \\* FROM cats").WillReturnRows(rows)

	// Set up controller.
	ctrl := controllers.CatsController{DB: gormDB}

	// Call method under test.
	cats, err := ctrl.GetCats(context.Background())

	// Verify results.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(cats, expectedCats) {
		t.Errorf("Expected cats %v but got %v", expectedCats, cats)
		// Verify mock database behavior.
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	}
}
