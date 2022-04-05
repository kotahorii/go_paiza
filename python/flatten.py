from typing import Any


def flatten(ls: list[Any]):
    def rec(xs: list[Any], r: list[int]):
        for e in xs:
            if type(e) is list:
                rec(e, r)
            else:
                r.append(e)
        return r

    return rec(ls, [])


if __name__ == "__main__":
    print(flatten([1, 2, [3, 4]]))
    print(flatten([1, [2, 3], [4, [5]]]))
