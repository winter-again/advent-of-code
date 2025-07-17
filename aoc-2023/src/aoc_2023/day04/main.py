import argparse
import collections
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
            card = line.split(": ")[1]
            data = card.strip().split(" | ")
            nums = []
            nums.append({int(x) for x in data[0].split()})
            nums.append({int(x) for x in data[1].split()})

            wins = nums[0].intersection(nums[1])
            n = len(wins)
            if n > 0:
                ans += 2 ** (n - 1)

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    cards = collections.defaultdict(int)
    with open(file) as f:
        for i, line in enumerate(f, 1):
            cards[i] += 1
            card = line.split(": ")[1]
            data = card.strip().split(" | ")
            nums = []
            nums.append({int(x) for x in data[0].split()})
            nums.append({int(x) for x in data[1].split()})

            wins = len(nums[0].intersection(nums[1]))
            for j in range(1, wins + 1):
                for _ in range(cards[i]):
                    cards[i + j] += 1

    ans = sum(cards.values())
    return ans


if __name__ == "__main__":
    raise SystemExit(main())
