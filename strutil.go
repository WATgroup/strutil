// SPDX-FileCopyrightText: © 2021-2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

//go:build !generate

package strutil


func IsAlphaA(s string) bool {
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


// IsAlnumA checks whether an ASCII string is alphanumeric only
func IsAlnumA(s string) bool {
	for _, r := range s { // runes
		if r > RuneSelf {
			return false
		}
		if /* r < RuneSelf && */ !tab_Alnum[r] { // !IsAsciiAlnum(r)
			return false
		}
	}
	return true
}

func IsIdent(s string) bool {

	for i, r := range s { // runes
		if '_' == r { // always good, including "just _"
			continue
		}
		if r > RuneSelf {	// only Ascii here...
			return false
		}
		if 0 == i && IsDigit(r) {
			return false // Boo!
		}
		if !tab_Alnum[r] {
			return false
		}
	}
	return true
}
