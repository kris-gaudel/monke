package parser

import (
	"monke/ast"
	"monke/lexer"
	"monke/token"
)

type Parser struct {
	l *lexer.lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read 2 tokens
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.nextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
