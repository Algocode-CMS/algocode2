package choices

import (
	"database/sql/driver"
	"errors"
	"slices"
)

var ErrFieldIsNil = errors.New("select field value can not be null")
var ErrUnknownChoice = errors.New("select field value is not one of possible choices")

type SelectChoice struct {
	DisplayValue string
	InnerValue   string
}

type SelectBase interface {
	ListValues() []SelectChoice
	AllowNil() bool
}

// Probably should be a better way

type Select[T SelectBase] struct {
	value *string
}

func (s *Select[T]) Scan(src any) error {
	var t T
	if src == nil {
		if !t.AllowNil() {
			return ErrFieldIsNil
		}
		s.value = nil
		return nil
	}

	value, ok := src.(string)
	if !ok {
		return ErrUnknownChoice
	}

	if !slices.ContainsFunc(t.ListValues(), func(choice SelectChoice) bool {
		return choice.InnerValue == value
	}) {
		return ErrUnknownChoice
	}

	s.value = &value
	return nil
}

func (s *Select[T]) Value() (driver.Value, error) {
	var t T
	if s.value == nil {
		if t.AllowNil() {
			return nil, nil
		} else {
			return nil, ErrFieldIsNil
		}
	}

	if !slices.ContainsFunc(t.ListValues(), func(choice SelectChoice) bool {
		return choice.InnerValue == *s.value
	}) {
		return nil, ErrUnknownChoice
	}

	return *s.value, nil
}
