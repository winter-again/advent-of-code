from .main import part_1, part_2


def test_part_1_sample() -> None:
    ans = part_1(True)
    assert ans == 357


def test_part_1() -> None:
    ans = part_1()
    assert ans == 17100


def test_part_2_sample() -> None:
    ans = part_2(True)
    assert ans == 3121910778619


def test_part_2() -> None:
    ans = part_2()
    assert ans == 170418192256861
