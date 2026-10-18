// SPDX-FileCopyrightText: © 2024 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package strutil

// from "unicode"
const (
	MaxRune         = '\U0010FFFF' // Maximum valid Unicode code point.
	BOM             = '\uFFFE'     // Byte-Order Marker (LE vs BE)
	ReplacementChar = '\uFFFD'     // Represents invalid code points.
	MaxASCII        = '\u007F'     // maximum ASCII value.
	MaxLatin1       = '\u00FF'     // maximum Latin-1 value.
	RuneSelf        = 0x80         // from unicode/utf8 -- remove dependency
	//
	RuneError = '\uFFFD' // the "error" Rune or "Unicode replacement character"
)

const (
	UTF8_ReplacementChar = "\xef\xbf\xbd"
	UTF8_BOM = "\xef\xbb\xbf"
)

// Code points in the surrogate range are not valid for UTF-8.
const (
	surrogateMin = 0xD800
	surrogateMax = 0xDFFF
)
