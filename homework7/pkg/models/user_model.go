package models

type User struct {
        Id		int	`json:"id"`
	NickName	string	`json:"nickname"`
	UserName	string	`json:"username"`
        Password     	string 	`json:"password"`
        Created_at 	string 	`json:"created_at"`
	Updated_at	string	`json:"updated_at"`
}
type Messages struct {
	Id		int		`json:"id"`
	UserId		int		`json:"userid"`
	Context		string		`json:"context"`
        Created_at 	string 		`json:"created_at"`
	Updated_at	string		`json:"updated_at"`
	IsDeleted	int		`json:"isdeleted"`
	ParentId	int		`json:"parentid"`
}

