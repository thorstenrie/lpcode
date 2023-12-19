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

// Line comment
func (code *Code) LineComment(c string) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("// %v\n", c)
	return code
}

// Block End with new line
func (code *Code) FuncEnd() *Code {
	if code == nil {
		return nil
	}
	code.c += "}\n\n"
	return code
}

// Block End
func (code *Code) BlockEnd() *Code {
	if code == nil {
		return nil
	}
	code.c += "}\n"
	return code
}

// Function Call
func (code *Code) Call(n string) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("%v(", n)
	return code
}

// Parameters End + new line
func (code *Code) ParamEndln() *Code {
	if code == nil {
		return nil
	}
	code.c += ")\n"
	return code
}

// Parameters End
func (code *Code) ParamEnd() *Code {
	if code == nil {
		return nil
	}
	code.c += ")"
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

// Type declaration for struct type
func (code *Code) TypeStruct(n string) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("type %v struct {\n", n)
	return code
}

type TypeArgs struct {
	Name, Type string
}

// Type
func (code *Code) Type(a *TypeArgs) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("%v %v\n", a.Name, a.Type)
	return code
}

// IdentifierLIst
func (code *Code) List() *Code {
	if code == nil {
		return nil
	}
	code.c += ", "
	return code
}

// IdentifierList with new line
func (code *Code) Listln() *Code {
	if code == nil {
		return nil
	}
	code.c += ",\n"
	return code
}

type SelArgs struct {
	Val, Sel string
}

// Field Selector
func (code *Code) SelField(a *SelArgs) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("%v.%v", a.Val, a.Sel)
	return code
}

// Method Selector
func (code *Code) SelMethod(a *SelArgs) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("%v.%v(", a.Val, a.Sel)
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

// Identifier
func (code *Code) Ident(n string) *Code {
	if code == nil {
		return nil
	}
	code.c += n
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

type ShortVarDeclArgs struct {
	Ident, Expr string
}

// Short variable declaration
func (code *Code) ShortVarDecl(a *ShortVarDeclArgs) *Code {
	if code == nil {
		return nil
	}
	code.c += fmt.Sprintf("%v := %v\n", a.Ident, a.Expr)
	return code
}

// KeyedElementArgs contains the key and element to generate
// a keyed element of a composite literal with KeyedElement
type KeyedElementArgs struct {
	Key, Element string
}

// KeyedElement generates a keyed element of a composite literal: Key: Element,\n. The key and element
// is provided by a.
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
	code.c += fmt.Sprintf("%v: %v,\n", a.Key, a.Element)
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
