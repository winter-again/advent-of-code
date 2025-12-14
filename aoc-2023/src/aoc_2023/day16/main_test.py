from .main import part_1, part_2


def test_part_1_sample() -> None:
    ans = part_1(True)
    assert ans == 46


def test_part_1() -> None:
    ans = part_1()
    assert ans == 7496


def test_part_2_sample() -> None:
    ans = part_2(True)
    assert ans == 51


def test_part_2() -> None:
    ans = part_2()
    assert ans == 7932
