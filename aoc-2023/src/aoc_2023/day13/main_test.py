from .main import part_1, part_2


def test_part_1_sample() -> None:
    ans = part_1(True)
    assert ans == 405


def test_part_1() -> None:
    ans = part_1()
    assert ans == 34100


def test_part_2_sample() -> None:
    ans = part_2(True)
    assert ans == 400


def test_part_2() -> None:
    ans = part_2()
    assert ans == 33106
