package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// A LocalVarDef represents a local variable definition.
//
// Examples:
//    %foo = add i32 13, 29
type LocalVarDef struct {
	// Name of the local variable.
	name string
	// Initial value.
	valInst ValueInst
}

// NewLocalVarDef returns a new local variable definition based on the given
// name and initial value instruction.
func NewLocalVarDef(name string, valInst ValueInst) (*LocalVarDef, error) {
	// TODO: Verify that name is not a local ID. Unnamed local variable
	// definitions should be assigned a local ID implicitly by the internal
	// localID counter of the given function rather than explicitly assigned.
	return &LocalVarDef{name: name, valInst: valInst}, nil
}

// Name returns the name of the defined local variable.
func (def *LocalVarDef) Name() string {
	return def.name
}

// TODO: Add note to SetName not set local IDs explicitly, as these are assigned
// implicitly by the internal localID counter.

// SetName sets the name of the defined local variable.
func (def *LocalVarDef) SetName(name string) {
	def.name = name
}

// ValInst returns the instruction defining the initial value of the local
// variable definition.
func (def *LocalVarDef) ValInst() ValueInst {
	return def.valInst
}

// Type returns the type of the value.
func (def *LocalVarDef) Type() types.Type {
	return def.valInst.RetType()
}

// String returns the string representation of the local variable definition;
// e.g.
//
//    %foo = add i32 13, 42
func (def *LocalVarDef) String() string {
	if types.IsVoid(def.Type()) || len(def.Name()) == 0 {
		return def.valInst.String()
	}
	return fmt.Sprintf("%s = %s", def.ValueString(), def.valInst)
}

// ValueString returns a string representation of the value.
func (def *LocalVarDef) ValueString() string {
	return enc.Local(def.Name())
}

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (def *LocalVarDef) isInst() {}
