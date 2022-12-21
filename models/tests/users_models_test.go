package models

import (
	"testing"

	"cat-backend/models"
)

func TestUsers(t *testing.T) {
	t.Run("Test valid user", func(t *testing.T) {
		username := "testuser"
		password := "testpass"
		displayName := "Test User"

		user := &models.Users{
			Username:    &username,
			Password:    &password,
			DisplayName: &displayName,
		}

		if user.Username == nil {
			t.Errorf("Expected non-nil username, got nil")
		}
		if *user.Username != username {
			t.Errorf("Expected username %q, got %q", username, *user.Username)
		}
		if user.Password == nil {
			t.Errorf("Expected non-nil password, got nil")
		}
		if *user.Password != password {
			t.Errorf("Expected password %q, got %q", password, *user.Password)
		}
		if user.DisplayName == nil {
			t.Errorf("Expected non-nil display name, got nil")
		}
		if *user.DisplayName != displayName {
			t.Errorf("Expected display name %q, got %q", displayName, *user.DisplayName)
		}
	})

	t.Run("Test nil username", func(t *testing.T) {
		password := "testpass"
		displayName := "Test User"

		user := &models.Users{
			Password:    &password,
			DisplayName: &displayName,
		}

		if user.Username != nil {
			t.Errorf("Expected nil username, got %q", *user.Username)
		}
		if user.Password == nil {
			t.Errorf("Expected non-nil password, got nil")
		}
		if *user.Password != password {
			t.Errorf("Expected password %q, got %q", password, *user.Password)
		}
		if user.DisplayName == nil {
			t.Errorf("Expected non-nil display name, got nil")
		}
		if *user.DisplayName != displayName {
			t.Errorf("Expected display name %q, got %q", displayName, *user.DisplayName)
		}
	})

	t.Run("Test nil password", func(t *testing.T) {
		username := "testuser"
		displayName := "Test User"

		user := &models.Users{
			Username:    &username,
			DisplayName: &displayName,
		}

		if user.Username == nil {
			t.Errorf("Expected non-nil username, got nil")
		}
		if *user.Username != username {
			t.Run("Test nil password", func(t *testing.T) {
				username := "testuser"
				displayName := "Test User"

				user := &models.Users{
					Username:    &username,
					DisplayName: &displayName,
				}

				if user.Username == nil {
					t.Errorf("Expected non-nil username, got nil")
				}
				if *user.Username != username {
					t.Errorf("Expected username %q, got %q", username, *user.Username)
				}
				if user.Password != nil {
					t.Errorf("Expected nil password, got %q", *user.Password)
				}
				if user.DisplayName == nil {
					t.Errorf("Expected non-nil display name, got nil")
				}
				if *user.DisplayName != displayName {
					t.Errorf("Expected display name %q, got %q", displayName, *user.DisplayName)
				}
			})

		}
	})
}
