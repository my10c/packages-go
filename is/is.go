// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package is

import (
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"os"
	"os/user"
	"path"
	"runtime"
	"strconv"
	"syscall"
)

type (
	Is struct{}
)

func getFileStat(file, infoType string) (string, error) {
	var msg string
	var object bool
	fileInfo, err := os.Stat(file)
	if err != nil {
		return "ERROR", err
	}

	if infoType == "file" || infoType == "directory" {
		dirPath, _ := path.Split(file)
		if _, err := os.Stat(dirPath); err != nil {
			return "ERROR", err
		}
		object = fileInfo.IsDir()
	}

	switch infoType {
	case "permissions":
		msg = fmt.Sprintf("%04o", fileInfo.Mode().Perm())

	case "owner":
		fileState := fileInfo.Sys().(*syscall.Stat_t)
		userId := strconv.FormatUint(uint64(fileState.Uid), 10)
		userName, err := user.LookupId(userId)
		if err != nil {
			return "ERROR", err
		}
		msg = userName.Username

	case "file", "directory":
		if object {
			msg = "directory"
		} else {
			msg = "file"
		}

	}
	return msg, err
}

func New() *Is {
	return &Is{}
}

func (i *Is) IsRoot() bool {
	if os.Geteuid() == 0 {
		return true
	}
	return false
}

func (i *Is) IsUser(runUser string) (bool, error) {
	user, err := user.Current()
	if user.Username == runUser {
		return true, err
	}
	return false, err
}

func (i *Is) IsOS() string {
	return runtime.GOOS
}

func (i *Is) IsRunning(process string, pid int) (bool, error) {
	procInfo, err := ps.FindProcess(pid)
	if procInfo == nil {
		return false, err
	}
	if procInfo.Executable() == process && procInfo.Pid() == pid {
		return true, err
	}
	return false, err
}

func (i *Is) IsRunningUser() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "ERROR", err
	}
	return user.Username, err
}

func (i *Is) IsFileOwner(file string, owners []string) (string, bool) {
	var foundUser bool = false
	fileOwner, err := getFileStat(file, "owner")
	if err != nil {
		return err.Error(), false
	}

	for _, owner := range owners {
		if fileOwner == owner {
			foundUser = true
			break
		}
	}
	if !foundUser {
		return fmt.Sprintf("ERROR: The file %s is not own by %v, is own by %s", file, owners, fileOwner), foundUser
	}
	return fmt.Sprintf("OK: The file %s is own by %s", file, fileOwner), foundUser
}

func (i *Is) IsFilePermission(file string, permissions []string) (string, bool) {
	var foundPermission bool = false
	filePerm, err := getFileStat(file, "permissions")
	if err != nil {
		return err.Error(), false
	}

	for _, filePermission := range permissions {
		if filePermission == filePerm {
			foundPermission = true
			break
		}
	}
	if !foundPermission {
		return fmt.Sprintf("ERROR: The file %s permission are to wide open %s", file, filePerm), foundPermission
	}
	return fmt.Sprintf("OK: The file %s has the correct permissions  %s", file, filePerm), foundPermission
}

func (i *Is) IsExist(fullPath, fileType string) (string, bool, error) {
	fileInfo, err := getFileStat(fullPath, fileType)
	if fileInfo == fileType {
		return fileInfo, true, err
	}
	return fileInfo, false, err
}

func (i *Is) IsInList(slice []string, val string) (bool) {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
