// Generated from ./grammar/VanLang.g4 by ANTLR 4.7.

package parser // VanLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by VanLangParser.
type VanLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VanLangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by VanLangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by VanLangParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by VanLangParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by VanLangParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by VanLangParser#datatype.
	VisitDatatype(ctx *DatatypeContext) interface{}

	// Visit a parse tree produced by VanLangParser#logicStatement.
	VisitLogicStatement(ctx *LogicStatementContext) interface{}

	// Visit a parse tree produced by VanLangParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by VanLangParser#conditionStatement.
	VisitConditionStatement(ctx *ConditionStatementContext) interface{}

	// Visit a parse tree produced by VanLangParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by VanLangParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by VanLangParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by VanLangParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by VanLangParser#atom.
	VisitAtom(ctx *AtomContext) interface{}
}
