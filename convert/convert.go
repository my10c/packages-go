// Copyright (c) 2017 Badassops
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//	* Redistributions of source code must retain the above copyright
//	notice, this list of conditions and the following disclaimer.
//	* Redistributions in binary form must reproduce the above copyright
//	notice, this list of conditions and the following disclaimer in the
//	documentation and/or other materials provided with the distribution.
//	* Neither the name of the <organization> nor the
//	names of its contributors may be used to endorse or promote products
//	derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSEcw
// ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Author		:	Luc Suryo <luc@badassops.com>
//
// Version		:	0.1
//
// Date			:	Jul 11, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Jul 11, 2017	LIS			Release as separate package
//
// TODO:

package utils

import (
	"bytes"
	"strings"
)

// Function to create a slice from the given strings
func StringsToSlice(args ...string) []string {
	return strings.Fields(strings.Join(args, " "))
}

// Function to create a string from the given slice
func SliceToString(array []*string) string {
	var buffer bytes.Buffer
	for idx := range array {
		buffer.WriteString(*array[idx])
	}
	return buffer.String()
}

// function to convert network speed from bytes/sec to X/sec
func BytesToBW(bytes uint64, unit string) float64 {
	// speed unit is 1000 and not 1024 for network speed!
	var converted float64 = 0
	switch unit {
	case "KB":
		converted = float64(bytes) / float64(1000)
	case "MB":
		converted = float64(bytes) / float64(1000*1000)
	case "GB":
		converted = float64(bytes) / float64(1000*1000*1000)
	case "TB":
		converted = float64(bytes) / float64(1000*1000*1000*1000)
	}
	return converted
}
