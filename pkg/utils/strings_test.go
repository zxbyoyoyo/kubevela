/*
 Copyright 2022 The KubeVela Authors.

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

package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSanitize(t *testing.T) {
	s := "abc\ndef\rgh"
	require.Equal(t, "abcdefgh", Sanitize(s))
}

func TestIgnoreVPrefix(t *testing.T) {
	require.Equal(t, "1.2.0", IgnoreVPrefix("v1.2.0"))
	require.Equal(t, "1.2.0", IgnoreVPrefix("1.2.0"))
}
