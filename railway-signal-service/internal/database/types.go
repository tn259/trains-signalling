package database

// ELR represents a railway line
type ELR struct {
	ID   int    `pg:"id,pk"`
	Name string `pg:"name,notnull,unique"`
}

// Signal represents a railway signal
type Signal struct {
	ID   int    `pg:"id,pk"`
	Name string `pg:"name"` // Allow null values
}

// Track represents a railway track section
type Track struct {
	ID     int    `pg:"id,pk"`
	Source string `pg:"source,notnull"`
	Target string `pg:"target,notnull"`
}

// TrackSignal represents the junction between tracks and signals with mileage information
type TrackSignal struct {
	ID       int     `pg:"id,pk"`
	Mileage  float64 `pg:"mileage"`
	ELRID    int     `pg:"elr_id,notnull"`
	SignalID int     `pg:"signal_id,notnull"`
	TrackID  int     `pg:"track_id,notnull"`

	// Relations
	ELR    *ELR    `pg:"rel:has-one"`
	Signal *Signal `pg:"rel:has-one"`
	Track  *Track  `pg:"rel:has-one"`
}
