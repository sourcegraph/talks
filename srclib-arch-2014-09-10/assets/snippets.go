import (
	"github.com/jmoiron/sqlx/types"
	"sourcegraph.com/sourcegraph/srclib/repo"
)

// START UNIT_STRUCT OMIT
type SourceUnit struct {
	// Name is an opaque identifier for this source unit that MUST be unique
	// among all other source units of the same type in the same repository.
	Name string

	// Type is the type of source unit this represents, such as "GoPackage".
	Type string

	Repo repo.URI

	Globs        []string
	Files        []string
	Dir          string
	Dependencies []interface{}

	Info   *Info
	Data   interface{}
	Config map[string]interface{}
}

// END UNIT_STRUCT OMIT

// START DEF_STRUCT OMIT
type Def struct {
	// DefKey
	Repo     repo.URI
	CommitID string
	UnitType string
	Unit     string
	Path     DefPath

	Kind     DefKind
	Name     string
	Callable bool
	Exported bool

	File     string
	DefStart int
	DefEnd   int

	Data types.JsonText
}

// END DEF_STRUCT OMIT

// START REF_STRUCT OMIT
type Ref struct {
	// The definition to which this reference points
	DefRepo     repo.URI
	DefUnitType string
	DefUnit     string
	DefPath     DefPath

	Repo     repo.URI
	CommitID string
	UnitType string
	Unit     string

	File  string
	Start int
	End   int
}

// END REF_STRUCT OMIT
