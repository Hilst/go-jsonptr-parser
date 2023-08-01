package syntax

import lx "github.com/Hilst/app-ui/lexer"

type SubstitutionType string

const (
	WhereSubs SubstitutionType = "where"
	IndexSubs SubstitutionType = "index"
)

type Analyser struct {
	tokens  []lx.Token
	subsMap map[SubstitutionType][][2]int
}

func NewSyntaxAnalyser(tokens []lx.Token) Analyser {
	a := Analyser{
		tokens,
		map[SubstitutionType][][2]int{
			WhereSubs: {},
			IndexSubs: {},
		},
	}
	a.mapSubstitutions()
	return a
}

func (a *Analyser) GetSubstitutions(subType SubstitutionType) [][2]int {
	return a.subsMap[subType]
}

func (a *Analyser) mapSubstitutions() {
	opened := false
	var startIndex int
	var subsType SubstitutionType
	for i := 1; i < len(a.tokens); i++ {
		if !opened && a.tokens[i].Type == lx.LSql && a.tokens[i-1].Type == lx.LSql {
			opened = true
			startIndex = i - 1
			continue
		}
		if opened && subsType == "" {
			switch a.tokens[i].Type {
			case lx.Where:
				subsType = WhereSubs
				break
			case lx.Index:
				subsType = IndexSubs
				break
			}
		}
		if opened && subsType != "" && a.tokens[i].Type == lx.RSql && a.tokens[i-1].Type == lx.RSql {
			opened = false
			a.subsMap[subsType] = append(a.subsMap[subsType], [2]int{startIndex, i})
			subsType = ""
		}
	}
}
