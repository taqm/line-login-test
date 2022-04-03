package domains

type User struct {
	ID          string
	DisplayName string
	PictureURL  string
}

func (u User) IsZeroValue() bool {
	return u == User{}
}
