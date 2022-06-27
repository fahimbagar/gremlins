/*
 * Copyright 2022 The Gremlins Authors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package mutator

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMutantTypeString(t *testing.T) {
	testCases := []struct {
		name       string
		mutantType MutantType
		expected   string
	}{
		{
			"CONDITIONALS_BOUNDARY",
			ConditionalsBoundary,
			"CONDITIONALS_BOUNDARY",
		},
		{
			"CONDITIONALS_NEGATION",
			ConditionalsNegation,
			"CONDITIONALS_NEGATION",
		},
		{
			"INCREMENT_DECREMENT",
			IncrementDecrement,
			"INCREMENT_DECREMENT",
		},
		{
			"INVERT_NEGATIVES",
			InvertNegatives,
			"INVERT_NEGATIVES",
		},
		{
			"ARITHMETIC_BASE",
			ArithmeticBase,
			"ARITHMETIC_BASE",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.mutantType.String() != tc.expected {
				t.Errorf(cmp.Diff(tc.mutantType.String(), tc.expected))
			}
		})
	}
}
