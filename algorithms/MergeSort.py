# a great example for devide and conquer algorithm
from utils import random_array


def merge(x, y):
    merged = []
    i = 0
    j = 0
    length_x = len(x)
    length_y = len(y)
    while True:
        if i >= length_x or j >= length_y:
            break
        if x[i] < y[j]:
            merged.append(x[i])
            i += 1
        else:
            merged.append(y[j])
            j += 1
    while i < length_x:
        merged.append(x[i])
        i += 1
    while j < length_y:
        merged.append(y[j])
        j += 1
    return merged


def mergeSort(array, start, end):
    if start == end:
        return [array[start]]
    # devide
    middle = (start + end) / 2
    # conquer
    x = mergeSort(array, start, middle)
    y = mergeSort(array, middle + 1, end)
    # merge
    return merge(x, y)


if __name__ == '__main__':
    array = random_array(20)
    print "before:", array
    result = mergeSort(array, 0, len(array) - 1)
    print "after:", result
