// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Peter Mattis (petermattis@gmail.com)

package util

var (
	buildSHA  string // SHA (or build revision)
	buildTag  string // Tag of this build
	buildUser string // User that built the artifact
	buildTime string // Unix time since EPOCH (string, in seconds)
)

// BuildInfo ...
type BuildInfo struct {
	SHA  string `json:"sha"`
	Tag  string `json:"tag"`
	User string `json:"user"`
	Time string `json:"time"`
}

// GetBuildInfo ...
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		SHA:  buildSHA,
		Tag:  buildTag,
		User: buildUser,
		Time: buildTime,
	}
}
