package main

import (
	"flag"
	"fmt"
	"strings"

	"example.com/tadpole/entity"
	"example.com/tadpole/repository"
	"example.com/tadpole/use_case"
)

type ConsolePresenter struct{}
type CLIParam struct {
	keyword string
}

func (p CLIParam) GetKeyword() string {
	return p.keyword
}

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
	keywordPtr := flag.String("keyword", "", "关键字")
	rootDirPtr := flag.String("root-dir", "/Users/liutos/Dropbox", "搜索目录")
	suffixPtr := flag.String("suffix", ".org", "后缀")
	flag.Parse()
	useCase := use_case.FindByKeywordUseCase{
		Param: CLIParam{
			keyword: *keywordPtr,
		},
		Presenter: ConsolePresenter{},
		Repository: repository.FSDocRepository{
			RootDir:       *rootDirPtr,
			ValidSuffixes: []string{*suffixPtr},
		},
	}
	useCase.Run()
}
