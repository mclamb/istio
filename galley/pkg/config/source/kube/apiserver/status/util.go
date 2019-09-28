// Copyright 2019 Istio Authors
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

package status

import (
	"strings"

	"istio.io/istio/galley/pkg/config/analysis/diag"
)

// toStatusValue converts a set of diag.Messages to a status value.
func toStatusValue(msgs diag.Messages) interface{} {
	if len(msgs) == 0 {
		return nil
	}

	var lines strings.Builder

	for _, m := range msgs {
		lines.WriteString(m.StatusString())
		lines.WriteString("\n")
	}

	return lines.String()
}