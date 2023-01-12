package main

type IFlyBehaviour interface {
	fly()
}

type IQuackBehaviour interface {
	quack()
}

// Two flying strategies
type StraightFlyStrategy struct {
}

func (d StraightFlyStrategy) fly() {
	print("straight fly")
}

type ZigZagFlyStrategy struct {
}

func (d ZigZagFlyStrategy) fly() {
	print("zigag fly")
}

// Two quacking strategies
type QuackStrategy struct {
}

func (d QuackStrategy) quack() {
	print("quack quack")
}

type NoQuackStrategy struct {
}

func (d NoQuackStrategy) quack() {
	print("No quack")
}

type Duck struct {
	IFlyBehaviour   IFlyBehaviour
	IQuackBehaviour IQuackBehaviour
	Name            string
}

func (d Duck) MakeDuckFly() {
	d.Fly()
}

func (d Duck) Fly() {
	d.IFlyBehaviour.fly()
}

func newDuck(FlyBehaviour IFlyBehaviour, QuackBehaviour IQuackBehaviour) *Duck {
	d := Duck{}
	d.Name = "duck"
	d.IFlyBehaviour = FlyBehaviour
	d.IQuackBehaviour = QuackBehaviour
	return &d
}

func main() {
	straightFly := &StraightFlyStrategy{}
	quack := &NoQuackStrategy{}
	duck := newDuck(*straightFly, *quack)
	duck.Fly()
}
