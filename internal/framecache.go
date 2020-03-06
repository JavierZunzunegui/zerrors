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

// Package internal contains stateful global variables internal to zerrors.
package internal

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	once    sync.Once
	cache   atomic.Value              // The underlying type is *map[uintptr]runtime.Frame
	current map[uintptr]runtime.Frame // The same as cache above, but split to avoid atomic access. Read-only.
	next    map[uintptr]runtime.Frame // Eventually becomes current. It is read-write and not shared.
	updates chan struct {
		pc    uintptr
		frame runtime.Frame
	} // The means by which new entries make it into the cache.

	updatesCount = 0
)

const (
	// Used to control how often 'next' is promoted to 'current'.
	// The value is arbitrary.
	updatesBeforeCacheSwitch = 10
)

func initCache() {
	updates = make(chan struct {
		pc    uintptr
		frame runtime.Frame
	}, 10)
	storeCache(make(map[uintptr]runtime.Frame))

	go runCacheLoop()
}

func storeCache(v map[uintptr]runtime.Frame) {
	current = v
	cache.Store(&v)
}

func runCacheLoop() {
	for update := range updates {
		// Checking 'current' not 'next' so that repeated use of the same frame does eventually lead to the promotion of 'next'.
		if _, ok := current[update.pc]; ok {
			continue
		}

		if next == nil {
			next = make(map[uintptr]runtime.Frame, len(current)+updatesBeforeCacheSwitch)
			for k, v := range current {
				next[k] = v
			}
		}

		next[update.pc] = update.frame

		updatesCount++

		if updatesCount > updatesBeforeCacheSwitch {
			storeCache(next)
			next = nil
			updatesCount = 0
		}
	}
}

// GetFramer returns a Framer with the latest global runtime.Frame cache.
func GetFramer() Framer {
	return *(cache.Load().(*map[uintptr]runtime.Frame))
}

// Framer is a runtime.Frame cache.
type Framer map[uintptr]runtime.Frame

// Get gets a frame from the cache, and if not found it evaluates it, stores it in the cache and returns.
func (f Framer) Get(pc uintptr) runtime.Frame {
	if frame, ok := f[pc]; ok {
		return frame
	}

	callers := [1]uintptr{pc}
	frame, _ := runtime.CallersFrames(callers[:]).Next()

	select {
	case updates <- struct {
		pc    uintptr
		frame runtime.Frame
	}{pc, frame}:
	default:
		// Do not block.
	}

	return frame
}
