from typing import Any, Union


def mapconcat(fn: Any, seq: Union[list[Union[str, int]], str], sep: str) -> str:
    return sep.join(fn(e) for e in seq)


def calc_cubic(c: str) -> str:
    return c * 3


def rjust_str(s: str) -> str:
    return s.rjust(10)


if __name__ == "__main__":
    print(mapconcat(str, ["foo", "bar", "baz"], "-"))
    print(mapconcat(str, [1, 2, 3], " "))
    print(mapconcat(calc_cubic, "abc", ""))
    print(mapconcat(rjust_str, ["foo", "bar", "baz"], ""))
