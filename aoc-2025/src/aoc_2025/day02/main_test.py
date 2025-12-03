from .main import part_1, part_2


def test_part_1_sample() -> None:
    ans = part_1(True)
    assert ans == 1227775554


def test_part_1() -> None:
    ans = part_1()
    assert ans == 19605500130


def test_part_2_sample() -> None:
    ans = part_2(True)
    assert ans == 4174379265


def test_part_2() -> None:
    ans = part_2()
    assert ans == 36862281418
