// Generated from ./grammar/VanLang.g4 by ANTLR 4.7.

package parser // VanLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVanLangListener is a complete listener for a parse tree produced by VanLangParser.
type BaseVanLangListener struct{}

var _ VanLangListener = &BaseVanLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVanLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVanLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVanLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVanLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseVanLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseVanLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseVanLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseVanLangListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseVanLangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseVanLangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseVanLangListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseVanLangListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseVanLangListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseVanLangListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BaseVanLangListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BaseVanLangListener) ExitDatatype(ctx *DatatypeContext) {}

// EnterLogicStatement is called when production logicStatement is entered.
func (s *BaseVanLangListener) EnterLogicStatement(ctx *LogicStatementContext) {}

// ExitLogicStatement is called when production logicStatement is exited.
func (s *BaseVanLangListener) ExitLogicStatement(ctx *LogicStatementContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseVanLangListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseVanLangListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterConditionStatement is called when production conditionStatement is entered.
func (s *BaseVanLangListener) EnterConditionStatement(ctx *ConditionStatementContext) {}

// ExitConditionStatement is called when production conditionStatement is exited.
func (s *BaseVanLangListener) ExitConditionStatement(ctx *ConditionStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseVanLangListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseVanLangListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVanLangListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVanLangListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseVanLangListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseVanLangListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseVanLangListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseVanLangListener) ExitFactor(ctx *FactorContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseVanLangListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseVanLangListener) ExitAtom(ctx *AtomContext) {}
