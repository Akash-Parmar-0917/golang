package models

import (
	"math"

	"github.com/Akash-Parmar-0917/ecommerce/pkg/config"
)

type PaginationData struct {
	Limit int
	NextPage int
	PreviousPage int
	CurrentPage  int    `json:"page"`
	TotalPages int
	TwoAfter int
	TwoBelow int
	ThreeAfter int
	Offset int
	Sort  string `json:"sort"`
	WhereQuery string
	ReqUrl string
}


func GetPaginationData(page int,model interface{},sort string,whereQuery string,reqUrl string) PaginationData{
	var limit int=2
	var totalRows int64
	db:=config.GetDB()
	if(whereQuery!=""){
		db.Model(model).Where(whereQuery).Count(&totalRows)
	}else{
		db.Model(model).Count(&totalRows)
	}
	
	totalPages:=math.Ceil(float64(totalRows)/float64(limit))
	return PaginationData{
		Limit: limit,
		NextPage: page+1,
		PreviousPage: page-1,
		CurrentPage: page,
		TotalPages: int(totalPages),
		TwoAfter: page+2,
		TwoBelow: page-2,
		ThreeAfter: page+3,
		Offset: (page-1)*limit,
		WhereQuery: whereQuery,
		ReqUrl: reqUrl,
		Sort: sort,
	}
}