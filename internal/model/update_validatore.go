package model

import "errors"

type UpdateUserInput struct {
	Name *string `json:"name"`
}

func (i UpdateUserInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
