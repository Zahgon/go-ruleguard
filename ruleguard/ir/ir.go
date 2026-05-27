package ir

type File struct {
	PkgPath string

	RuleGroups []RuleGroup

	CustomDecls []string

	BundleImports []BundleImport
}

type BundleImport struct {
	Line int

	PkgPath string
	Prefix  string
}

type RuleGroup struct {
	Line        int
	Name        string
	MatcherName string

	DocTags    []string
	DocSummary string
	DocBefore  string
	DocAfter   string
	DocNote    string

	Imports []PackageImport

	Rules []Rule
}

type PackageImport struct {
	Path string
	Name string
}

type Rule struct {
	Line int

	SyntaxPatterns  []PatternString
	CommentPatterns []PatternString

	ReportTemplate  string
	SuggestTemplate string
	DoFuncName      string

	WhereExpr FilterExpr

	LocationVar string
}

type PatternString struct {
	Line  int
	Value string
}

// stringer -type=FilterOp -trimprefix=Filter

//go:generate go run ./gen_filter_op.go
type FilterOp int

func (op FilterOp) String() string { _ = "STUB: not implemented"; return "" }

type FilterExpr struct {
	Line int

	Op    FilterOp
	Src   string
	Value interface{}
	Args  []FilterExpr
}

func (e FilterExpr) IsValid() bool { _ = "STUB: not implemented"; return false }

func (e FilterExpr) IsBinaryExpr() bool { _ = "STUB: not implemented"; return false }
func (e FilterExpr) IsBasicLit() bool   { _ = "STUB: not implemented"; return false }
func (e FilterExpr) HasVar() bool       { _ = "STUB: not implemented"; return false }

func (e FilterExpr) String() string { _ = "STUB: not implemented"; return "" }

const (
	flagIsBinaryExpr uint64 = 1 << iota
	flagIsBasicLit
	flagHasVar
)
