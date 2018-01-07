// Copyright 2014 loolgame Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package room

import "testing"

func assertEqual(t *testing.T, val interface{}, exp interface{}) {
	if val != exp {
		t.Errorf("Expected %v, got %v.", exp, val)
	}
}
func TestBuildBigRoomId(t *testing.T) {
	bigroom := BuildBigRoomId(9, 102, 11)
	assertEqual(t, bigroom, "000901020011")

	bigroom1 := BuildBigRoomId(9001, 1020, 11)
	assertEqual(t, bigroom1, "900110200011")
}
