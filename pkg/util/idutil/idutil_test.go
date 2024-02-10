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

package idutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestGetUUID(t *testing.T) {
//	fmt.Println(GetUUID(""))
//}

func TestGetUUID36(t *testing.T) {
	fmt.Println(GetUUID36(""))
}

func TestGetManyUuid(t *testing.T) {
	for i := 0; i < 10000; i++ {
		testID := GetUUID36("")
		if len(testID) != 14 {
			t.Errorf("GetUUID failed, expected uuid length 14, got: %d", len(testID))
		}
	}
}

func TestRandString(t *testing.T) {
	str := randString(Alphabet62, 50)
	assert.Equal(t, 50, len(str))
	t.Log(str)

	str = randString(Alphabet62, 255)
	assert.Equal(t, 255, len(str))
	t.Log(str)
}
