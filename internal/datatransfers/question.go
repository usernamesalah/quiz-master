package datatransfers

type Answer struct {
	Value      string `json:"value" valid:"required"`
	QuestionID int    `json:"questionID"`
}
