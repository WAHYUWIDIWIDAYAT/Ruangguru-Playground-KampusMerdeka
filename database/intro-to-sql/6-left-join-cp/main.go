package main

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db}
}

func (r *MovieRepository) FetchMoviesDirector() ([]model.MovieDirector, error) {
	var sqlStmt string

	// Task: Fetch all movies and their director
	// Hint: Use LEFT JOIN
	// See model.MovieDirector for the structure of the result

	
	//beginanswer
	//no such table genres
	sqlStmt = `
	SELECT
		m.id
		, m.title
		, m.description
		, m.year
		, d.Name AS director_name
	FROM movies m
	LEFT JOIN directors d ON m.director_id = d.id`
	
	//endanswer
	// TODO: answer here

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var movies []model.MovieDirector

	for rows.Next() {
		var movie model.MovieDirector

		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.DirectorName,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}
