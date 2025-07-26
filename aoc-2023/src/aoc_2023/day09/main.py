import argparse
from collections.abc import Sequence
from pathlib import Path

parent = Path(__file__).parent


def main(argv: Sequence[str] | None = None) -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("part", nargs="?", default=0, type=int)
    args = parser.parse_args(argv)

    if args.part == 0:
        print(f"Part 1(sample): {part_1(True)}")
        print(f"Part 1: {part_1()}")

        print(f"Part 2 (sample): {part_2(True)}")
        print(f"Part 2: {part_2()}")
    elif args.part == 1:
        print(f"Part 1 (sample): {part_1(True)}")
        print(f"Part 1: {part_1()}")
    elif args.part == 2:
        print(f"Part 2 (sample): {part_2(True)}")
        print(f"Part 2: {part_2()}")

    return 0


def part_1(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    with open(file) as f:
        for line in f:
            nums = [int(n) for n in line.split()]
            ans += proc_p1(nums)

    return ans


def proc_p1(nums: list[int]) -> int:
    if len(set(nums)) == 1:
        return nums[0]

    new = [s - f for f, s in zip(nums, nums[1:])]
    return nums[-1] + proc_p1(new)


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    with open(file) as f:
        for line in f:
            nums = [int(n) for n in line.split()]
            ans += proc_p2(nums)

    return ans


def proc_p2(nums: list[int]) -> int:
    if len(set(nums)) == 1:
        return nums[0]

    new = [s - f for f, s in zip(nums, nums[1:])]
    return nums[0] - proc_p2(new)


if __name__ == "__main__":
    raise SystemExit(main())
