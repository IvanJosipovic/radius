/*
Copyright 2023 The Radius Authors.

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

package v20220315privatepreview

import (
	"bytes"
	"io"
	"os"
	"time"
)

func loadTestData(testfile string) ([]byte, error) {
	d, err := os.ReadFile(testfile)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func loadTestDataAsReader(testfile string) (io.Reader, error) {
	d, err := os.ReadFile(testfile)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(d), nil
}

func unmarshalTimeString(ts string) *time.Time {
	var tt timeRFC3339
	_ = tt.UnmarshalText([]byte(ts))
	return (*time.Time)(&tt)
}
