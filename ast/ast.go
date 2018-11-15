package ast

import "github.com/Interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token //token.IDENT
	Value string
}

type LetStatement struct {
	Token token.Token //token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (i *Identifier) expressionNode() {

}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
