// +build go1.11

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * importer_go1_11.go
 *
 *  Created on Nov 16, 2019
 *      Author Massimiliano Ghilardi
 */

package genimport

import (
	"fmt"
	"go/build"
	"go/importer"
	"go/types"
	"os"
	"strings"

	"github.com/cosmos72/gomacro/base/paths"
	"golang.org/x/tools/go/packages"
)

const GoModuleSupported bool = true

func (imp *Importer) Load(path string, enableModule bool) (p *types.Package, err error) {
	if !enableModule {
		return importer.Default().Import(path)
	}

	imp.output.Debugf("looking for package %q ...", path)

	defer func() {
		if p == nil && err == nil {
			r := recover()
			if rerr, ok := r.(error); ok {
				err = rerr
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	// Go >= 1.14 requires a valid go.mod file in the directory used for packages.Config.Dir
	gomod := createPluginGoModFile(imp.output, path)

	cfg := packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedImports,
		Env:  environForCompiler(enableModule),
		Dir:  paths.DirName(gomod),
		Logf: nil, // imp.output.Debugf,
	}

	list, err := packages.Load(&cfg, "pattern="+path)
	if err != nil {
		return nil, err
	}
	for _, pkg := range list {
		if pkg.PkgPath == path {
			if len(pkg.Errors) != 0 {
				err = errorList{pkg.Errors, mergeErrorMessages(pkg.Errors)}
				return nil, err
			}
			return pkg.Types, nil
		}
	}
	return nil, fmt.Errorf("packages.Load() could not find package %q", path)
}

type errorList struct {
	errors []packages.Error
	str    string
}

func (e errorList) Error() string {
	return e.str
}

func mergeErrorMessages(errors []packages.Error) string {
	str := make([]string, len(errors))
	for i, err := range errors {
		str[i] = err.Error()
	}
	return strings.Join(str, "\n")
}

func environForCompiler(enableModule bool) []string {
	env := append(os.Environ(),
		"GOARCH="+build.Default.GOARCH,
		"GOOS="+build.Default.GOOS,
		"GOROOT="+build.Default.GOROOT)
	if enableModule {
		env = append(env, "GO111MODULE=on")
	} else {
		env = append(env, "GO111MODULE=off")
	}
	return env
}
