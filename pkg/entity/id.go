package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func ParseUid(s string) (ID, error) {
	ui, err := uuid.Parse(s)

	if err != nil {
		return uuid.Nil, err
	}
	return ui, nil
}
