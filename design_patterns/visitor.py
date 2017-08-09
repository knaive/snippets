from abc import ABCMeta, abstractmethod


class Visitor(object):
    __metaclass__ = ABCMeta

    @abstractmethod
    def visit(self, container):
        pass


class Child(Visitor):
    def visit(self, container):
        name = 'visit_' + type(container).__name__.lower()
        if hasattr(self, name):
            visit_method = getattr(self, name)
            visit_method(container)
    
    def visit_zoo(self, container):
        print container.get_annimals()

    def visit_school(self, container):
        print container.greet_teachers()


class Visitable(object):
    __metaclass__ = ABCMeta

    @abstractmethod
    def accept(self, visitor):
        pass


class School(Visitable):
    def __init__(self):
        self.address = 'Henan road 19'

    def accept(self, visitor):
        visitor.visit(self)
    
    def greet_teachers(self):
        return 'hello'


class Zoo(Visitable):
    def __init__(self):
        self.address = 'Shanghai road 19'

    def accept(self, visitor):
        visitor.visit(self)
    
    def get_annimals(self):
        return ['monkey', 'tiger', 'pika', 'rabbit']


def main():
    zoo = Zoo()
    school = School()
    child = Child()
    zoo.accept(child)
    school.accept(child)


if __name__ == '__main__':
    main()
