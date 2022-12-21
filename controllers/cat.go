package controllers

import (
	"context"

	"example.com/golang-restfulapi/models"

	"gorm.io/gorm"
)

type CatsController struct {
	DB *gorm.DB
}

func (c *CatsController) GetCats(ctx context.Context) ([]models.Cats, error) {
	var cats []models.Cats
	if err := c.DB.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (c *CatsController) GetCatByID(ctx context.Context, id uint) (*models.Cats, error) {
	var cat models.Cats
	if err := c.DB.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (c *CatsController) CreateCat(ctx context.Context, cat *models.Cats) error {
	if err := c.DB.Create(cat).Error; err != nil {
		return err
	}
	return nil
}

func (c *CatsController) UpdateCat(ctx context.Context, cat *models.Cats) error {
	if err := c.DB.Save(cat).Error; err != nil {
		return err
	}
	return nil
}

func (c *CatsController) DeleteCat(ctx context.Context, id uint) error {
	cat, err := c.GetCatByID(ctx, id)
	if err != nil {
		return err
	}
	if err := c.DB.Delete(cat).Error; err != nil {
		return err
	}
	return nil
}
