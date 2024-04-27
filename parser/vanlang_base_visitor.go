// Generated from ./grammar/VanLang.g4 by ANTLR 4.7.

package parser // VanLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseVanLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVanLangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitDatatype(ctx *DatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitLogicStatement(ctx *LogicStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitConditionStatement(ctx *ConditionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVanLangVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}
