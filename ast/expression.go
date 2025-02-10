package ast

type Operator int

const (
	PLUS Operator = iota
	ASSIGN
	MINUS
	FSLASH
	ASTERIX
	GT
	LT
	GTE
	LTE
	EQUALS
	NEQUALS
	AND
	OR
	DOT
)

func StrToOp(op string) Operator {
	switch op {
	case "+":
		return PLUS
	case "-":
		return MINUS
	case "*":
		return ASTERIX
	case "/":
		return FSLASH
	case "=":
		return ASSIGN
	case ">":
		return GT
	case ">=":
		return GTE
	case "<":
		return LT
	case "<=":
		return LTE
	case "==":
		return EQUALS
	case "!=":
		return NEQUALS
	case "&&":
		return AND
	case "||":
		return OR
	case ".":
		return DOT
	}
	panic("Could not find operator")
}

func OpToStr(op Operator) string {
	switch op {
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case ASTERIX:
		return "*"
	case FSLASH:
		return "/"
	case ASSIGN:
		return "="
	case GT:
		return ">"
	case GTE:
		return ">="
	case LT:
		return "<"
	case LTE:
		return "<="
	case EQUALS:
		return "=="
	case NEQUALS:
		return "!="
	case AND:
		return "&&"
	case OR:
		return "||"
	case DOT:
		return "."
	}
	panic("Could not determine operator")
}

type Expression interface {
	String() string
}