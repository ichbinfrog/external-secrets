/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (

	// nolint:gosec
	"crypto/md5"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// MergeByteMap merges map of byte slices.
func MergeByteMap(dst, src map[string][]byte) map[string][]byte {
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// ConvertName converts a string into a secret-key compatible string.
// Replaces any non-alphanumeric characters with its unicode code.
func ConvertName(in string) string {
	out := make([]string, len(in))
	rs := []rune(in)
	for k, r := range rs {
		if !unicode.IsNumber(r) &&
			!unicode.IsLetter(r) &&
			r != '-' &&
			r != '.' &&
			r != '_' {
			out[k] = fmt.Sprintf("_U%04x_", r)
		} else {
			out[k] = string(r)
		}
	}
	return strings.Join(out, "")
}

// MergeStringMap performs a deep clone from src to dest.
func MergeStringMap(dest, src map[string]string) {
	for k, v := range src {
		dest[k] = v
	}
}

// IsNil checks if an Interface is nil.
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	value := reflect.ValueOf(i)
	if value.Type().Kind() == reflect.Ptr {
		return value.IsNil()
	}
	return false
}

// ObjectHash calculates md5 sum of the data contained in the secret.
// nolint:gosec
func ObjectHash(object interface{}) string {
	textualVersion := fmt.Sprintf("%+v", object)
	return fmt.Sprintf("%x", md5.Sum([]byte(textualVersion)))
}

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
