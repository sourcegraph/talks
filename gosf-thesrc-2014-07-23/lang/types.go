package lang

type Def struct {
	Name string
	Type string
}

type Ref struct{}

type Dep struct{}

// START 2 OMIT

type Analyzer interface {
	// Analyze a package to find all definitions (funcs, types, vars, etc.) and references
	Analyze(pkg string) ([]*Def, []*Ref, error) // HL
}

type DependencyLister interface {
	// ListDependencies of pkg, which is either a file (e.g., package.json, setup.py)
	// or a dir (e.g., a Go package directory).
	ListDependencies(pkg string) ([]*Dep, error)
}

// END 2 OMIT
