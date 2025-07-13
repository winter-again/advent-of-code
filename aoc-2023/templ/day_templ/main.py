from pathlib import Path

parent = Path(__file__).parent


def main() -> int:
    print(part_1(True))
    print(part_1())

    print(part_2(True))
    print(part_2())

    return 0


def part_1(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl_2.txt"

    ans = 0

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
