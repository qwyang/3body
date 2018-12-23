class Node(object):
    def __init__(self, key, left=None, right=None):
        self.key = key
        self.height = 1
        self.left = left
        self.right = right

    def __str__(self):
        left = id(self.left) if self.left else None
        right = id(self.right) if self.right else None
        return "<key:%s,id:%s, left:%s,right:%s,height:%s>" % (self.key, id(self), left, right, self.height)


def Height(tree):
    if tree:
        return tree.height
    else:
        return 0


def RR_rotate(k1):
    k2 = k1.right
    k1.right = k2.left
    k2.left = k1
    k1.height = max(Height(k1.left), Height(k1.right)) + 1
    k2.height = max(Height(k1), Height(k2.right)) + 1
    return k2


def LL_rotate(k1):
    k2 = k1.left
    k1.left = k2.right
    k2.right = k1
    k1.height = max(Height(k1.left), Height(k1.right)) + 1
    k2.height = max(Height(k2.left), Height(k1)) + 1
    return k2


def RL_rotate(k1):
    k1.right = LL_rotate(k1.right)
    return RR_rotate(k1)


def LR_rotate(k1):
    k1.left = RR_rotate(k1.left)
    return LL_rotate(k1)


class AvlTree(object):
    def __init__(self):
        self.root = None

    def insert(self, key):
        node = Node(key)
        self.root = self._insert(self.root, node)

    def display(self):
        self._display(self.root)

    def remove(self, key):
        self._remove(self.root, key)

    def find(self, key):
        return self._find(self.root, key)

    def _find(self, tree, key):
        while tree:
            if key == tree.key:
                break
            elif key < tree.key:
                tree = tree.left
            else:
                tree = tree.right
        return tree

    def _remove(self, tree, key):
        if not tree:
            return None
        if key < tree.key:
            tree.left = self._remove(tree.left, key)
        elif key > tree.key:
            tree.right = self._remove(tree.right, key)
        else:
            if not tree.left:
                return tree.right
            if tree.right:
                tmp = tree.right
                while tmp.left:
                    tmp = tmp.left
                tree.key, tmp.key = tmp.key, tree.key
                tree.right = self._remove(tree.right, key)
        tree = self._balance(tree, key)
        tree.height = max(Height(tree.left), Height(tree.right)) + 1
        return tree

    def _balance(self, tree, key):
        balance = Height(tree.left) - Height(tree.right)
        if balance < -1:
            if key > tree.right.key:
                tree = RL_rotate(tree)
            else:
                tree = RR_rotate(tree)
        elif balance > 1:
            if key < tree.left.key:
                tree = LL_rotate(tree)
            else:
                tree = LR_rotate(tree)
        return tree

    def _display(self, tree):
        if tree:
            print tree
            self._display(tree.left)
            self._display(tree.right)

    def _insert(self, tree, node):
        if tree:
            if node.key < tree.key:
                tree.left = self._insert(tree.left, node)
            elif node.key > tree.key:
                tree.right = self._insert(tree.right, node)
            else:
                print "%s already exists." % node.key
                return tree
            balance = Height(tree.left) - Height(tree.right)
            if balance < -1:
                if node.key > tree.right.key:
                    tree = RR_rotate(tree)
                else:
                    tree = RL_rotate(tree)
            elif balance > 1:
                if node.key < tree.left.key:
                    tree = LL_rotate(tree)
                else:
                    tree = LR_rotate(tree)
            else:
                tree.height = max(Height(tree.left), Height(tree.right)) + 1
            return tree
        else:
            return node


if __name__ == "__main__":
    t = AvlTree()
    t.insert(1)
    t.insert(2)
    t.insert(3)
    t.insert(4)
    t.insert(5)
    t.insert(6)
    t.insert(7)
    t.insert(8)
    t.insert(9)
    t.insert(10)
    t.display()
    print '-' * 40
    t.remove(1)
    t.remove(4)
    t.remove(6)
    t.remove(7)
    t.display()
    print t.find(1)
    print t.find(10)
