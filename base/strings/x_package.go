// this file was generated by gomacro command: import _i "github.com/lifepod-solutions/gomacro/base/strings"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package strings

import (
	r "reflect"
	"github.com/lifepod-solutions/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/lifepod-solutions/gomacro/base/strings"
func init() {
	imports.Packages["github.com/lifepod-solutions/gomacro/base/strings"] = imports.Package{
	Binds: map[string]r.Value{
		"FindFirstToken":	r.ValueOf(FindFirstToken),
		"MaybeUnescapeString":	r.ValueOf(MaybeUnescapeString),
		"Split2":	r.ValueOf(Split2),
		"UnescapeChar":	r.ValueOf(UnescapeChar),
		"UnescapeString":	r.ValueOf(UnescapeString),
	}, 
	}
}
