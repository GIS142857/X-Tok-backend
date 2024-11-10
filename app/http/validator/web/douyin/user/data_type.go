package user

type Phone struct {
	Phone *string `form:"phone" json:"phone" binding:"required,len=11"`
}

type Password struct {
	Password *string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

type Uid struct {
	Uid *float64 `form:"uid" json:"uid" binding:"required,gt=0"`
}

type PageNo struct {
	PageNo *float64 `form:"pageNo" json:"pageNo" binding:"required,min=0"`
}

type PageSize struct {
	PageSize *float64 `form:"pageSize" json:"pageSize" binding:"required,min=0"`
}
