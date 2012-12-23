// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This LevelDB Go implementation is based on LevelDB C++ implementation.
// Which contains the following header:
//   Copyright (c) 2011 The LevelDB Authors. All rights reserved.
//   Use of this source code is governed by a BSD-style license that can be
//   found in the LEVELDBCPP_LICENSE file. See the LEVELDBCPP_AUTHORS file
//   for names of contributors.

package table

import "testing"

func TestIter_Empty(t *testing.T) {
	cc := []struct {
		name string
		c    Constructor
	}{
		{"block", &BlockConstructor{}},
		{"table", &TableConstructor{}},
		{"memdb", &MemDBConstructor{}},
		{"merged", &MergedMemDBConstructor{}},
	}

	for _, p := range cc {
		c, name := p.c, p.name
		err := c.Init()
		if err != nil {
			t.Error(name+": error when initializing constructor:", err.Error())
			continue
		}
		size, err := c.Finish(t)
		if err != nil {
			t.Error(name+": error when finishing constructor:", err.Error())
			continue
		}
		t.Logf(name+": final size is %d bytes", size)
		iter := c.NewIterator()
		if iter.Valid() {
			t.Error(name + ": Valid() return true")
		}
		if iter.Next() {
			t.Error(name + ": Next() return true")
		}
		if iter.Prev() {
			t.Error(name + ": Prev() return true")
		}
		if iter.Seek(nil) {
			t.Error(name + ": Seek(nil) return true")
		}
		if iter.Seek([]byte("v")) {
			t.Error(name + ": Seek('v') return true")
		}
		if iter.First() {
			t.Error(name + ": First() return true")
		}
		if iter.Last() {
			t.Error(name + ": Last() return true")
		}
	}
}