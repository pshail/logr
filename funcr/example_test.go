/*
Copyright 2021 The logr Authors.

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

package funcr_test

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
)

func ExampleUnderlier() {
	var log logr.Logger = funcr.New(func(prefix, args string) {
		fmt.Println(prefix, args)
	}, funcr.Options{})

	if underlier, ok := log.Sink.(funcr.Underlier); ok {
		fn := underlier.GetUnderlying()
		fn("hello", "world")
	}
	// Output: hello world
}
