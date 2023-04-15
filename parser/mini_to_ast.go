package parser // Mini

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/keep-cp/compiler-project-c/ast"
)

// MiniToAst is the entry function for creating AST from ANTLR parse tree
func MiniToAst(ctx IProgramContext) ast.Root {
	var root ast.Root
	root.Types = typeDeclarationsToAst(ctx.Types().AllTypeDeclaration())
	root.Declarations = declarationsToAst(ctx.Declarations().AllDeclaration())
	root.Functions = functionsToAst(ctx.Functions().AllFunction())

	return root
}

func constructPosition(tok antlr.Token) *ast.Position {
	return &ast.Position{
		Line:   tok.GetLine(),
		Column: tok.GetColumn(),
	}
}

func typeDeclarationsToAst(typeDecls []ITypeDeclarationContext) []*ast.TypeDeclaration {
	ret := make([]*ast.TypeDeclaration, len(typeDecls))

	for i, v := range typeDecls {
		td := ast.TypeDeclaration{
			Id:       v.ID().GetText(),
			Fields:   gatherFieldDeclaration(v.NestedDecl().AllDecl()),
			Position: constructPosition(v.ID().GetSymbol()),
		}
		ret[i] = &td

	}

	return ret
}

func gatherFieldDeclaration(decls []IDeclContext) []*ast.Declaration {
	ret := make([]*ast.Declaration, len(decls))
	for i, v := range decls {
		decl := ast.Declaration{
			Name: v.ID().GetText(),
			Type: typeToAst(v.Type_()),
		}
		ret[i] = &decl
	}
	return ret
}

func typeToAst(t ITypeContext) ast.Type {
	switch v := t.(type) {
	case *IntTypeContext:
		return ast.IntType{}
	case *BoolTypeContext:
		return ast.BoolType{}
	case *StructTypeContext:
		return ast.StructType{Id: v.ID().GetText()}
	}
	return nil
}

func declarationsToAst(declarations []IDeclarationContext) []*ast.Declaration {
	var td []*ast.Declaration

	fmt.Println(declarations[0].Type_().GetText())
	return td
}

func functionsToAst(functions []IFunctionContext) []*ast.Function {
	var td []*ast.Function

	fmt.Println(functions[0].Parameters().AllDecl()[0].Type_().GetText())
	return td
}
