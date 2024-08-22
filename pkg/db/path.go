package db

import (
	"context"
	"time"

	"github.com/feldtsen/farrago/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PathRepository interface {
	GetPaths() ([]models.Path, error) // TODO: should probably ad the possibility to get all path between givien coordinates (would be good for the map). Will simplify this for now and fetch everything
	CreatePath() (models.Path, error)
	UpdatePath() (models.Path, error)
	DeletePath() (models.DataManipulationResult, error)
}

type PgxPathRepository struct {
	DB *pgxpool.Pool
}

func (repo *PgxPathRepository) GetPaths() ([]models.Path, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, title, description, image_lookup_id, latitude, longitude
		FROM paths 
	`

	rows, err := repo.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paths []models.Path

	for rows.Next() {
		var path models.Path
		if err := rows.Scan(&path.ID, &path.Title, &path.ImageLookupId, &path.Coordinate.Latitude, &path.Coordinate.Longitude); err != nil {
			return nil, err
		}

		paths = append(paths, path)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paths, nil
}

func (repo *PgxPathRepository) CreatePath() {

}

func (repo *PgxPathRepository) UpdatePath() {

}

func (repo *PgxPathRepository) DeletePath() {

}
