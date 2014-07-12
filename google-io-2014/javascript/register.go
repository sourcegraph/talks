package javascript

import "github.com/sourcegraph/talks/google-io-2014/lang" // OMIT
//OMIT
func init() {
	lang.Register("js", &JSAnalyzer{}) // HL
}
