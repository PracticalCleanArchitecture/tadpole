package repository

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"example.com/tadpole/entity"
)

type FSDocRepository struct {
	RootDir       string
	ValidSuffixes []string
}

func collectFileNames(dir string) ([]string, error) {
	fmt.Println("处理目录" + dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	fileNames := []string{}
	for _, file := range files {
		if file.IsDir() {
			subDir := dir + string(os.PathSeparator) + file.Name()
			subFileNames, err := collectFileNames(subDir)
			if err != nil {
				return nil, err
			}
			fileNames = append(fileNames, subFileNames...)
		} else {
			fileNames = append(fileNames, dir+string(os.PathSeparator)+file.Name())
		}
	}
	return fileNames, nil
}

func (r FSDocRepository) Find(keyword string) ([]entity.MatchedData, error) {
	dir := r.RootDir
	fileNames, err := collectFileNames(dir)
	if err != nil {
		return nil, err
	}
	fmt.Println("遍历完毕")

	matchedDataList := []entity.MatchedData{}
	validSuffixes := r.ValidSuffixes
	for _, fileName := range fileNames {
		hasValidSuffix := false
		for _, validSuffix := range validSuffixes {
			if strings.HasSuffix(fileName, validSuffix) {
				hasValidSuffix = true
				break
			}
		}
		if !hasValidSuffix {
			fmt.Println("忽略文件" + fileName)
			continue
		}
		fmt.Println("处理文件" + fileName)
		// 文本文件都很小，直接读取所有内容不会有问题。
		fileContent, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		text := string(fileContent)
		isContentMatched := strings.Contains(text, keyword)
		isNameMatched := strings.Contains(fileName, keyword)
		if isNameMatched || isContentMatched {
			lineNums := []int{}
			if isContentMatched {
				lines := strings.Split(text, "\n")
				for lineNum, line := range lines {
					if strings.Contains(line, keyword) {
						lineNums = append(lineNums, lineNum)
					}
				}
			}

			doc := entity.Doc{
				Content: text,
				Name:    fileName,
			}
			matchedData := entity.MatchedData{
				Doc:              doc,
				IsContentMatched: isContentMatched,
				IsNameMatched:    isNameMatched,
				LineNums:         lineNums,
			}
			matchedDataList = append(matchedDataList, matchedData)
		}
	}

	return matchedDataList, nil
}
