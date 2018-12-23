import random


def random_array(size):
    result = []
    for i in range(size):
        result.append(random.randint(1, 10))
    return result
