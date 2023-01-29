from abc import ABC, abstractmethod


class IFlyBehaviour(ABC):
    @abstractmethod
    def fly(self):
        pass


class StraightFlyBehaviour(IFlyBehaviour):
    def fly(self):
        print("Straight Fly!!!!")


class ZigzagFlyBehaviour(IFlyBehaviour):
    def fly(self):
        print("ZigZag Fly!!!!")


class IQuackBehaviour(ABC):
    @abstractmethod
    def quack(self):
        pass


class QuackBehaviour(IQuackBehaviour):
    def quack(self):
        print("Quack Quack!!!!")


class NoQuackBehaviour(IQuackBehaviour):
    def quack(self):
        print("No Quack Quack!!!!")
