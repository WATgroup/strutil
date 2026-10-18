// SPDX-FileCopyrightText: © 2021,2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package strutil_test

import (
	// "fmt"
	"github.com/WATgroup/assert"
	. "github.com/WATgroup/strutil"
	"testing"
)

var bm1 = BuildSpanBM_A("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var bm2 = BuildSpanBM_A("abcdefghijklmnopqrstuvwxyz")

func TestBuildTable(t *testing.T) {

	/*t1*/_ = BuildSpanBM_A("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// fmt.Printf("%016x\n", t1)

	/*t2*/_ = BuildSpanBM_A("abcdefghijklmnopqrstuvwxyz")
	// fmt.Printf("%016x\n", t2)
}

func TestStrspnT(t *testing.T) {

	t.Run("uppercase", func(t *testing.T) {
		a := StrspnT(&bm1, "PEPEPOTAMO")
		// fmt.Println(a)
		assert.True(t, 0 == a)
	})
	t.Run("uppercaseC", func(t *testing.T) {
		a := StrcspnT(&bm1, "pepepotamo")
		// fmt.Println(a)
		assert.True(t, 10 == a) // =len("pepepotamo")
	})
	t.Run("lowercase", func(t *testing.T) {
		b := StrspnT(&bm2, "maguilagorilla")
		// fmt.Println(b)
		assert.True(t, 0 == b)
	})
	t.Run("lowercaseC", func(t *testing.T) {
		b := StrcspnT(&bm2, "GORILLAMAGILA")
		// fmt.Println(b)
		assert.True(t, 13 == b)
	})
}
