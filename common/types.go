/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package common

import (
	"errors"
)

var (
	// ErrCantCompareInterface error can't compare interface
	ErrCantCompareInterface = errors.New("Can't compare interface")
	// ErrFieldWrongType error field has wrong type
	ErrFieldWrongType = errors.New("Field has wrong type")
	// ErrNotFound error no result was found
	ErrNotFound = errors.New("No result found")
	// ErrTimeout network timeout
	ErrTimeout = errors.New("Timeout")
	// ErrNotImplemented unimplemented feature
	ErrNotImplemented = errors.New("Not implemented")
)

// SortOrder describes ascending or descending order
type SortOrder string

const (
	// SortAscending sorting order
	SortAscending SortOrder = "ASC"
	// SortDescending sorting order
	SortDescending SortOrder = "DESC"
)

// Metric defines a common metric interface
type Metric interface {
	// part of the Getter interface
	GetFieldInt64(field string) (int64, error)
	GetFieldKeys() []string

	Add(m Metric) Metric
	Sub(m Metric) Metric
	Split(cut int64) (Metric, Metric)
	GetStart() int64
	SetStart(start int64)
	GetLast() int64
	SetLast(last int64)
	IsZero() bool
}
