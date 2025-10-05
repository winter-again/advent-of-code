from .main import part_1, part_2


def test_part_1_sample() -> None:
    ans = part_1(True)
    assert ans == 374


def test_part_1() -> None:
    ans = part_1()
    assert ans == 9521776


def test_part_2() -> None:
    ans = part_2()
    assert ans == 553224415344
