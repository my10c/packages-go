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
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

// Function to get the MD5 value of the given file
func FileMD5(filePath string) (string, error) {
	var md5Str string
	// Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return md5Str, err
	}
	// Tell the program to call the following function when the current function returns
	defer file.Close()
	//Open a new hash interface to write to
	hash := md5.New()
	// Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		log.Printf("%s\n", err.Error())
		return md5Str, err
	}
	// Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	// Convert the bytes to a string
	md5Str = hex.EncodeToString(hashInBytes)
	return md5Str, nil
}

// Function to get the MD5 value of the given file pointer
func FilePtrMD5(filePtr *os.File) (string, error) {
	var md5Str string
	//Open a new hash interface to write to
	hash := md5.New()
	// Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, filePtr); err != nil {
		log.Printf("%s\n", err.Error())
		return md5Str, err
	}
	// Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	// Convert the bytes to a string
	md5Str = hex.EncodeToString(hashInBytes)
	return md5Str, nil
}
