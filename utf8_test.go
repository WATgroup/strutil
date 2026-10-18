// SPDX-FileCopyrightText: © 2021,2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package strutil_test

import (
	"github.com/WATgroup/assert"
	. "github.com/WATgroup/strutil"
	"testing"
)



func testUTF8str(t *testing.T, ss []string) {
	for _, s := range ss {
		assert.True(t, ValidUTF8String(s))
	}
}

func TestUTF8set(t *testing.T) {

	// Check regular ASCII
	assert.True(t, ValidUTF8String("abcdefgh"))

	// Unicode UTF8 tests
	// Reference: https://www.w3.org/2001/06/utf-8-test/UTF-8-demo.html
	t.Run("Unicode scientific", func(t *testing.T){
		uniText0 := []string {
			"∮ E⋅da = Q,  n → ∞, ∑ f(i) = ∏ g(i), ∀x∈ℝ: ⌈x⌉ = −⌊−x⌋, α ∧ ¬β = ¬(¬α ∨ β),",
			"ℕ ⊆ ℕ₀ ⊂ ℤ ⊂ ℚ ⊂ ℝ ⊂ ℂ, ⊥ < a ≠ b ≡ c ≤ d ≪ ⊤ ⇒ (A ⇔ B),",
			"2H₂ + O₂ ⇌ 2H₂O, R = 4.7 kΩ, ⌀ 200 mm",
		}
		testUTF8str(t, uniText0)
	})
	t.Run("Unicode linguistics", func(t *testing.T){
		uniText1 := []string {
			"ði ıntəˈnæʃənəl fəˈnɛtık əsoʊsiˈeıʃn",
			"Y [ˈʏpsilɔn], Yen [jɛn], Yoga [ˈjoːgɑ]",
			"((V⍳V)=⍳⍴V)/V←,V    ⌷←⍳→⍴∆∇⊃‾⍎⍕⌈,",
		}
		testUTF8str(t, uniText1)
	})
	t.Run("Unicode misc", func(t *testing.T){
		uniText := []string {
			`“We’ve been here”`,
			`‚deutsche‘ „Anführungszeichen“ `,
			`†, ‡, ‰, •, 3–4, —, −5/+5, ™, …`,
			`1lI|, 0OD, 8B €æßðđŋħ`,
		}
		testUTF8str(t, uniText)
	})

	// Emojis interspersed with some characters
	assert.True(t, ValidUTF8String("💩abc🔴🟡🟢xyz🕗Ç🏷ñ🏠"))
}
