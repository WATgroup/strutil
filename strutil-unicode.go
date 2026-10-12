// SPDX-FileCopyrightText: © 2021-2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

//go:build !generate

package strutil

import (
	"unicode"
)

func IsAlpha(s string) bool {
	for _, r := range s { // runes
		if r > RuneSelf {
			return false
		}
		if /* r < RuneSelf && */ !tab_Alpha[r] { // !IsAsciiAlpha(r)
			return false
		}
	}
	return true
}


// IsAlnum checks whether a (UNICODE) string is alphanumeric only
func IsAlnum(s string) bool {

	for _, r := range s { // runes

		if r < RuneSelf {
			if !tab_Alnum[r] { // !IsAsciiAlnum(r)
				return false
			}
		} else {
			// previous check must have caught digits....
			if !unicode.IsLetter(r) {
				return false
			}
		}
	}
	return true
}


// IsIdentPlus checks whether string is a valid (extended!) identifier
func IsIdentPlus(s string) bool {

	for i, r := range s { // runes
		if '_' == r { // always good, including "just _"
			continue
		}
		switch {
		case unicode.IsLetter(r): // always good
			continue

		case 0 == i && IsDigit(r):
			return false // Boo!

		case unicode.IsSymbol(r):
			continue

		default:
			return false
		} // end-switch
	}
	return true
}
