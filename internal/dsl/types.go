// Package dsl defines the scenario DSL: types, YAML parsing, and validation.
package dsl

// ScenarioType is the kind of security test (OWASP-style).
type ScenarioType string

const (
	TypeSQLi          ScenarioType = "sqli"
	TypeIDOR          ScenarioType = "idor"
	TypeSSRF          ScenarioType = "ssrf"
	TypeBruteForce    ScenarioType = "brute_force"
	TypeSecurityHeaders ScenarioType = "security_headers"
	TypeRateLimit     ScenarioType = "rate_limit"
)

// Scenario is the root DSL structure: a named security test with one or more steps.
type Scenario struct {
	Name        string        `yaml:"name"`
	Description string        `yaml:"description"`
	Type        ScenarioType  `yaml:"type"`
	Steps       []ScenarioStep `yaml:"steps"`
}

// ScenarioStep is a single HTTP request within a scenario.
type ScenarioStep struct {
	Name    string            `yaml:"name"`
	Method  string            `yaml:"method"`
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}
