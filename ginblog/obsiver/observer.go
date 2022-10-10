package obsiver

type Observer interface {
	Update(string)
	GetID() string
}
