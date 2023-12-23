// Does not check for correct syntax or arguments
//
// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpcode

// Import Go standard library packages and tserr
import (
	"fmt"       // fmt
	"go/format" // format

	"github.com/thorstenrie/tserr" // tserr
)

// Code contains the source code as string. The source code is amended by
// its methods. The source code can be retrieved with String and formatted
// with Format.
type Code struct {
	c string // the source code
}

// NewCode returns a pointer to a new Code instance.
// The source code in the new Code instance is empty.
func NewCode() *Code {
	// Return a new Code instance
	return &Code{}
}

// String returns the source code in code as string.
// It returns an empty string if code is nil.
func (code *Code) String() string {
	// Return an empty string if code is nil
	if code == nil {
		// Return an empty string
		return ""
	}
	// Return the source code as string
	return code.c
}

// LineComment adds a line comment and a new line to code: // c\n. The comment is provided by argument c.
func (code *Code) LineComment(c string) *Code {
	// Return nil if code is nil
	if code == nil {
		return nil
	}
	// Add a line comment and a new line to code
	code.c += fmt.Sprintf("// %v\n", c)
	// Return code
	return code
}

// FuncEnd adds a block end and two new lines to code: }\n\n.
func (code *Code) FuncEnd() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a block end and two new lines to code
	code.c += "}\n\n"
	// Return code
	return code
}

// BlockEnd adds a block ending to code: }\n.
func (code *Code) BlockEnd() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a block ending to code
	code.c += "}\n"
	// Return code
	return code
}

// Call adds a function call to code: n(. The function name is
// provided by n.
func (code *Code) Call(n string) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a function call to code
	code.c += fmt.Sprintf("%v(", n)
	// Return code
	return code
}

// ParamEndln adds a parameters ending and a new line to code: )\n.
func (code *Code) ParamEndln() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a parameters ending and a new line to code
	code.c += ")\n"
	// Return code
	return code
}

// ParamEnd adds a parameters ending to code: ).
func (code *Code) ParamEnd() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add parameters ending to code
	code.c += ")"
	// Return code
	return code
}

type Func1Args struct {
	Name, Var, Type, Return string
}

func (code *Code) Func1(a *Func1Args) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("func %v(%v %v) %v {\n", a.Name, a.Var, a.Type, a.Return)
	return code
}

// TypeStruct adds a type declaration for a struct type to code: type n struct {\n.
// The name of the type is provided with n.
func (code *Code) TypeStruct(n string) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a type declaration for a struct type to code
	code.c += fmt.Sprintf("type %v struct {\n", n)
	// Return code
	return code
}

// VarSpecArgs contains the identifier Ident and type Type to generate
// a variable specification with VarSpec
type VarSpecArgs struct {
	Ident, Type string
}

// VarSpec adds a variable specification to code: Ident Type\n. The identifier and type
// is provided by a.
func (code *Code) VarSpec(a *VarSpecArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add a variable specification to code
	code.c += fmt.Sprintf("%v %v\n", a.Ident, a.Type)
	// Return code
	return code
}

// List adds an identifier list to code: , .
func (code *Code) List() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add an identifier list to code
	code.c += ", "
	// Return code
	return code
}

// Listln adds an identifier list and a new line to code: ,\n.
func (code *Code) Listln() *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add an identifier list and a new line to code
	code.c += ",\n"
	// Return code
	return code
}

type SelArgs struct {
	Val, Sel string
}

// SelField adds a field selector to code: val.sel. The value val and selector sel are
// provided by a. It returns nil if a is nil.
func (code *Code) SelField(a *SelArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add a field selector to code
	code.c += fmt.Sprintf("%v.%v", a.Val, a.Sel)
	// Return code
	return code
}

// SelMethod adds a method selector to code: val.sel(. The value val and selector sel are
// provided by a. It returns nil if a is nil.
func (code *Code) SelMethod(a *SelArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add a method selector to code
	code.c += fmt.Sprintf("%v.%v(", a.Val, a.Sel)
	// Return code
	return code
}

// Expression
type IfArgs struct {
	ExprLeft, ExprRight, Operator string
}

// If statement
func (code *Code) If(a *IfArgs) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("if %v %v %v {\n", a.ExprLeft, a.Operator, a.ExprRight)
	return code
}

type IfErrArgs struct {
	Method, Operator string
}

// If statement for error handling using a simple statement
func (code *Code) IfErr(a *IfErrArgs) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("if err := %v; err %v nil {\n", a.Method, a.Operator)
	return code
}

// Return
func (code *Code) Return() *Code {
	if code == nil {
		return nil
	}
	code.c += "return "
	return code
}

// Address operator
func (code *Code) Addr() *Code {
	if code == nil {
		return nil
	}
	code.c += "&"
	return code
}

// Ident adds an identifier to code: n. The identifier is provided by argument n.
func (code *Code) Ident(n string) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Add identifier n to code
	code.c += n
	// Return code
	return code
}

type AssignmentArgs struct {
	ExprLeft, ExprRight string
}

// Assignment
func (code *Code) Assignment(a *AssignmentArgs) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("%v = %v", a.ExprLeft, a.ExprRight)
	return code
}

// Composite Literal
func (code *Code) CompositeLit(LiteralType string) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("%v{", LiteralType)
	return code
}

// ShortVarDeclArgs contains the identifier Ident and expression Expr to generate
// a short variable declaration with ShortVarDecl
type ShortVarDeclArgs struct {
	Ident, Expr string
}

// ShortVarDecl generates a short variable declaration: Ident := Expr\n. The identifier
// and expression is provided by a. It returns nil if a is nil.
func (code *Code) ShortVarDecl(a *ShortVarDeclArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add a short variable declaration to code
	code.c += fmt.Sprintf("%v := %v\n", a.Ident, a.Expr)
	// Return code
	return code
}

// KeyedElementArgs contains the key Key and element Elem to generate
// a keyed element of a composite literal with KeyedElement
type KeyedElementArgs struct {
	Key, Elem string
}

// KeyedElement generates a keyed element of a composite literal: Key: Element,\n. The key and element
// is provided by a. It returns nil if a is nil.
func (code *Code) KeyedElement(a *KeyedElementArgs) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case a is nil
	if a == nil {
		return nil
	}
	// Add keyed element of a composite literal to code
	code.c += fmt.Sprintf("%v: %v,\n", a.Key, a.Elem)
	// Return code
	return code
}

// Testvars contains the configuration for the generation of test variables with Testvariables
// A test variable is generated if the corresponding type is not equal to zero.
type Testvars struct {
	String, Error, Int, Float int
}

// Testvariables generates test variables for unit tests. The test variables are generated based
// on t. A test variable is generated if the corresponding type in t is not equal to zero.
// It returns nil if t is nil.
func (code *Code) Testvariables(t *Testvars) *Code {
	// Return nil in case code is nil
	if code == nil {
		return nil
	}
	// Return nil in case t is nil
	if t == nil {
		return nil
	}
	// Initialize text
	text := ""
	// Add a string test variable to text if String is not equal to zero
	if t.String != 0 {
		text += "strFoo string = \"foobar\" // test variable type string\n"
	}
	// Add an error test variable to text if Error is not equal to zero
	if t.Error != 0 {
		text += "errFoo error = fmt.Errorf(strFoo) // test variable type error\n"
	}
	// Add an integer test variable to text if Int is not equal to zero
	if t.Int != 0 {
		text += "intFoo int64 = 1234 // test variable type int64\n"
	}
	// Add a float test variable to text if Float is not equal to zero
	if t.Float != 0 {
		text += "floatFoo float64 = 1234 // test variable type float64\n"
	}
	// If test variables exist, add them to the variable declaration
	if text != "" {
		code.c += "var (\n"
		code.c += text
		code.c += ")\n\n"
	}
	// Return generated code
	return code
}

// Format formats the source code in code in canonical gfmt style.
// It uses Source from the go/format package. Format returns an error
// if code is nil or if go/format returns an error.
func (code *Code) Format() error {
	// Return an error in cae code is nil
	if code == nil {
		return tserr.NilPtr()
	}
	// Convert source code into a slice of bytes
	b := []byte(code.c)
	// Format the source code using Source from the go/format package
	o, e := format.Source(b)
	// Return an error in case Source fails
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "format source", Fn: "code", Err: e})
	}
	// Convert the formatted source code to string and store it in code
	code.c = string(o)
	// Return nil
	return nil
}
