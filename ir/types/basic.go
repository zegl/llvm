// === [ Basic types ] =========================================================

package types

import "fmt"

// --- [ void ] ----------------------------------------------------------------

// VoidType represents a void type.
//
// References:
//    http://llvm.org/docs/LangRef.html#void-type
type VoidType struct {
}

// String returns the LLVM syntax representation of the type.
func (t *VoidType) String() string {
	return "void"
}

// Equal reports whether t and u are of equal type.
func (t *VoidType) Equal(u Type) bool {
	_, ok := u.(*VoidType)
	return ok
}

// --- [ label ] ---------------------------------------------------------------

// LabelType represents a label type, which is used for basic block values.
//
// References:
//    http://llvm.org/docs/LangRef.html#label-type
type LabelType struct {
}

// String returns the LLVM syntax representation of the type.
func (t *LabelType) String() string {
	return "label"
}

// Equal reports whether t and u are of equal type.
func (t *LabelType) Equal(u Type) bool {
	_, ok := u.(*LabelType)
	return ok
}

// --- [ integer ] -------------------------------------------------------------

// IntType represents an integer type.
//
// References:
//    http://llvm.org/docs/LangRef.html#integer-type
type IntType struct {
	// Bit size.
	size int
}

// NewInt returns a new integer type based on the given bit size.
func NewInt(size int) *IntType {
	return &IntType{size: size}
}

// String returns the LLVM syntax representation of the type.
func (t *IntType) String() string {
	return fmt.Sprintf("i%d", t.size)
}

// Equal reports whether t and u are of equal type.
func (t *IntType) Equal(u Type) bool {
	if u, ok := u.(*IntType); ok {
		return t.size == u.size
	}
	return false
}

// Size returns the bit size of the integer type.
func (t *IntType) Size() int {
	return t.size
}

// --- [ floating-point ] ------------------------------------------------------

// FloatType represents a floating-point type.
//
// References:
//    http://llvm.org/docs/LangRef.html#floating-point-types
type FloatType struct {
	// Floating-point kind.
	kind FloatKind
}

// String returns the LLVM syntax representation of the type.
func (t *FloatType) String() string {
	return t.kind.String()
}

// Equal reports whether t and u are of equal type.
func (t *FloatType) Equal(u Type) bool {
	if u, ok := u.(*FloatType); ok {
		return t.kind == u.kind
	}
	return false
}

// Kind returns the kind of the floating-point type.
func (t *FloatType) Kind() FloatKind {
	return t.kind
}

// FloatKind represents the set of floating-point kinds.
type FloatKind int

// Floating point kinds.
const (
	FloatKindIEEE_16           FloatKind = iota // half:      16-bit floating point type
	FloatKindIEEE_32                            // float:     32-bit floating point type
	FloatKindIEEE_64                            // double:    64-bit floating point type
	FloatKindIEEE_128                           // fp128:     128-bit floating point type (112-bit mantissa)
	FloatKindDoubleExtended_80                  // x86_fp80:  80-bit floating point type (x87)
	FloatKindDoubleDouble_128                   // ppc_fp128: 128-bit floating point type (two 64-bits, PowerPC)
)

// String returns the LLVM syntax representation of the floating-point kind.
func (kind FloatKind) String() string {
	switch kind {
	case FloatKindIEEE_16:
		return "half"
	case FloatKindIEEE_32:
		return "float"
	case FloatKindIEEE_64:
		return "double"
	case FloatKindIEEE_128:
		return "fp128"
	case FloatKindDoubleExtended_80:
		return "x86_fp80"
	case FloatKindDoubleDouble_128:
		return "ppc_fp128"
	}
	return fmt.Sprintf("<unknown floating-point kind %d>", int(kind))
}