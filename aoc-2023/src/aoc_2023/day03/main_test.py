from .main import part_1, part_2


def test_part_1_sample() -> None:
    ans = part_1(True)
    assert ans == 4361


def test_part_1() -> None:
    ans = part_1()
    assert ans == 528799


def test_part_2_sample() -> None:
    ans = part_2(True)
    assert ans == 467835


def test_part_2() -> None:
    ans = part_2()
    assert ans == 84907174
