// Copyright 2019 Andrei Tudor Călin
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package hiername_test

import (
	"testing"

	"acln.ro/hiername"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		segments []string
		want     string
	}{
		{
			name:     "",
			segments: []string{"hello"},
			want:     "hello",
		},
		{
			name:     "",
			segments: []string{"x", "y"},
			want:     "x.y",
		},
		{
			name:     "foo",
			segments: []string{"bar"},
			want:     "foo.bar",
		},
		{
			name:     "foo",
			segments: []string{"bar", "baz"},
			want:     "foo.bar.baz",
		},
	}
	for _, tt := range tests {
		got := hiername.Append(tt.name, tt.segments...)
		if got != tt.want {
			t.Errorf("Append(%q, %q) == %q, want %q", tt.name, tt.segments, got, tt.want)
		}
	}
}

func TestSegments(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "",
			want: []string{""},
		},
		{
			name: "foo",
			want: []string{"foo"},
		},
		{
			name: "foo.bar",
			want: []string{"foo", "bar"},
		},
		{
			name: "foo.bar.baz",
			want: []string{"foo", "bar", "baz"},
		},
	}
	for _, tt := range tests {
		got := hiername.Segments(tt.name)
		for i, seg := range got {
			if seg != tt.want[i] {
				t.Errorf("Segments(%q): segment %d == %q, want %q", tt.name, i, seg, tt.want[i])
			}
		}
	}
}
