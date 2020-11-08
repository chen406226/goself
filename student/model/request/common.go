package request

// Paging common input paramenter structure
type PageInfo struct {
	Page		int	`json:"page" form:"page"`
	PageSize	int	`json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id	float64	`json:"id" form:"id"`
}

type IdsReq struct {
	Ids	[]int	`json:"ids" form:"ids"`
}