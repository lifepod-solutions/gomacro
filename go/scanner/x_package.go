// this file was generated by gomacro command: import _i "github.com/lifepod-solutions/gomacro/go/scanner"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package scanner

import (
	r "reflect"

	"github.com/lifepod-solutions/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/lifepod-solutions/gomacro/go/scanner"
func init() {
	imports.Packages["github.com/lifepod-solutions/gomacro/go/scanner"] = imports.Package{
		Binds: map[string]r.Value{
			"PrintError":   r.ValueOf(PrintError),
			"ScanComments": r.ValueOf(ScanComments),
		},
		Types: map[string]r.Type{
			"Error":        r.TypeOf((*Error)(nil)).Elem(),
			"ErrorHandler": r.TypeOf((*ErrorHandler)(nil)).Elem(),
			"ErrorList":    r.TypeOf((*ErrorList)(nil)).Elem(),
			"Mode":         r.TypeOf((*Mode)(nil)).Elem(),
			"Scanner":      r.TypeOf((*Scanner)(nil)).Elem(),
		},
		Proxies: map[string]r.Type{}}
}
