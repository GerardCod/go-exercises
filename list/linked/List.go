package linked

//List es la interfaz que define los metodos para las implementaciones, ya sea una LinkedList o una ArrayList.
type List interface {
  First() (int, error)
  Last() (int, error)
  Find(number int) (int, error)
  IndexOf(number int) int
  Append(number int)
  Shift(number int)
  Unshift() (int, error)
  Pop() (int, error)
  Size() int
  IsEmpty() bool
  Show() error
}
