class LinkNode(object):
    def __init__(self, data, left=None, right=None):
        self.data = data
        self.left = left
        self.right = right

    def __eq__(self, other):
        if isinstance(other, LinkNode):
            return id(other) == id(self)
        else:
            return False


class DeQueue(object):
    def __init__(self):
        self._left = None
        self._right = None

    def append_left(self, data):
        s = LinkNode(data, None, None)
        if self._left is None and self._right is None:
            self._left = s
            self._right = s
        else:
            self._left.left = s
            s.right = self._left
            self._left = s

    def append_right(self, data):
        s = LinkNode(data, None, None)
        if self._left is None and self._right is None:
            self._left = s
            self._right = s
        else:
            self._right.right = s
            s.left = self._right
            self._right = s

    def pop_left(self):
        if self._left == self._right:
            tmp = self._left
            self._right = None
            self._left = None
        else:
            tmp = self._left
            self._left = tmp.right
        return tmp.data

    def pop_right(self):
        if self._right == self._left:
            tmp = self._right
            self._right = None
            self._left = None
        else:
            tmp = self._right
            self._right = tmp.left
        return tmp.data

    def empty(self):
        return self._right is None and self._left is None


if __name__ == "__main__":
    dq = DeQueue()
    dq.append_left(1)
    dq.append_left(2)
    dq.append_left(3)
    dq.append_right(4)
    dq.append_right(5)
    dq.append_right(6)
    dq.append_left(0)
    dq.append_right(7)
    while not dq.empty():
        data = dq.pop_right()
        print data,