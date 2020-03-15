//Copyright 2020 Google LLC
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//https://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package internal

var (
	basic  func(error) string
	detail func(error) string
)

// SetBasic modifies how the `Error() string` method serialises zerrors. Note this should not rely on the `Error()`
// method of zerrors or it may enter an infinite recursion.
func SetBasic(f func(error) string) {
	if f == nil {
		return
	}
	basic = f
}

// Basic serialises an error using the basic encoder function. It is used to produce `Error() string` for zerrors.
func Basic(err error) string {
	return basic(err)
}

// SetDetail modifies how the `zerrors.Detail(error) string` method serialises zerrors. Note this should not rely on the
// `Detail(error)` function of zerrors or it may enter an infinite recursion.
func SetDetail(f func(error) string) {
	if f == nil {
		return
	}
	detail = f
}

// Detail serialises an error using the detail encoder function. It is used to produce `zerrors.Detail(error) string`.
func Detail(err error) string {
	return detail(err)
}
