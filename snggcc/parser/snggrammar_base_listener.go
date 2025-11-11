// Code generated from snggrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // snggrammar
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BasesnggrammarListener is a complete listener for a parse tree produced by snggrammarParser.
type BasesnggrammarListener struct{}

var _ snggrammarListener = &BasesnggrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesnggrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesnggrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesnggrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesnggrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasesnggrammarListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasesnggrammarListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasesnggrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasesnggrammarListener) ExitExpr(ctx *ExprContext) {}
