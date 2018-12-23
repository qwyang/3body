"""
heap sort.
"""


class heap(object):
    def __init__(self, array):
        self.data = array
        self.n = len(array)
        self.__heapify()

    def __heapify(self):
        for i in range(self.n / 2 - 1, -1, -1):
            self.__justify(i)

    def __justify(self, i):
        if 2 * i + 2 < self.n:  # have two children
            pos = 2 * i + 1 if self.data[2 * i + 1] <= self.data[2 * i + 2] else 2 * i + 2
            if self.data[pos] < self.data[i]:
                self.data[pos], self.data[i] = self.data[i], self.data[pos]
            self.__justify(pos)
        elif 2 * i + 1 < self.n:  # have one child
            if self.data[2 * i + 1] < self.data[i]:
                self.data[2 * i + 1], self.data[i] = self.data[i], self.data[2 * i + 1]
        else:  # have no children
            pass

    def pop(self):
        tmp = self.data[0]
        self.data[0] = self.data[self.n - 1]
        del self.data[self.n - 1]
        self.n -= 1
        self.__justify(0)
        return tmp

    def append(self, d):
        self.data.append(d)
        self.n += 1
        i = self.n - 1
        p = (i - 1) / 2
        while p >= 0:
            if self.data[i] < self.data[p]:
                self.data[i], self.data[p] = self.data[p], self.data[i]
            i = p
            p = (i - 1) / 2

    def empty(self):
        return self.n == 0

    def __str__(self):
        return str(self.data)


if __name__ == "__main__":
    a = []
    h = heap(a)
    h.append(1)
    h.append(5)
    h.append(4)
    h.append(3)
    h.append(2)
    while not h.empty():
        print h.pop(),
