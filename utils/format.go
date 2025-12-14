package utils

import (
	"errors"
	"fmt"
)

var ErrCategoryNotFound = errors.New("category not found")
var ErrItemNotFound = errors.New("item not found")

func ErrorMessage(err error) string {
	return fmt.Sprintf("\033[31m[FAILED] Something went wrong: %v\033[0m", err.Error())
}

func SuccessMsg(add string) string {
	return fmt.Sprintf("\033[32m[SUCCESS] %s successfully\033[0m", add)
}

func PrintSuccess(add string) {
	fmt.Printf("\033[32m[SUCCESS] %s successfully\033[0m\n", add)
}

func PrintErr(err error) {
	fmt.Printf("\033[31m[FAILED] Something went wrong: %v\033[0m", err.Error())
}