package processes

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Process, error)
	GetProcessByID(ctx context.Context, id int) (Process, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]Process, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, description, start_at, end_at FROM processes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var processes []Process
	for rows.Next() {
		var p Process
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.StartAt, &p.EndAt); err != nil {
			return nil, err
		}
		processes = append(processes, p)
	}
	return processes, nil
}

func (r *repository) GetProcessByID(ctx context.Context, id int) (Process, error) {
	var p Process
	err := r.db.QueryRowContext(ctx,
		`SELECT id, title, description, start_at, end_at
			 FROM processes
			WHERE id = $1`, id,
	).Scan(&p.ID, &p.Title, &p.Description, &p.StartAt, &p.EndAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return Process{}, nil
		}
		return Process{}, err
	}

	return p, nil

}
