// this file was generated by gomacro command: import _i "github.com/lifepod-solutions/gomacro/base/output"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package output

import (
	r "reflect"
	"github.com/lifepod-solutions/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/lifepod-solutions/gomacro/base/output"
func init() {
	imports.Packages["github.com/lifepod-solutions/gomacro/base/output"] = imports.Package{
	Binds: map[string]r.Value{
		"Debugf":	r.ValueOf(Debugf),
		"Error":	r.ValueOf(Error),
		"Errorf":	r.ValueOf(Errorf),
		"MakeRuntimeError":	r.ValueOf(MakeRuntimeError),
		"ShowPackageHeader":	r.ValueOf(ShowPackageHeader),
		"Warnf":	r.ValueOf(Warnf),
	}, Types: map[string]r.Type{
		"Output":	r.TypeOf((*Output)(nil)).Elem(),
		"RuntimeError":	r.TypeOf((*RuntimeError)(nil)).Elem(),
		"Stringer":	r.TypeOf((*Stringer)(nil)).Elem(),
	}, Wrappers: map[string][]string{
		"Output":	[]string{"Copy","ErrorAt","Errorf","Fprintf","IncLine","IncLineBytes","MakeRuntimeError","Position","Sprintf","ToString",},
	}, 
	}
}
