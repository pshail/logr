/*
Copyright 2020 The logr Authors.

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

package logr

// Discard returns a Logger that discards all messages logged to it.  It can be
// used whenever the caller is not interested in the logs.  Logger instances
// produced by this function always compare as equal.
func Discard() Logger {
	return Logger{
		level: 0,
		Sink:  discardLogSink{},
	}
}

// discardLogSink is a LogSink that discards all messages.
type discardLogSink struct{}

// Verify that it actually implements the interface
var _ LogSink = discardLogSink{}

func (l discardLogSink) Init(RuntimeInfo) {
}

func (l discardLogSink) Enabled(int) bool {
	return false
}

func (l discardLogSink) Info(int, string, ...interface{}) {
}

func (l discardLogSink) Error(error, string, ...interface{}) {
}

func (l discardLogSink) WithValues(...interface{}) LogSink {
	return l
}

func (l discardLogSink) WithName(string) LogSink {
	return l
}
