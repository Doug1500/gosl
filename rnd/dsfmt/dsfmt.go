// Copyright 2015 Dorival Pedroso and Raul Durand. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dsfmt

/*
#cgo CFLAGS: -O3 -finline-functions -fomit-frame-pointer -DNDEBUG -fno-strict-aliasing --param max-inline-insns-single=1800 -Wall -std=c99 -msse2 -DHAVE_SSE2 -DDSFMT_MEXP=19937
#include "connectdsfmt.h"
*/
import "C"

import "time"

// Init initialises random numbers generator
//  Input:
//   seed -- seed value; use seed <= 0 to use current time
func Init(seed int) {
	if seed <= 0 {
		seed = int(time.Now().Unix())
	}
	C.DsfmtInit(C.long(seed))
}

// Rand generates pseudo random real numbers between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit (closed)
//   high -- upper limit (open)
//  Output:
//   random float64
func Rand(low, high float64) float64 {
	return float64(C.DsfmtRand(C.double(low), C.double(high)))
}
