def permute(array, start, end):
    if start >= end:
        print array
    else:
        for i in range(start, end + 1):
            array[start], array[i] = array[i], array[start]
            permute(array, start + 1, end)
            array[start], array[i] = array[i], array[start]


class StackNode(object):
    def __init__(self, i, s, e):
        self.i = i
        self.s = s
        self.e = e

    def __str__(self):
        return "(%s,%s,%s)" % (self.i, self.s, self.e)


def p_stack(stack):
    for i in stack:
        print i
    print


def permute2(array, s, e):
    stack = []
    node = StackNode(s, s, e)
    stack.append(node)
    flag = False
    while stack:
        front = stack[len(stack) - 1]
        i, s, e = front.i, front.s, front.e
        if s == e:  # the stop condition comes, output one of the permutation.
            print array
            stack.pop()
            flag = True  # come out from a child permutation
        elif i == e:  # finish a child sequence permutation.
            # print '#2:swap %s, %s' % (s,i)
            array[s], array[i] = array[i], array[s]  # recover the sequence
            stack.pop()
            flag = True  # come out from a child permutation.
        else:  # must continue to push stack.
            if flag:  # a child permutation finished.
                # print '#2:swap %s, %s' % (s,i)
                array[s], array[i] = array[i], array[s]  # recover the sequence
                i += 1
                front.i = i
                # print '#1:swap %s, %s' % (s,i)
                array[s], array[i] = array[i], array[s]
                stack.append(StackNode(s + 1, s + 1, e))
                flag = False
            else:
                stack.append(StackNode(s + 1, s + 1, e))
                flag = False  # not finish any child permutation by default.
                # p_stack(stack)


if __name__ == "__main__":
    permute2([1, 2, 3, 4, 5, 6], 0, 5)
