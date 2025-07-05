package patterns

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {

}

type Square struct{}

func (s *Square) Draw() {

}

func CreateShape(stype string) Shape {
	switch stype {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		return nil
	}
}

func main() {
	circle := CreateShape("circle")
	circle.Draw()

	square := CreateShape("square")
	square.Draw()
}
