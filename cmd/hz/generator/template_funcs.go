/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package generator

import (
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/cloudwego/hertz/cmd/hz/util"
)

var funcMap = func() template.FuncMap {
	m := template.FuncMap{
		"GetUniqueHandlerOutDir": getUniqueHandlerOutDir,
		"ToSnakeCase":            util.ToSnakeCase,
		"Split":                  strings.Split,
		"Trim":                   strings.Trim,
		"EqualFold":              strings.EqualFold,
		"BaseName":               util.BaseName,
		"IsRedundantAlias":       isRedundantAlias,
	}
	for key, f := range sprig.TxtFuncMap() {
		m[key] = f
	}
	return m
}()

// getUniqueHandlerOutDir uses to get unique "api.handler_path"
func getUniqueHandlerOutDir(methods []*HttpMethod) (ret []string) {
	outDirMap := make(map[string]string)
	for _, method := range methods {
		if _, exist := outDirMap[method.OutputDir]; !exist {
			outDirMap[method.OutputDir] = method.OutputDir
			ret = append(ret, method.OutputDir)
		}
	}

	return ret
}

// isRedundantAlias checks if the alias is redundant (same as package name)
func isRedundantAlias(alias, packagePath string) bool {
	packageName := util.BaseName(packagePath, "")

	// Check if alias is the same as the original package name
	if alias == packageName {
		return true
	}

	// Check if alias is the same as the camelCase version of package name
	camelCaseAlias := util.ToCamelCaseAlias(packageName)
	if alias == camelCaseAlias {
		return true
	}

	return false
}
