package main

import (
	"fmt"
	"strings"

	"example.com/tadpole/entity"
	"example.com/tadpole/repository"
	"example.com/tadpole/use_case"
)

type ConsolePresenter struct{}

func (p ConsolePresenter) PrintMatchedData(matchedData entity.MatchedData) {
	if matchedData.IsContentMatched {
		fmt.Println("找到了匹配内容的文件" + matchedData.Doc.Name)
		lineNums := matchedData.LineNums
		lines := strings.Split(matchedData.Doc.Content, "\n")
		for _, lineNum := range lineNums {
			content := lines[lineNum]
			fmt.Printf("%s:%d:%s\n", matchedData.Doc.Name, lineNum, content)
		}
	}
	if matchedData.IsNameMatched {
		fmt.Println("找到了匹配名字的文件" + matchedData.Doc.Name)
	}
}

func main() {
	useCase := use_case.FindByKeywordUseCase{
		Presenter:  ConsolePresenter{},
		Repository: repository.FSDocRepository{ValidSuffixes: []string{".org"}},
	}
	useCase.Run()
}
