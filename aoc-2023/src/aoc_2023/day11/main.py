import argparse
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

    g = []
    with open(file) as f:
        for i, line in enumerate(f):
            g.append(line.strip())

    row_exp = set()
    for i, row in enumerate(g):
        if row == "." * len(row):
            row_exp.add(i)

    col_exp = set()
    for j, col in enumerate(zip(*g)):
        if all(c == "." for c in col):
            col_exp.add(j)

    galaxies = []
    for i, row in enumerate(g):
        for j, c in enumerate(row):
            if c == "#":
                galaxies.append((i, j))

    ans = 0
    expand = 2
    for i, (oi, oj) in enumerate(galaxies):
        for di, dj in galaxies[:i]:
            for r in range(min(oi, di), max(oi, di)):
                ans += expand if r in row_exp else 1
            for c in range(min(oj, dj), max(oj, dj)):
                ans += expand if c in col_exp else 1

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl_2.txt"

    g = []
    with open(file) as f:
        for i, line in enumerate(f):
            g.append(line.strip())

    row_exp = set()
    for i, row in enumerate(g):
        if row == "." * len(row):
            row_exp.add(i)

    col_exp = set()
    for j, col in enumerate(zip(*g)):
        if all(c == "." for c in col):
            col_exp.add(j)

    galaxies = []
    for i, row in enumerate(g):
        for j, c in enumerate(row):
            if c == "#":
                galaxies.append((i, j))

    ans = 0
    expand = 1000000
    for i, (oi, oj) in enumerate(galaxies):
        for di, dj in galaxies[:i]:
            for r in range(min(oi, di), max(oi, di)):
                ans += expand if r in row_exp else 1
            for c in range(min(oj, dj), max(oj, dj)):
                ans += expand if c in col_exp else 1

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
