import argparse
import collections
import functools
import operator
from collections.abc import Sequence
from pathlib import Path

parent = Path(__file__).parent


def main(argv: Sequence[str] | None = None) -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("part", nargs="?", default=0, type=int)
    parser.add_argument(
        "-s", "--sample-only", action=argparse.BooleanOptionalAction, default=False
    )
    args = parser.parse_args(argv)

    if args.part == 0:
        print(f"Part 1(sample): {part_1(True)}")
        print(f"Part 1: {part_1()}")

        print(f"Part 2 (sample): {part_2(True)}")
        print(f"Part 2: {part_2()}")
    elif args.part == 1:
        print(f"Part 1 (sample): {part_1(True)}")
        if not args.sample_only:
            print(f"Part 1: {part_1()}")
    elif args.part == 2:
        print(f"Part 2 (sample): {part_2(True)}")
        if not args.sample_only:
            print(f"Part 2: {part_2()}")

    return 0


def part_1(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    cols = collections.defaultdict(list)
    ops = []
    with open(file) as f:
        for line in f:
            ll = line.strip().split()
            if ll[0] == "+" or ll[0] == "*":
                ops = ll
                continue

            for i, c in enumerate(ll):
                cols[i].append(int(c))

    for col, op in zip(cols.values(), ops):
        if op == "+":
            ans += sum(col)
        else:
            ans += functools.reduce(operator.mul, col, 1)

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    with open(file) as f:
        lines = [line.strip("\n") for line in f]

    rows = [list(line) for line in lines[:-1]]
    ops = lines[-1].split()

    nums = collections.defaultdict(list)
    i = 0
    for col in zip(*rows):
        if all([x == " " for x in col]):
            i += 1
            continue

        nums[i].append(int("".join(col)))

    for block, op in zip(nums.values(), ops):
        if op == "+":
            ans += sum(block)
        else:
            ans += functools.reduce(operator.mul, block, 1)

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
