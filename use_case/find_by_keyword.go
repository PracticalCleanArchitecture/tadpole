package use_case

import (
	"log"

	"example.com/tadpole/entity"
)

type IParam interface {
	GetKeyword() string
}

type IPresenter interface {
	PrintMatchedData(entity.MatchedData)
}

type FindByKeywordUseCase struct {
	Param      IParam
	Presenter  IPresenter
	Repository entity.IDocRepository
}

func (u FindByKeywordUseCase) Run() {
	param := u.Param
	r := u.Repository
	keyword := param.GetKeyword()
	matchedDataList, err := r.Find(keyword)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, matchedData := range matchedDataList {
		u.Presenter.PrintMatchedData(matchedData)
	}
}
