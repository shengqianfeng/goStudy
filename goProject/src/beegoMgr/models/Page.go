package models

import "math"

type Page struct {
	PageIndex int16 //当前页码
	PageCount int16 //总页数
	TotalCount int64 //总条数
	Pagesize  int8 //每页条数
}


func GetPage(pageSize int8,pageIndex int16,totalCount int64) *Page{
	page:=Page{
		PageIndex : pageIndex,
		Pagesize :pageSize,
		TotalCount:totalCount,
		PageCount : int16(math.Ceil(float64(totalCount) / float64(pageSize))) ,  //总页数
	}
	//当前页码小于0或者当前页大于最大页数，自动pageIndex页码为第一页
	if pageIndex <=0 || pageIndex > page.PageCount {
		page.PageIndex=1
	}
	return &page
}


/*
1.在视图中定义视图函数函数名，语法： 参数| funcName

2.一般在main.go里面实现试图函数

3.在main函数里面把实现的函数和试图函关联起来   beego.AddFuncMap()
 */
func ShowPrePage(pageindex int16)(preIndex int16){
	preIndex = pageindex - 1
	return
}

func ShowNextPage(pageindex int16)(nextIndex int16){
	nextIndex = pageindex + 1
	return
}