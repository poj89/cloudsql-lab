package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/poj89/cloudsql-lab/internal/instance"
)

var (
	ErrAlreadyExists = errors.New("instance already exists")
	ErrNotFound      = errors.New("instance not found")
)

// JSONStore persists instance metadata as a JSON array.
type JSONStore struct {
	path string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}

func (s *JSONStore) List() ([]instance.Instance, error) {
	instances, err := s.read()
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (s *JSONStore) Create(item instance.Instance) error {
	if err := item.Validate(); err != nil {
		return err
	}

	instances, err := s.read()
	if err != nil {
		return err
	}
	for _, current := range instances {
		if current.Name == item.Name {
			return fmt.Errorf("%w: %q", ErrAlreadyExists, item.Name)
		}
	}

	instances = append(instances, item)
	return s.write(instances)
}

func (s *JSONStore) Delete(name string) error {
	if err := (instance.Instance{Name: name}).Validate(); err != nil {
		return err
	}

	instances, err := s.read()
	if err != nil {
		return err
	}
	for index, current := range instances {
		if current.Name == name {
			instances = append(instances[:index], instances[index+1:]...)
			return s.write(instances)
		}
	}
	return fmt.Errorf("%w: %q", ErrNotFound, name)
}

func (s *JSONStore) read() ([]instance.Instance, error) {
	data, err := os.ReadFile(s.path)
	if errors.Is(err, os.ErrNotExist) {
		return []instance.Instance{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read instance registry %q: %w", s.path, err)
	}

	var instances []instance.Instance
	if err := json.Unmarshal(data, &instances); err != nil {
		return nil, fmt.Errorf("decode instance registry %q: %w", s.path, err)
	}
	if instances == nil {
		return nil, fmt.Errorf("decode instance registry %q: expected a JSON array", s.path)
	}
	for index, item := range instances {
		if err := item.Validate(); err != nil {
			return nil, fmt.Errorf("invalid instance at index %d in registry %q: %w", index, s.path, err)
		}
		for previous := 0; previous < index; previous++ {
			if instances[previous].Name == item.Name {
				return nil, fmt.Errorf("duplicate instance name %q in registry %q", item.Name, s.path)
			}
		}
	}
	return instances, nil
}

func (s *JSONStore) write(instances []instance.Instance) error {
	data, err := json.MarshalIndent(instances, "", "  ")
	if err != nil {
		return fmt.Errorf("encode instance registry: %w", err)
	}
	data = append(data, '\n')

	directory := filepath.Dir(s.path)
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("create instance registry directory %q: %w", directory, err)
	}
	tempFile, err := os.CreateTemp(directory, "."+filepath.Base(s.path)+"-*.tmp")
	if err != nil {
		return fmt.Errorf("create temporary instance registry: %w", err)
	}
	tempPath := tempFile.Name()
	closed := false
	defer func() {
		if !closed {
			_ = tempFile.Close()
		}
		_ = os.Remove(tempPath)
	}()

	if _, err := tempFile.Write(data); err != nil {
		return fmt.Errorf("write temporary instance registry: %w", err)
	}
	if err := tempFile.Sync(); err != nil {
		return fmt.Errorf("sync temporary instance registry: %w", err)
	}
	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("close temporary instance registry: %w", err)
	}
	closed = true
	if err := os.Rename(tempPath, s.path); err != nil {
		return fmt.Errorf("replace instance registry %q: %w", s.path, err)
	}
	return nil
}
