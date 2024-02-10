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

package retryutil

import (
	"context"
	"fmt"
	"math"
	"time"
)

var RetryAbleErr = fmt.Errorf("retry")
var TimeoutErr = fmt.Errorf("timeout")

func RetryUntilTimeout(ctx context.Context, interval time.Duration, timeout time.Duration, do func() error) error {
	err := do()
	if err == nil {
		return nil
	}

	if err != RetryAbleErr {
		return err
	}

	if timeout == 0 {
		timeout = time.Duration(math.MaxInt64)
	}

	t := time.NewTimer(timeout)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			return TimeoutErr
		case <-time.After(interval):
			err := do()
			if err == nil {
				return nil
			}

			if err != RetryAbleErr {
				return err
			}
		}
	}
}
