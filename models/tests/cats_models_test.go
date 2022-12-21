package models

import (
	"encoding/json"
	"testing"

	"cat-backend/models"
)

func TestCatsMarshalJSON(t *testing.T) {
	cat := models.Cats{
		ID:    1,
		Name:  "Fluffy",
		Image: "http://example.com/fluffy.jpg",
	}
	expectedJSON := `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"id":1,"name":"Fluffy","image":"http://example.com/fluffy.jpg"}`
	bytes, err := json.Marshal(cat)
	if err != nil {
		t.Errorf("Error marshaling Cats to JSON: %v", err)
	}
	if string(bytes) != expectedJSON {
		t.Errorf("Expected JSON %s but got %s", expectedJSON, string(bytes))
	}
}

func TestCatsUnmarshalJSON(t *testing.T) {
	jsonStr := `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"id":1,"name":"Fluffy","image":"http://example.com/fluffy.jpg"}`
	expectedCat := models.Cats{
		ID:    1,
		Name:  "Fluffy",
		Image: "http://example.com/fluffy.jpg",
	}
	var cat models.Cats
	err := json.Unmarshal([]byte(jsonStr), &cat)
	if err != nil {
		t.Errorf("Error unmarshaling JSON to Cats: %v", err)
	}
	if cat != expectedCat {
		t.Errorf("Expected Cats %v but got %v", expectedCat, cat)
	}
}
