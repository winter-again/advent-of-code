import re
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

    tot = 0
    with open(file) as f:
        for line in f:
            nums = [c for c in line if c.isdigit()]
            tot += int(nums[0] + nums[-1])

    return tot


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl_2.txt"

    mapper = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    tot = 0
    with open(file) as f:
        for line in f:
            matches: list[str] = re.findall(
                r"(?=([0-9]|one|two|three|four|five|six|seven|eight|nine))", line
            )
            nums = [int(x) if x.isdigit() else mapper[x] for x in matches]
            tot += nums[0] * 10 + nums[-1]

    return tot


if __name__ == "__main__":
    raise SystemExit(main())
