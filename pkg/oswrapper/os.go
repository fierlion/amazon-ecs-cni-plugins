// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package oswrapper

import "os"

// OS wraps methods from the 'os' package
type OS interface {
	FindProcess(pid int) (OSProcess, error)
}

// OSProcess wraps methods from the 'os.Process' struct
type OSProcess interface {
	Kill() error
}

type _os struct {
}

// NewOS creates a new OS object
func NewOS() OS {
	return &_os{}
}

func (*_os) FindProcess(pid int) (OSProcess, error) {
	return os.FindProcess(pid)
}