package repository

import (
	"context"

	"github.com/Turtel216/collecix-web-service/internal/models"
	"github.com/Turtel216/collecix-web-service/internal/repository/collection"
)

// Repo interface for data base methods
type Repo interface {
	Insert(ctx context.Context, collection models.Collection) error
	FindById(ctx context.Context, id uint64) (models.Collection, error)
	DeleteById(ctx context.Context, id uint64) error
	Update(ctx context.Context, collection models.Collection) error
	FindAll(ctx context.Context, page collection.FindAllPage) (collection.FindResult, error)
}
