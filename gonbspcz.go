package gonbspcz

import "regexp"

// structure of the expressions to find and replace
type setup struct {
	regex       string
	replacement string
}

// list of all replacements to do
var list = []setup{
	setup{regex: `(\d) `, replacement: "$1 "},
	setup{regex: `^(A|I|O|U|K|S|V|Z) `, replacement: "$1 "},
	setup{regex: ` (A|I|O|U|K|S|V|Z|a|i|o|u|k|s|v|z) `, replacement: " $1 "},
	setup{regex: `(\d\.) `, replacement: "$1 "},
	setup{regex: `([A-Z].) `, replacement: "$1 "},
}

// Replace will retrieve the original string and return the modified one
func Replace(s string) string {
	for {
		// replace all cases by its replacement string
		gen := s
		for i := 0; i < len(list); i++ {
			gen = regexp.MustCompile(list[i].regex).ReplaceAllString(gen, list[i].replacement)
		}

		// do the comparison
		if s == gen {
			return s
		}

		// assign new value
		s = gen
	}
}
