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

    with open(file) as f:
        times, dists = [list(map(int, line.split(": ")[1].split())) for line in f]

    ways = 1
    for time, dist in zip(times, dists):
        n = 0
        for hold in range(1, time):
            d = hold * (time - hold)
            if d > dist:
                n += 1
        ways *= n

    return ways


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        data = [int(line.strip().split(": ")[1].strip().replace(" ", "")) for line in f]

    ways = 0
    for hold in range(1, data[0]):
        dist = hold * (data[0] - hold)
        if dist > data[1]:
            ways += 1

    return ways


if __name__ == "__main__":
    raise SystemExit(main())
