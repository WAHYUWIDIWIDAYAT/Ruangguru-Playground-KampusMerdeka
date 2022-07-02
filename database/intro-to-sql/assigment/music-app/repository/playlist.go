package repository

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/assigment/music-app/model"
)

type PlaylistRepositoryInterface interface {
	FetchUserPlaylists(userID int64) ([]model.UserPlaylist, error)
	CreatePlaylist(playlist model.Playlist) error
	UpdateUserPlaylist(userID int64, playlist model.Playlist) error
	FetchPlaylistTrack(playlistID int64) ([]model.PlaylistTrack, error)
}

type PlaylistRepository struct {
	db *sql.DB
}

func NewPlaylistRepository(db *sql.DB) PlaylistRepositoryInterface {
	return &PlaylistRepository{db}
}

func (p *PlaylistRepository) FetchUserPlaylists(userID int64) ([]model.UserPlaylist, error) {
	var sqlStatement string
	var playlists []model.UserPlaylist

	sqlStatement = "SELECT playlists.id, playlists.name, playlists.user_id, users.name FROM playlists LEFT JOIN users ON playlists.user_id = users.id WHERE user_id = ?"


	rows, err := p.db.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var playlist model.UserPlaylist
		err := rows.Scan(
			&playlist.PlaylistID,
			&playlist.PlaylistName,
			&playlist.UserID,
			&playlist.UserName,
		)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (p *PlaylistRepository) CreatePlaylist(playlist model.Playlist) error {
	var sqlStatement string

	sqlStatement = "INSERT INTO playlists (id, name, created_at, user_id) VALUES(?, ?, ?, ?)"

	_, err := p.db.Exec(sqlStatement, playlist.ID, playlist.Name, playlist.CreatedAt, playlist.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlaylistRepository) UpdateUserPlaylist(userID int64, playlist model.Playlist) error {
	var sqlStatement string

	sqlStatement = "UPDATE playlists SET name = ?, created_at = ? where user_id = ? and id = ?"

	_, err := p.db.Exec(sqlStatement, playlist.Name, playlist.CreatedAt, playlist.UserID, playlist.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlaylistRepository) FetchPlaylistTrack(playlistID int64) ([]model.PlaylistTrack, error) {
	var sqlStatement string

	sqlStatement = "SELECT playlists.id, playlists.name, tracks.id, tracks.name, tracks.artist FROM playlist_tracks LEFT JOIN playlists ON playlist_tracks.playlist_id = playlists.id LEFT JOIN tracks ON playlist_tracks.track_id = tracks.id  where playlists.id = ?"

	rows, err := p.db.Query(sqlStatement, playlistID)
	if err != nil {
		return nil, err
	}

	var playlistTracks []model.PlaylistTrack
	for rows.Next() {
		playlistTrack := model.PlaylistTrack{}
		err := rows.Scan(
			&playlistTrack.PlaylistID,
			&playlistTrack.PlaylistName,
			&playlistTrack.TrackID,
			&playlistTrack.TrackName,
			&playlistTrack.TrackArtist,
		)
		if err != nil {
			return nil, err
		}

		playlistTracks = append(playlistTracks, playlistTrack)
	}

	return playlistTracks, nil
}
