// SPDX-FileCopyrightText: © 2021-2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package strutil_test

import (
	// "fmt"
	"github.com/WATgroup/assert"
	. "github.com/WATgroup/strutil"
	"testing"
)

func TestAlnum(t *testing.T) {

	assert.True(t, IsAlnum("q1w2e3r4t5y6"))

}

func TestAlpha(t *testing.T) {

	assert.True(t, IsAlpha("zxcvbn"))
	assert.False(t,IsAlpha("1qaz2wsx3edc4rfv5tgb6yhn7ujm8ik,"))
}

