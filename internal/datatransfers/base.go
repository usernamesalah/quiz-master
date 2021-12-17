package datatransfers

import (
	"github.com/dgrijalva/jwt-go"
)

type ErrorData struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Success bool        `json:"success"`
	Error   *ErrorData  `json:"errors"`
	Data    interface{} `json:"data"`
	Paging  *PageData   `json:"paging"`
}

type PageData struct {
	HasNext     bool  `json:"hasNext"`
	CurrentPage int   `json:"currentPage"`
	TotalPages  int   `json:"totalPages"`
	TotalData   int64 `json:"totalData"`
	Limit       int   `json:"limit"`
}

type ListQueryParams struct {
	Limit   int
	Offset  int
	Query   string
	Keyword string
}

type JWTClaims struct {
	UserData
	jwt.StandardClaims
}

type UserData struct {
	UID     string `json:"uid"`
	IsAdmin bool   `json:"isAdmin"`
}
