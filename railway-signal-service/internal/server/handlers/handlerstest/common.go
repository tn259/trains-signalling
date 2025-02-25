package handlerstest

import "crosstech-hw/railway-signal-service/internal/database"

// FakeDAO
// Strategy is to first check error corresponding to the menthod and entity
// If no error is found, then check if the entity is present in operations except for create
// We control test preconditions by setting the error and the entity appropriately
type FakeDao struct {
	CreateSignalError error
	SignalsError      error
	CreateTrackError  error
	TracksError       error
	SignalErrorsByID  map[int]error
	TrackErrorsByID   map[int]error
	SignalsByID       map[int]*database.Signal
	TracksByID        map[int]*database.Track
}

func NewFakeDao() *FakeDao {
	return &FakeDao{
		SignalErrorsByID: make(map[int]error),
		TrackErrorsByID:  make(map[int]error),
		SignalsByID:      make(map[int]*database.Signal),
		TracksByID:       make(map[int]*database.Track),
	}
}

func (d *FakeDao) Signals() ([]*database.Signal, error) {
	if d.SignalsError != nil {
		return nil, d.SignalsError
	}
	signals := make([]*database.Signal, 0, len(d.SignalsByID))
	for _, s := range d.SignalsByID {
		signals = append(signals, s)
	}
	return signals, nil
}

func (d *FakeDao) SignalByID(id int) (*database.Signal, error) {
	if err, ok := d.SignalErrorsByID[id]; ok {
		return nil, err
	}
	s, ok := d.SignalsByID[id]
	if !ok {
		return nil, nil
	}
	return s, nil
}

func (d *FakeDao) CreateSignal(s *database.Signal) error {
	if d.CreateSignalError != nil {
		return d.CreateSignalError
	}
	d.SignalsByID[s.ID] = s
	return nil
}

func (d *FakeDao) UpdateSignal(s *database.Signal) (bool, error) {
	if err, ok := d.SignalErrorsByID[s.ID]; ok {
		return false, err
	}
	if _, ok := d.SignalsByID[s.ID]; !ok {
		return false, nil
	}
	d.SignalsByID[s.ID] = s
	return true, nil
}

func (d *FakeDao) DeleteSignal(id int) (bool, error) {
	if err, ok := d.SignalErrorsByID[id]; ok {
		return false, err
	}
	if _, ok := d.SignalsByID[id]; !ok {
		return false, nil
	}
	delete(d.SignalsByID, id)
	return true, nil
}

func (d *FakeDao) Tracks() ([]*database.Track, error) {
	if d.TracksError != nil {
		return nil, d.TracksError
	}
	tracks := make([]*database.Track, 0, len(d.TracksByID))
	for _, t := range d.TracksByID {
		tracks = append(tracks, t)
	}
	return tracks, nil
}

func (d *FakeDao) TrackByID(id int) (*database.Track, error) {
	if err, ok := d.TrackErrorsByID[id]; ok {
		return nil, err
	}
	t, ok := d.TracksByID[id]
	if !ok {
		return nil, nil
	}
	return t, nil
}

func (d *FakeDao) CreateTrack(t *database.Track) error {
	if d.CreateTrackError != nil {
		return d.CreateTrackError
	}
	d.TracksByID[t.ID] = t
	return nil
}

func (d *FakeDao) UpdateTrack(t *database.Track) (bool, error) {
	if err, ok := d.TrackErrorsByID[t.ID]; ok {
		return false, err
	}
	if _, ok := d.TracksByID[t.ID]; !ok {
		return false, nil
	}
	d.TracksByID[t.ID] = t
	return true, nil
}

func (d *FakeDao) DeleteTrack(id int) (bool, error) {
	if err, ok := d.TrackErrorsByID[id]; ok {
		return false, err
	}
	if _, ok := d.TracksByID[id]; !ok {
		return false, nil
	}
	delete(d.TracksByID, id)
	return true, nil
}
