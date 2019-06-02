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

// Package hiername provides hierarchical "."-separated names for use with
// logging and metrics collection systems.
package hiername

import "strings"

// Sep is the name separator.
const Sep = "."

// Append appends the specified segments to the specified name. For example:
//
// 	Append("", "apiserver") == "apiserver"
//	Append("", "x", "y") == "x.y"
//	Append("apiserver", "proxy") == "apiserver.proxy"
//	Append("foo", "bar", "baz") == "foo.bar.baz"
func Append(name string, segments ...string) string {
	if name == "" {
		return strings.Join(segments, Sep)
	}
	return name + Sep + strings.Join(segments, Sep)
}

// Segments returns the segments in a hierarchical name.
func Segments(name string) []string {
	return strings.Split(name, Sep)
}
