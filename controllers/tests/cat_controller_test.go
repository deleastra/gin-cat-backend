package controllers

import (
	"cat-backend/models"
	"cat-backend/routers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCatsController(t *testing.T) {
	// Set up mock DB and dialector.
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_dns",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	// Set up Gin router and test server.
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	r := routers.SetCatRoutes(gormDB, router)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Test GetCats.
	t.Run("GetCats", func(t *testing.T) {
		// Set up mock DB response.
		cats := []models.Cats{
			{Name: "Fluffy", Image: "fluffy.jpg"},
			{Name: "Whiskers", Image: "whiskers.jpg"},
		}
		rows := sqlmock.NewRows([]string{"name", "image"})
		for _, cat := range cats {
			rows = rows.AddRow(cat.Name, cat.Image)
		}

		// Send request and assert on response.
		res, err := http.Get(ts.URL + "/cats")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
