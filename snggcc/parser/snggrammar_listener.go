// Code generated from snggrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // snggrammar
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// snggrammarListener is a complete listener for a parse tree produced by snggrammarParser.
type snggrammarListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
