# LSB(Least Significant Bit first) algorithm.
from utils import random_array


def get_total_digit_num(d, base=10):
    n = 0
    while d:
        d /= base
        n += 1
    return n


def get_digit(d, n, base=10):
    for i in range(n):
        d /= base
    return d % base


def radixSort(array, base=10):
    buckets = [[] for _ in range(base)]
    n = get_total_digit_num(max(array))
    for i in range(n):  # n times distribute and collect and then the array will be in order.
        # distribute data to buckets
        for item in array:
            index = get_digit(item, i)
            buckets[index].append(item)
        # collect data from buckets
        count = 0
        for bucket in buckets:
            for item in bucket:
                array[count] = item
                count += 1
        buckets = [[] for _ in range(base)]


if __name__ == "__main__":
    array = random_array(20)
    print "before:", array
    radixSort(array)
    print "after:", array
