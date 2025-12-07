from .main import part_1, part_2


def test_part_1_sample() -> None:
    ans = part_1(True)
    assert ans == 21


def test_part_1() -> None:
    ans = part_1()
    assert ans == 1662


def test_part_2_sample() -> None:
    ans = part_2(True)
    assert ans == 40


def test_part_2() -> None:
    ans = part_2()
    assert ans == 40941112789504
