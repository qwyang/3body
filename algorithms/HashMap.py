import math


def hash(key, mask):
    h = 0;
    for c in key:
        h = (h << 4) + ord(c);
        g = h & 0xF0000000L;
        if g:
            h ^= g >> 24;
        h &= ~g;
    return h & mask;


class HashMap(object):
    def __init__(self, capacity=16, load_factor=1):
        self.capacity = int(pow(2, int(math.log(capacity, 2) + 0.99)))
        self.mask = self.capacity - 1
        self.load_factor = load_factor
        self.size = 0
        self.maximum = int(self.load_factor * self.capacity)
        self.container = [[] for i in range(self.capacity)]

    def put(self, key, value):
        pos = hash(key, self.mask)
        for item in self.container[pos]:
            if item[0] == key:
                item[1] = value
                break
        else:
            self.container[pos].append([key, value])
            self.size += 1
        if self.size > self.maximum:
            self.resize()

    def pop(self, key):
        pos = hash(key, self.mask)
        for item in self.container[pos]:
            if item[0] == key:
                self.size -= 1
                return item[1]
        else:
            return None

    def get(self, key):
        pos = hash(key, self.mask)
        for item in self.container[pos]:
            if item[0] == key:
                return item[1]
        else:
            return None

    def resize(self):
        self.capacity *= 2
        self.mask = self.capacity - 1
        self.maximum = int(self.capacity * self.load_factor)
        container = [[] for i in range(self.capacity)]
        for link in self.container:
            for item in link:
                pos = hash(item[0], self.mask)
                container[pos].append(item)
        del self.container
        self.container = container

    def __str__(self):
        return "capacity:%s, load_factor:%s, max:%s, mask:0x%x" % (
        self.capacity, self.load_factor, self.maximum, self.mask)


if __name__ == "__main__":
    import json

    hm = HashMap(4, 1)
    for k, v in [('x', 2), ('x', 1), ('y', 1), ('a', 1), ('b', 2), ('c', 3), ('d', 4), ('e', 5), ('f', 6), ('g', 7)]:
        hm.put(k, v)
    print hm.get('g')
    print hm
    print json.dumps(hm.container, indent=4)
