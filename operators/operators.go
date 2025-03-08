package operators

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
	NOT
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
	case "!":
		return NOT
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
	case NOT:
		return "!"
	}
	panic("Could not determine operator")
}

func IsArithmeticOp(op Operator) bool {
	if op == PLUS || op == MINUS || op == FSLASH || op == ASTERIX {
		return true
	}
	return false
}

func IsLogicalOp(op Operator) bool {
	if op == OR || op == AND || op == NOT {
		return true
	}
	return false
}

func IsComparisonOp(op Operator) bool {
	if op == GT || op == LT || op == GTE || op == LTE {
		return true
	}
	return false
}

func IsEqualityOp(op Operator) bool {
	if op == EQUALS || op == NEQUALS {
		return true
	}
	return false
}

func OpToLLVM(op Operator) string {
	switch op {
	case PLUS:
		return "add"
	case MINUS:
		return "sub"
	case ASTERIX:
		return "mul"
	case FSLASH:
		return "sdiv"
	case GT:
		return "sgt"
	case GTE:
		return "sge"
	case LT:
		return "slt"
	case LTE:
		return "sle"
	case EQUALS:
		return "eq"
	case NEQUALS:
		return "ne"
	case AND:
		return "and"
	case OR:
		return "or"
	case NOT:
		return "!"
	}
	panic("Could not determine operator")
}