package data

// START OMIT

type Def struct {
	Path       string // unique identifier for this definition in its package
	Name       string
	Start, End int
	File       string
	Repo       string
	Lang       string
	Data       interface{} // extra language-specific info about this definition
}

type Ref struct {
	DefPath    string
	DefRepo    string
	File       string
	Repo       string
	Start, End int
	Lang       string
}

// END OMIT
