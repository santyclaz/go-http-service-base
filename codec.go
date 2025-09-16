package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func encode[T any](
	w http.ResponseWriter,
	r *http.Request,
	status int,
	v T,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

type ValidatorFunc[T any] func(o T) (problems ProblemsMap)
type ProblemsMap map[string]string

func decode[T any](
	r *http.Request,
	validatorFn ValidatorFunc[T],
) (T, ProblemsMap, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, nil, fmt.Errorf("decode json: %w", err)
	}

	if validatorFn == nil {
		validatorFn = func(o T) ProblemsMap { return nil }
	}
	if problems := validatorFn(v); len(problems) > 0 {
		return v, problems, fmt.Errorf("invalid %T: %d problems", v, len(problems))
	}

	return v, nil, nil
}
