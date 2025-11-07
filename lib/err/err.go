package err

import "fmt"

func Wrap(msg string, err error) error {
	return fmt.Errorf("can't do request: %w", err)
}
