// Copyright 2024 The seacraft Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sliceutil

// RemoveString remove string from slice if function return true.
func RemoveString(slice []string, remove func(item string) bool) []string {
	for i := 0; i < len(slice); i++ {
		if remove(slice[i]) {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return slice
}

// FindString return true if target in slice, return false if not.
func FindString(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}

// FindInt return true if target in slice, return false if not.
func FindInt(slice []int, target int) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}

// FindUint return true if target in slice, return false if not.
func FindUint(slice []uint, target uint) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}
