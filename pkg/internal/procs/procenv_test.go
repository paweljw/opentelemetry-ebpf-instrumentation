// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package procs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvStrParsing(t *testing.T) {
	strs := []string{
		"ok=\"=  =\"",
		"nothing",
		"=wrong",
		"something=somethingelse",
		"something_empty=",
		"something= else",
		"weird==  =",
		"resources=a=b,c=d,e=  fg",
		"",
	}

	res := envStrsToMap(strs)
	assert.Equal(t, map[string]string{"something": "else", "ok": "\"=  =\"", "weird": "=  =", "resources": "a=b,c=d,e=  fg"}, res)
}

func BenchmarkEnvStrsToMap(b *testing.B) {
	vars := make([]string, 32)
	for i := range vars {
		vars[i] = fmt.Sprintf("KEY_%d=value_%d", i, i)
	}
	padded := make([]string, 8192)
	copy(padded, vars)

	for _, tc := range []struct {
		name string
		vars []string
	}{
		{name: "ordinary", vars: vars},
		{name: "nul_padded", vars: padded},
	} {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				envStrsToMap(tc.vars)
			}
		})
	}
}
