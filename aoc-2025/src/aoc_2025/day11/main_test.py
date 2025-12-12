from .main import part_1, part_2


def test_part_1_sample() -> None:
    ans = part_1(True)
    assert ans == 5


def test_part_1() -> None:
    ans = part_1()
    assert ans == 662


def test_part_2_sample() -> None:
    ans = part_2(True)
    assert ans == 2


def test_part_2() -> None:
    ans = part_2()
    assert ans == 429399933071120
