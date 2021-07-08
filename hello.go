package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mgutz/ansi"

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
		fmt.Println(ansi.Color(matchedData.Doc.Name, "magenta"))
		lineNums := matchedData.LineNums
		lines := strings.Split(matchedData.Doc.Content, "\n")
		for _, lineNum := range lineNums {
			content := lines[lineNum]
			coloredLineNum := ansi.Color(fmt.Sprintf("%d", lineNum), "green")
			coloredContent := strings.ReplaceAll(content, matchedData.Keyword, ansi.Color(matchedData.Keyword, "red"))
			fmt.Printf("%s:%s\n", coloredLineNum, coloredContent)
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
