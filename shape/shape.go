package shape

type Color int

const (
    Red Color = iota
    Blue
    Green
)

type Shape struct {
    U Color
    D Color
    L Color
    R Color
}

