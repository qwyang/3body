import json
from copy import deepcopy


def dict_merge(a, b):
    """recursively merges dict's. not just simple a['key'] = b['key'], if
    both a and bhave a key who's value is a dict then dict_merge is called
    on both values and the result stored in the returned dictionary."""
    if not isinstance(b, dict):
        return b
    result = deepcopy(a)
    for k, v in b.iteritems():
        if k in result and isinstance(result[k], dict):
            result[k] = dict_merge(result[k], v)
        else:
            result[k] = deepcopy(v)
    return result


if __name__ == "__main__":
    a = dict(a_only=42, a_and_b=43, a_only_dict={'a': 44}, a_and_b_dict={'a_o': 45, 'a_a_b': 46})
    b = dict(b_only=45, a_and_b=46, b_only_dict={'a': 47}, a_and_b_dict={'b_o': 48, 'a_a_b': 49})
    c = dict_merge(a, b)
    print json.dumps(c, indent=4)
