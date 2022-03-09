package uidgen

import "github.com/google/uuid"

type UIDGen interface {
	NewUuid() string
}

type uidgen struct{}

func New() UIDGen {
	return &uidgen{}
}

func (u uidgen) NewUuid() string {
	return uuid.New().String()
}
