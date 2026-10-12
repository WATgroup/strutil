// SPDX-FileCopyrightText: © 2021-2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package strutil

//go:generate go clean
//go:generate rm -f tables_generated.go
//go:generate go run tablegen.go -- tables_generated.go

func IsDigit(c rune) bool {
	return (c >= '0' && c <= '9')
}

func IsAsciiAlpha(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a') && (c <= 'z')
}

func IsAsciiAlnum(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a') && (c <= 'z')
}

func IsAsciiIdentBeg(c rune) bool {
	return ('_' == c) || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func IsAsciiIdentMid(c rune) bool {
	return (c >= '0' && c <= '9') || (c == '_') || (c >= 'A' && c <= 'Z') || (c >= 'a') && (c <= 'z')
}
