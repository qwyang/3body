from utils import random_array


def partition(array, start, end):
    r = start - 1  # the last element that is smaller than key
    key = array[end]
    for i in range(start, end):  # find out all the elements that are smaller than key and swap them before key
        if array[i] < key:
            array[r + 1], array[i] = array[i], array[r + 1]  # swap the smaller element to r+1 position
            r += 1  # r+1 position is now smaller than key, update r
    array[r + 1], array[end] = array[end], array[r + 1]  # last swap
    return r + 1


def qsort(array, start, end):
    if start >= end:
        return
    p = partition(array, start, end)
    qsort(array, start, p - 1)
    qsort(array, p + 1, end)


if __name__ == "__main__":
    array = random_array(10)
    qsort(array, 0, len(array) - 1)
    print array
