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
        g = [line.strip() for line in f]

    dirs = [(-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (1, -1), (0, -1), (-1, -1)]
    m = len(g)
    n = len(g[0])
    num_coords: set[tuple[int, int]] = set()
    for i in range(m):
        for j in range(n):
            if g[i][j].isdigit() or g[i][j] == ".":
                continue

            for di, dj in dirs:
                ni = i + di
                nj = j + dj
                if ni < 0 or ni >= m or nj < 0 or nj >= n or not g[ni][nj].isdigit():
                    continue
                while nj > 0 and g[ni][nj - 1].isdigit():
                    nj -= 1

                num_coords.add((ni, nj))

    ans = 0
    for i, j in num_coords:
        digits = ""
        while j < n and g[i][j].isdigit():
            digits += g[i][j]
            j += 1

        ans += int(digits)

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    with open(file) as f:
        g = [line.strip() for line in f]

    dirs = [(-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (1, -1), (0, -1), (-1, -1)]
    m = len(g)
    n = len(g[0])
    num_coords: set[tuple[int, int]] = set()
    cands = collections.defaultdict(set)
    for i in range(m):
        for j in range(n):
            if g[i][j] != "*":
                continue

            for di, dj in dirs:
                ni = i + di
                nj = j + dj
                if ni < 0 or ni >= m or nj < 0 or nj >= n or not g[ni][nj].isdigit():
                    continue
                while nj > 0 and g[ni][nj - 1].isdigit():
                    nj -= 1

                if (ni, nj) not in cands:
                    cands[(i, j)].add((ni, nj))

                num_coords.add((ni, nj))

    ans = 0
    for _, nums in cands.items():
        gears = []
        if len(nums) == 2:
            for i, j in nums:
                digits = ""
                while j < n and g[i][j].isdigit():
                    digits += g[i][j]
                    j += 1

                gears.append(int(digits))

        if len(gears) > 0:
            ans += gears[0] * gears[1]

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
