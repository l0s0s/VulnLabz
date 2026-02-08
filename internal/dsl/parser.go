package dsl

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ParseYAML parses scenario from YAML bytes. Call Validate to check required fields.
func ParseYAML(data []byte) (*Scenario, error) {
	var s Scenario
	if err := yaml.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("yaml: %w", err)
	}
	return &s, nil
}

// LoadFile reads and parses a scenario from a YAML file.
func LoadFile(path string) (*Scenario, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	return ParseYAML(data)
}

// ParseAndValidate parses YAML and validates the scenario. Returns parsed scenario or error.
func ParseAndValidate(data []byte) (*Scenario, error) {
	s, err := ParseYAML(data)
	if err != nil {
		return nil, err
	}
	if err := Validate(s); err != nil {
		return nil, err
	}
	return s, nil
}
