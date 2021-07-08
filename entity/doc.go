package entity

type MatchedLine struct {
	Content string
}

type Doc struct {
	Content string
	Name    string
}

type MatchedData struct {
	Doc              Doc
	IsContentMatched bool
	IsNameMatched    bool
	Keyword          string
	LineNums         []int
}

type IDocRepository interface {
	Find(keyword string) ([]MatchedData, error)
}
