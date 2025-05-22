package proposals

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Proposal, error)
	GetByProcessID(ctx context.Context, processID int) ([]Proposal, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]Proposal, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, description, start_at, end_at FROM proposal")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var processes []Proposal
	for rows.Next() {
		var p Proposal
		if err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		processes = append(processes, p)
	}
	return processes, nil
}

func (r *repository) GetByProcessID(ctx context.Context, processID int) ([]Proposal, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, description, start_at, end_at FROM proposal WHERE process_id = $1", processID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var proposals []Proposal
	for rows.Next() {
		var p Proposal
		if err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		proposals = append(proposals, p)
	}
	return proposals, nil
}
