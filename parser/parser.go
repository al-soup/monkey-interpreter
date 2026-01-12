package parser

import (
	"github.com/al-soup/monkey-interpreter/ast"
	"github.com/al-soup/monkey-interpreter/lexer"
	"github.com/al-soup/monkey-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// set cur- and peerToken
	p.l.NextToken()
	p.l.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
