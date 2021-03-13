package main

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", errors.New(fmt.Sprintf("%s not found", key))
	}
	return val, nil
}