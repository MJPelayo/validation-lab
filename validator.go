package main

type Validator struct {
	Errors map[string]string
}

// Constructor
func newValidator() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

// Check runs validation tests
func (v *Validator) Check(ok bool, field, message string) {
	if !ok {
		v.AddError(field, message)
	}
}

// AddError adds an error to the map
func (v *Validator) AddError(field, message string) {
	if _, exists := v.Errors[field]; !exists {
		v.Errors[field] = message
	}
}

// Valid checks if there are no errors
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}