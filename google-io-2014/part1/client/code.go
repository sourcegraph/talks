package client

// START IFACE OMIT
type CodeService interface {
	Get(def DefSpec, opt *CodeGetOptions) (*Def, Response, error)
	List(opt *CodeListDefOptions) ([]*Def, Response, error)
	ListRefs(def DefSpec, opt *CodeListRefsOptions) ([]*Ref, Response, error)
	// ...
}

// END IFACE OMIT

type Def struct{}
type Ref struct{}
type DefSpec struct{}
type CodeGetOptions struct{}
type CodeListDefOptions struct{}
type CodeListRefsOptions struct{}
type CodeDependency struct{}

type codeService struct{}
