/*
Copyright © 2024 David Muckle <dvdmuckle@dvdmuckle.xyz>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package helper

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang/glog"
)

var verboseErrLog bool

func GetVerboseErrLogAddr() *bool {
	return &verboseErrLog
}

func LogErrorAndExit(args ...interface{}) {
	if verboseErrLog {
		glog.Fatal(args...)
	} else {
		logArgs := []interface{}{"Error: "}
		logArgs = append(logArgs, args...)
		fmt.Println(strings.Trim(fmt.Sprint(logArgs...), "[]"))
		os.Exit(1)
	}
}
