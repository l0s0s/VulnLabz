package dsl

import (
	"errors"
	"testing"
)

func TestParseYAML(t *testing.T) {
	valid := []byte(`
name: "SQLi login bypass"
description: "Attempt SQL injection on login form"
type: sqli
steps:
  - name: "Send payload"
    method: POST
    url: "https://example.com/login"
    headers:
      Content-Type: application/x-www-form-urlencoded
    body: "username=admin'--&password=x"
`)
	s, err := ParseYAML(valid)
	if err != nil {
		t.Fatalf("ParseYAML: %v", err)
	}
	if s.Name != "SQLi login bypass" {
		t.Errorf("name: got %q", s.Name)
	}
	if s.Type != TypeSQLi {
		t.Errorf("type: got %q", s.Type)
	}
	if len(s.Steps) != 1 {
		t.Fatalf("steps: got %d", len(s.Steps))
	}
	if s.Steps[0].Method != "POST" || s.Steps[0].URL != "https://example.com/login" {
		t.Errorf("step: method=%q url=%q", s.Steps[0].Method, s.Steps[0].URL)
	}
}

func TestParseAndValidate(t *testing.T) {
	valid := []byte(`
name: "Minimal"
steps:
  - method: GET
    url: "https://example.com"
`)
	s, err := ParseAndValidate(valid)
	if err != nil {
		t.Fatalf("ParseAndValidate: %v", err)
	}
	if s.Name != "Minimal" || len(s.Steps) != 1 {
		t.Errorf("unexpected scenario: %+v", s)
	}
}

func TestValidate_errors(t *testing.T) {
	tests := []struct {
		name string
		yaml string
		want error
	}{
		{"empty name", "name: \"\"\nsteps:\n  - method: GET\n    url: u", ErrNameRequired},
		{"no steps", "name: x\nsteps: []", ErrStepsRequired},
		{"step no method", "name: x\nsteps:\n  - url: u", ErrStepMethod},
		{"step no url", "name: x\nsteps:\n  - method: GET", ErrStepURL},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := ParseYAML([]byte(tt.yaml))
			if err != nil {
				t.Fatalf("parse: %v", err)
			}
			err = Validate(s)
			if err == nil || !errors.Is(err, tt.want) {
				t.Errorf("Validate: got %v, want %v", err, tt.want)
			}
		})
	}
}
