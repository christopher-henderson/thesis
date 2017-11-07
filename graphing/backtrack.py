import subprocess
import shlex
import time
from index import INDEX

GRAPH = set()

class Node(object):

    NODE_ID = 1

    NODE_NODES = dict()

    NODE_WHITE = 'white'
    NODE_GREEN = '#839356'
    NODE_RED = '#CD594A'
    NODE_BLUE = '#3C6478'

    def __init__(self, node, parent=None):
        self.NODE_parent = parent
        self.NODE_id = self.NODE_ID
        Node.NODE_ID += 1
        self.NODE_node = node
        self.NODE_color = self.NODE_BLUE
        self.NODE_NODES[self.NODE_id] = self

    def __getattr__(self, attr):
        return self.NODE_node.__getattribute__(attr)

    def __repr__(self):
        return self.NODE_node.__repr__()

    def __str__(self):
        return self.NODE_node.__str__()

    def __unicode__(self):
        return self.NODE_node.__unicode__()

    @classmethod
    def NODE__OUTPUT(cls):
        nodes = []
        edges = []
        for _, node in cls.NODE_NODES.items():

            nodes.append({
                'id': node.NODE_id, 
                'label': str(node), 
                'color': node.NODE_color,
            })
            if node.NODE_parent is not None:
                edges.append({"from": node.NODE_parent, "to": node.NODE_id})
        rendered_html = INDEX.replace('%%EDGES%%', str(edges)).replace('%%NODES%%', str(nodes))
        with open('index.html', 'w+') as index:
            index.write(rendered_html)
        subprocess.Popen(["open  index.html"], shell=True)


def _add(solution, candidate):
    solution.append(candidate)

def _remove(solution):
    solution.pop()

def _output(solution):
    print(solution) 


def __accept(accept_function):
    def inner(solution):
        if accept_function(solution):
            for node in solution:
                node.NODE_color = Node.NODE_GREEN
            return True
        return False
    return inner

def __reject(reject_function):
    def inner(solution, candidate):
        if reject_function(solution, candidate.NODE_node):
            candidate.NODE_color = Node.NODE_RED
            return True
        return False
    return inner

def __child(child_function):
    def inner(root):
        child = child_function(root.NODE_node)
        return None if child is None else Node(child, parent=root.NODE_id)
    return inner

def __sibling(sibling_function):
    def inner(node):
        sibling = sibling_function(node.NODE_node)
        return None if sibling is None else Node(sibling, parent=node.NODE_parent)
    return inner

def __append(append_function):
    def inner(solution, candidate):
        solution.append(candidate)
        if append_function is not None:
            append_function(candidate.NODE_node)
    return inner

ACCEPT = '__accept__'
REJECT = '__reject__'
CHILD = '__child__'
SIBLING = '__sibling__'
ADD = '__append__'
REMOVE = '__remove__'
OUTPUT = '__output__'
SINGLE_SOLUTION = '__single_solution__'


def backtrack(root):
    # Procedure resolutions.
    accept = __accept(getattr(type(root), ACCEPT))
    reject = __reject(getattr(type(root), REJECT))
    child = __child(getattr(type(root), CHILD))
    sibling = __sibling(getattr(type(root), SIBLING))
    add = __append(getattr(type(root), ADD) if hasattr(root, ADD) else None)
    remove = getattr(type(root), REMOVE) if hasattr(root, REMOVE) else _remove
    output = getattr(type(root), OUTPUT) if hasattr(root, OUTPUT) else _output

    single_solution = getattr(type(root), SINGLE_SOLUTION) if hasattr(root, SINGLE_SOLUTION) else False

    # Stack initialization.
    root_stack = list()
    solution = list()

    root = Node(root)

    # Algorithm.
    add(solution, root)
    while root is not None: 
        candidate = child(root)
        while candidate is not None:
            if reject(solution, candidate):
                candidate = sibling(candidate)
                continue
            add(solution, candidate)
            if accept(solution):
                output(solution)
                if single_solution:
                    return
                remove(solution)
                candidate = sibling(candidate)
                continue
            root_stack.append(root)
            root = candidate
            break
        else:
            root = sibling(root)
            remove(solution)
            while root_stack:
                if root is None:
                    root = sibling(root_stack.pop())
                    remove(solution)
                    continue
                if reject(solution, root):
                    root = sibling(root)
                    continue
                break
            if root is not None:
                add(solution, root)
    Node.NODE__OUTPUT()
    