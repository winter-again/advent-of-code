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

    ans = 0
    with open(file) as f:
        g = [line.strip() for line in f]

    dirs = [
        (1, 0),
        (1, 1),
        (0, 1),
        (-1, 1),
        (-1, 0),
        (-1, -1),
        (0, -1),
        (1, -1),
    ]
    m = len(g)
    n = len(g[0])

    for i, row in enumerate(g):
        for j, _ in enumerate(row):
            if g[i][j] == "@":
                rolls = 0
                for di, dj in dirs:
                    ni = i + di
                    nj = j + dj
                    if 0 <= ni < m and 0 <= nj < n and g[ni][nj] == "@":
                        rolls += 1

                if rolls < 4:
                    ans += 1

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    with open(file) as f:
        g = [list(line.strip()) for line in f]

    dirs = [
        (1, 0),
        (1, 1),
        (0, 1),
        (-1, 1),
        (-1, 0),
        (-1, -1),
        (0, -1),
        (1, -1),
    ]
    m = len(g)
    n = len(g[0])

    while True:
        removed = 0
        for i, row in enumerate(g):
            for j, _ in enumerate(row):
                if g[i][j] == "@":
                    rolls = 0
                    for di, dj in dirs:
                        ni = i + di
                        nj = j + dj
                        if 0 <= ni < m and 0 <= nj < n and g[ni][nj] == "@":
                            rolls += 1

                    if rolls < 4:
                        g[i][j] = "x"
                        ans += 1
                        removed += 1
        if removed == 0:
            break

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
