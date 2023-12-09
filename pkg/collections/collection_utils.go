// Copyright Ⓒ 2023 Pavlo Moisieienko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package collections contains collections and their manipulation, including utility functions.
package collections

// CopyMap creates and return copy of the original map
func CopyMap[K comparable, V any](originalMap map[K]V) map[K]V {
	result := make(map[K]V, len(originalMap))
	for k, v := range originalMap {
		result[k] = v
	}
	return result
}
