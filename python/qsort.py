from typing import Callable


def qsort(f: Callable[[int, int], bool], lst: list[int]) -> list[int]:
    if lst == []:
        return lst
    else:

        def partition(x: int, nums: list[int]) -> tuple[list[int], list[int]]:
            a: list[int] = []
            b: list[int] = []

            for num in nums:
                if f(num, x):
                    a.append(num)
                else:
                    b.append(num)
            return a, b

    xs, ys = partition(lst[0], lst[1:])
    return qsort(f, xs) + [lst[0]] + qsort(f, ys)


if __name__ == "__main__":
    asc = qsort(lambda x, y: x < y, [2, 4, -90, 3, 10])
    des = qsort(lambda x, y: x > y, [2, 4, -90, 3, 10])

    print(asc)
    print(des)
