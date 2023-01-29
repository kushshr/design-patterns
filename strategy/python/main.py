from behaviours import StraightFlyBehaviour, QuackBehaviour, NoQuackBehaviour, ZigzagFlyBehaviour
from ducks import Duck

if __name__ == '__main__':
    wild_duck = Duck(fly_behaviour=StraightFlyBehaviour(), quack_behaviour=QuackBehaviour())
    wild_duck.fly()
    wild_duck.quack()
    rubber_duck = Duck(fly_behaviour=ZigzagFlyBehaviour(), quack_behaviour=NoQuackBehaviour())
    rubber_duck.fly()
    rubber_duck.quack()
    rubber_duck.quack_behaviour = QuackBehaviour()
    rubber_duck.quack()
