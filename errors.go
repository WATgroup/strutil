// SPDX-FileCopyrightText: © 2024 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package strutil

// Refactor to avoid dependencies on "WATgroup/errors"
// ...given that this intends to be a "leaf" / freestanding package
type strUtilError string

func (e strUtilError) Error() string {
	return string(e)
}
