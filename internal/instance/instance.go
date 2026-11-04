package instance

import "fmt"

// Instance is the metadata Phase 1 stores for a PostgreSQL instance.
type Instance struct {
	Name   string `json:"name"`
	Engine string `json:"engine"`
	Status string `json:"status"`
}

// Validate checks that the name is safe to use as a future Podman container name.
func (i Instance) Validate() error {
	if i.Name == "" || !isAlphaNumeric(i.Name[0]) {
		return fmt.Errorf("invalid instance name %q: it must start with a letter or digit", i.Name)
	}
	for index := 1; index < len(i.Name); index++ {
		character := i.Name[index]
		if !isAlphaNumeric(character) && character != '.' && character != '_' && character != '-' {
			return fmt.Errorf("invalid instance name %q: use only letters, digits, '.', '_' or '-'", i.Name)
		}
	}
	return nil
}

func isAlphaNumeric(character byte) bool {
	return character >= 'a' && character <= 'z' ||
		character >= 'A' && character <= 'Z' ||
		character >= '0' && character <= '9'
}
