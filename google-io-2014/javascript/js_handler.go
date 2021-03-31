package javascript

import (
	"log"
	"regexp"

	"github.com/sourcegraph/talks/google-io-2014/lang"
)

// START OMIT

type JSAnalyzer struct{}

var jsdef = regexp.MustCompile(`(var|function) (\w+)`)

func (_ JSAnalyzer) Analyze(file string) ([]*lang.Def, []*lang.Ref, error) { // HL
	src, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var defs []*lang.Def
	for _, m := range jsdef.FindAllStringSubmatch(string(src), -1) {
		defs = append(defs, &lang.Def{Name: m[2], Type: m[1]})
	}
	return defs, nil, nil
}

// END OMIT

// dummy
func (_ JSAnalyzer) Scan(dir string) ([]string, error) { return nil, nil }

// dummy
func (_ JSAnalyzer) ListDependencies(pkg string) ([]*lang.Dep, error) { return nil, nil }

var _ lang.Analyzer = &JSAnalyzer{}
