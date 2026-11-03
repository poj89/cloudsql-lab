
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/poj89/cloudsql-lab/internal/instance"
	"github.com/poj89/cloudsql-lab/internal/storage"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string, output io.Writer) error {
	if len(args) == 0 {
		return errors.New("usage: cloudsql <create|list|delete> [name]")
	}

	store := storage.NewJSONStore(filepath.Join("data", "instances.json"))

	switch args[0] {
	case "create":
		if len(args) != 2 {
			return errors.New("usage: cloudsql create <name>")
		}
		if err := store.Create(instance.Instance{Name: args[1]}); err != nil {
			return err
		}
		_, err := fmt.Fprintf(output, "Created instance %q\n", args[1])
		return err
	case "list":
		if len(args) != 1 {
			return errors.New("usage: cloudsql list")
		}
		instances, err := store.List()
		if err != nil {
			return err
		}
		if len(instances) == 0 {
			_, err := fmt.Fprintln(output, "No instances found.")
			return err
		}
		if _, err := fmt.Fprintln(output, "NAME"); err != nil {
			return err
		}
		for _, item := range instances {
			if _, err := fmt.Fprintln(output, item.Name); err != nil {
				return err
			}
		}
		return nil
	case "delete":
		if len(args) != 2 {
			return errors.New("usage: cloudsql delete <name>")
		}
		if err := store.Delete(args[1]); err != nil {
			return err
		}
		_, err := fmt.Fprintf(output, "Deleted instance %q\n", args[1])
		return err
	default:
		return fmt.Errorf("unknown command %q; usage: cloudsql <create|list|delete> [name]", args[0])
	}
}
