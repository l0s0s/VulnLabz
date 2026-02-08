package dsl

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNameRequired   = errors.New("scenario name is required")
	ErrStepsRequired  = errors.New("scenario must have at least one step")
	ErrStepMethod     = errors.New("step method is required")
	ErrStepURL        = errors.New("step url is required")
)

// Validate checks required fields and returns the first validation error.
func Validate(s *Scenario) error {
	if s == nil {
		return fmt.Errorf("scenario: %w", ErrNameRequired)
	}
	if strings.TrimSpace(s.Name) == "" {
		return ErrNameRequired
	}
	if len(s.Steps) == 0 {
		return ErrStepsRequired
	}
	for i := range s.Steps {
		if err := validateStep(&s.Steps[i], i); err != nil {
			return err
		}
	}
	return nil
}

func validateStep(st *ScenarioStep, index int) error {
	method := strings.TrimSpace(strings.ToUpper(st.Method))
	if method == "" {
		return fmt.Errorf("step[%d]: %w", index, ErrStepMethod)
	}
	if strings.TrimSpace(st.URL) == "" {
		return fmt.Errorf("step[%d]: %w", index, ErrStepURL)
	}
	// Normalize method into struct for future use
	st.Method = method
	return nil
}
