// Copyright 2016 Richard Hawkins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package pyxel contains utility functions for reading data from pyxel files.
package pyxel

// Context of a .pyxel file
type Context struct {
	Version string
}

// Load returns a pointer to a Context of the pyxel file.
func Load(path string) (*Context, error) {
	ctx := Context{}
	return &ctx, nil
}
