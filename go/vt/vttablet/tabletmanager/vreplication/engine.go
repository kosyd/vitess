/*
Copyright 2018 The Vitess Authors.

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

package vreplication

import (
	"sync"

	"vitess.io/vitess/go/vt/mysqlctl"
)

// Engine is the engine for handling vreplication.
type Engine struct {
	mu     sync.Mutex
	isOpen bool

	mysqld *mysqlctl.MysqlDaemon
}

// NewEngine creates a new Engine.
func NewEngine(mysqld *mysqlctl.MysqlDaemon) *Engine {
	return &Engine{
		mysqld: mysqld,
	}
}

// Open starts the Engine service.
func (vre *Engine) Open() error {
	if vre.isOpen {
		return nil
	}
	vre.isOpen = true
	return nil
}

// Close closes the Engine service.
func (vre *Engine) Close() {
	vre.mu.Lock()
	defer vre.mu.Unlock()
	if !vre.isOpen {
		return
	}
	vre.isOpen = false
}
