package lang

var analyzers = make(map[string]Analyzer) // HL

func Register(language string, a Analyzer) { // HL
	// maybe check for dupes or non-nil
	analyzers[language] = a
}
