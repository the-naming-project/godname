// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//  ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗
//  ████╗  ██║██╔═████╗██║   ██║██╔══██╗
//  ██╔██╗ ██║██║██╔██║██║   ██║███████║
//  ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║
//  ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║
//  ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝

package godname

// #cgo LDFLAGS: -ldname
//
// #include "dname.h"
import "C"

// DNameDigest Represents a dname_digest.
//
// typedef struct dname_digest {
//     char *input;
//     char *name;
//     unsigned char sha256hash[DNAME_SHA256_DIGEST_32];
// } dname_digest;
//
type DNameDigest struct {
	Input string
	Name  string
}

// DNameLookup will call the internal library to lookup the
// state of the machine at runtime.
func DNameLookup() *DNameDigest {
	raw := C.dname_lookup_raw()
	lookup := C.GoString(raw)
	return DName(lookup)
}

// DName can used to calculate a deterministic name given
// an arbitrary input.
func DName(input string) *DNameDigest {
	digest := C.dname(C.CString(input))
	return &DNameDigest{
		Input: input,
		Name:  C.GoString(digest.name),
	}
}
