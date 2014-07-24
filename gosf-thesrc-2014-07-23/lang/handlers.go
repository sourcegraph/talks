package lang

import "path/filepath"

// START OMIT

func AnalyzeFile(file string, hs map[string]Analyzer) {
	hs[filepath.Ext(file)[1:]].Analyze(file)
}

// END OMIT

func doOtherStuff(file string, hs map[string]Analyzer) {} // dummy

func filesIn(dir string) []string { return nil } // dummy
