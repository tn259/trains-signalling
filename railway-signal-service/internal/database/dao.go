package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Dao interface {
	Signals() ([]*Signal, error)
	SignalByID(id int) (*Signal, error)
	CreateSignal(s *Signal) error
	UpdateSignal(s *Signal) (bool, error)
	DeleteSignal(id int) (bool, error)

	Tracks() ([]*Track, error)
	TrackByID(id int) (*Track, error)
	CreateTrack(t *Track) error
	UpdateTrack(t *Track) (bool, error)
	DeleteTrack(id int) (bool, error)
}

type PGDao struct {
	db *DB
}

func NewPGDao(db *DB) *PGDao {
	return &PGDao{
		db: db,
	}
}

func (d *PGDao) Signals() ([]*Signal, error) {
	var signals []*Signal
	err := d.db.Model(&signals).Select()
	if err != nil {
		return nil, fmt.Errorf("d.db.Model(&signals).Select(): %v", err)
	}
	return signals, nil
}

func (d *PGDao) SignalByID(id int) (*Signal, error) {
	signal := &Signal{ID: id}
	err := d.db.Model(signal).WherePK().Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("d.db.Model(signal).WherePK().Select(): %v", err)
	}
	return signal, nil
}

func (d *PGDao) CreateSignal(s *Signal) error {
	_, err := d.db.Model(s).Insert()
	if err != nil {
		return fmt.Errorf("d.db.Model(s).Insert(): %v", err)
	}
	return nil
}

func (d *PGDao) UpdateSignal(s *Signal) (bool, error) {
	r, err := d.db.Model(s).WherePK().Update()
	if err != nil {
		return false, fmt.Errorf("d.db.Model(s).WherePK().Update(): %v", err)
	}
	return r.RowsAffected() > 0, nil
}

func (d *PGDao) DeleteSignal(id int) (bool, error) {
	signal := &Signal{ID: id}
	r, err := d.db.Model(signal).WherePK().Delete()
	if err != nil {
		return false, fmt.Errorf("d.db.Model(signal).WherePK().Delete(): %v", err)
	}
	return r.RowsAffected() > 0, nil
}

func (d *PGDao) Tracks() ([]*Track, error) {
	var tracks []*Track
	err := d.db.Model(&tracks).Select()
	if err != nil {
		return nil, fmt.Errorf("d.db.Model(&tracks).Select(): %v", err)
	}
	return tracks, nil
}

func (d *PGDao) TrackByID(id int) (*Track, error) {
	track := &Track{ID: id}
	err := d.db.Model(track).WherePK().Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("d.db.Model(track).WherePK().Select(): %v", err)
	}
	return track, nil
}

func (d *PGDao) CreateTrack(t *Track) error {
	_, err := d.db.Model(t).Insert()
	if err != nil {
		return fmt.Errorf("d.db.Model(t).Insert(): %v", err)
	}
	return nil
}

func (d *PGDao) UpdateTrack(t *Track) (bool, error) {
	r, err := d.db.Model(t).WherePK().Update()
	if err != nil {
		return false, fmt.Errorf("d.db.Model(t).WherePK().Update(): %v", err)
	}
	return r.RowsAffected() > 0, nil
}

func (d *PGDao) DeleteTrack(id int) (bool, error) {
	track := &Track{ID: id}
	r, err := d.db.Model(track).WherePK().Delete()
	if err != nil {
		return false, fmt.Errorf("d.db.Model(track).WherePK().Delete(): %v", err)
	}
	return r.RowsAffected() > 0, nil
}
