// SPDX-FileCopyrightText: © 2021,2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package strutil

import (
// 	"unicode/utf8"
)

//revive:disable-next-line:var-naming
const k_spnThreshold = 8

/** strspn
 * Implementation loosely inspired by various implementations found
 * in the public domain (or assumed to be PD)
 * ...only modernized and independently ported to go
 */
func Strspn(s string, charset string) (size_t, error) {
	if 0 == len(charset) {
		return 0, nil
	}
	for i := range len(s) {
		if s[i] >= RuneSelf {
			// It's not ASCII...
		}
	}
	if len(s) <= k_spnThreshold {
		// Doesn't make sense to even build the table
		return strspnA(s, charset), nil
	}

	// fastpath: use StrspnA
	if r := StrspnA([]byte(s), charset); r >= 0 {
		// All good !!
		return size_t(r), nil
	}
	// strspnA returned -1
	return 0, strUtilError(`Encoding error or preconditions not met`)
}

func StrspnA(s []byte, charset string) ssize_t {
	if 0 == len(charset) {
		return 0
	}
	if len(charset) < k_spnThreshold {
		return ssize_t(strspnA(string(s), charset))
	}
	bm := BuildSpanBM_A(charset)
	var i int
	for i = 0; i < len(s); i++ {
		if !_opTEST(&bm, s[i]) {
			return ssize_t(i)
		}
	} // end-for
	return -1
}

// ~strpbrk
func StrpA(s string, accept string) size_t {
	var bm AsciiBitmap
	buildTableA(&bm, accept)
	return size_t(StrcspnT(&bm, s))
}

func StrspnT(tab *AsciiBitmap, s string) int {
	sd := []byte(s)
	var i int
	for ; i < len(sd); i++ {
		if !_opTEST(tab, sd[i]) {
			break
		}
	}
	return len(sd) - i
}

func StrcspnT(tab *AsciiBitmap, s string) int {
	sd := []byte(s)
	var i int
	for ; i < len(sd); i++ {
		if _opTEST(tab, sd[i]) {
			break
		}
	}
	return i
}

type AsciiBitmap = [4]uint64 // ~[32]byte, but more efficient

func BuildSpanBM_A(charset string) (ret AsciiBitmap) {
	buildTableA(&ret, charset)
	return
}

func MergeSpanBM_A(bitmap *AsciiBitmap, charset string) bool {
	buildTableA(bitmap, charset)
	return true
}

func buildTableA(res *AsciiBitmap, charset string) {
	// expecting the compiler to notice it won't be modified... and avoid the copy
	cs := []byte(charset)
	for i := 0; i < len(cs); i++ {
		_opSET(res, cs[i])
	}
}

// slow (quadratic) implementation
// ...but it's actually bounded (<=8) iterations
func strspnA(s, charset string) size_t {
	// lc := len(charset)
	var i int
	var c rune
	for i, c = range s { // runes !!
		// j := 0
		// loop over the charset, one rune at a time... :O
		for _, x := range charset {
			if x == c { // found, so loop
				goto strspnA_loop
			}
			// j++
		}
		// if j == lc { // not found, hence return early
		return size_t(i)
		// }
	strspnA_loop:
	}
	return size_t(i) // must be == lc if here
}

////////////////////////////////////////////////////////////////////////////////

// set the corresponding bit in the bitmap
func _opSET(ret *AsciiBitmap, x byte) {
	// the constant here is "64" bits in an uint64 ;)
	// hence, make it into bitops: SHR + AND
	// ret[ x / (8*Sizeof(uint64)) ] |= 1 << x % (8*Sizeof(uint64))
	ret[x>>6] |= 1 << (x & 0x3F)
}

// check if the bit is set
func _opTEST(ret *AsciiBitmap, x byte) bool {
	return 0 != (ret[x>>6] & (1 << (x & 0x3F)))
}

/*
 * Reference:	man 3 strspn
 * Reference:   musl libC's implementation
 *
 */
