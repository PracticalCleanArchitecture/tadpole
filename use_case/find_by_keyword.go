package use_case

import (
	"log"

	"example.com/tadpole/entity"
)

type IPresenter interface {
	PrintMatchedData(entity.MatchedData)
}

type FindByKeywordUseCase struct {
	Presenter  IPresenter
	Repository entity.IDocRepository
}

func (u FindByKeywordUseCase) Run() {
	r := u.Repository
	matchedDataList, err := r.Find("Acc")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, matchedData := range matchedDataList {
		u.Presenter.PrintMatchedData(matchedData)
	}
}
