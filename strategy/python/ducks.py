from behaviours import IFlyBehaviour, IQuackBehaviour


class Duck():
    def __init__(self, fly_behaviour: IFlyBehaviour, quack_behaviour: IQuackBehaviour) -> None:
        self._fly_behaviour = fly_behaviour
        self._quack_behaviour = quack_behaviour

    @property
    def fly_behaviour(self) -> IFlyBehaviour:
        return self._fly_behaviour

    @fly_behaviour.setter
    def fly_behaviour(self, behaviour: IFlyBehaviour) -> None:
        self._fly_behaviour = behaviour

    @property
    def quack_behaviour(self) -> IQuackBehaviour:
        return self._quack_behaviour

    @quack_behaviour.setter
    def quack_behaviour(self, behaviour: IQuackBehaviour) -> None:
        self._quack_behaviour = behaviour

    def fly(self):
        self._fly_behaviour.fly()

    def quack(self):
        self._quack_behaviour.quack()
