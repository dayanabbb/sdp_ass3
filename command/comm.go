from abc import ABC, abstractmethod


class Command(ABC):
    @abstractmethod
    def execute(self):
        pass


class Chef:
    def prepare_soup(self):
        print("chef is preparing soup")

    def prepare_besh(self):
        print("chef is grilling besh")


class SoupCommand(Command):
    def execute(self):
        Chef().prepare_soup()

class BeshCommand(Command):
    def execute(self):
        Chef().prepare_besh()


class Waiter:
    def __init__(self):
        self.order = None

    def take_order(self, order):
        self.order = order

    def place_order(self):
        self.order.execute()



if __name__ == "__main__":
    waiter = Waiter()

 
    waiter.take_order(SoupCommand())
    waiter.place_order()

    
    waiter.take_order(BeshCommand())
    waiter.place_order()
