package response

type List[T any] struct {
	Total int64 `json:"total"`
	List  []T   `json:"list"`
}
