// Package repository provides all repositories
package repository

import "github.com/studentkickoff/gobp/pkg/db"

// Repository is used to create specific repositories
type Repository struct {
	db db.DB
}

// New creates a new repository
func New(db db.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) NewUser() User {
	return &userRepo{db: r.db}
}
