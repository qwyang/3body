def move(n, src, dst):
    print 'move %dth from %s to %s dst' % (n, src, dst)


def hanoi(src, dst, h, n):
    if n == 0:
        return
    hanoi(src, h, dst, n - 1)
    move(n, src, dst)
    hanoi(h, dst, src, n - 1)


if __name__ == "__main__":
    hanoi("A", "C", "B", 4)
