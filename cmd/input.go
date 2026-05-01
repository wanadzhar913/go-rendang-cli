package cmd

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"go-rendang-cli/internal/rendang"
)

func collectOrder(reader *bufio.Reader, out io.Writer, order rendang.Order, allowPrompt bool) (rendang.Order, error) {
	var err error

	if strings.TrimSpace(order.FirstName) == "" {
		if !allowPrompt {
			return order, fmt.Errorf("first name is required")
		}
		order.FirstName, err = prompt(reader, out, "First name")
		if err != nil {
			return order, err
		}
	}

	if strings.TrimSpace(order.LastName) == "" {
		if !allowPrompt {
			return order, fmt.Errorf("last name is required")
		}
		order.LastName, err = prompt(reader, out, "Last name")
		if err != nil {
			return order, err
		}
	}

	if strings.TrimSpace(order.Email) == "" {
		if !allowPrompt {
			return order, fmt.Errorf("email is required")
		}
		order.Email, err = prompt(reader, out, "Email")
		if err != nil {
			return order, err
		}
	}

	if order.Packs == 0 {
		if !allowPrompt {
			return order, fmt.Errorf("packs must be greater than 0")
		}
		order.Packs, err = promptUint(reader, out, "Packs")
		if err != nil {
			return order, err
		}
	}

	return order, nil
}

func prompt(reader *bufio.Reader, out io.Writer, label string) (string, error) {
	fmt.Fprintf(out, "%s: ", style(out, label, ansiBold, ansiCyan))
	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(value), nil
}

func promptUint(reader *bufio.Reader, out io.Writer, label string) (uint, error) {
	for {
		value, err := prompt(reader, out, label)
		if err != nil {
			return 0, err
		}

		parsed, err := strconv.ParseUint(value, 10, 64)
		if err == nil {
			return uint(parsed), nil
		}

		fmt.Fprintln(out, style(out, "Please enter a whole number.", ansiYellow))
	}
}

func promptYesNo(reader *bufio.Reader, out io.Writer, label string) (bool, error) {
	answer, err := prompt(reader, out, label+" [y/N]")
	if err != nil {
		return false, err
	}

	answer = strings.ToLower(strings.TrimSpace(answer))
	return answer == "y" || answer == "yes", nil
}
