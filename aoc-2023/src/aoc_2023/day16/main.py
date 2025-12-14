import argparse
import collections
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

    with open(file) as f:
        g = [list(line) for line in f.read().splitlines()]

    m = len(g)
    n = len(g[0])
    dirs = {
        "^": (-1, 0),
        ">": (0, 1),
        "v": (1, 0),
        "<": (0, -1),
    }

    q = collections.deque([(">", 0, -1)])
    vis = set()

    while len(q) > 0:
        d, ci, cj = q.popleft()
        di, dj = dirs[d]
        ni, nj = ci + di, cj + dj

        if ni < 0 or ni >= m or nj < 0 or nj >= n:
            continue

        if (d, ni, nj) in vis:
            continue

        nxt = g[ni][nj]

        if nxt == ".":
            q.append((d, ni, nj))
        elif nxt == "/":
            if d == "^":
                q.append((">", ni, nj))
            elif d == ">":
                q.append(("^", ni, nj))
            elif d == "v":
                q.append(("<", ni, nj))
            elif d == "<":
                q.append(("v", ni, nj))
        elif nxt == "\\":
            if d == "^":
                q.append(("<", ni, nj))
            elif d == ">":
                q.append(("v", ni, nj))
            elif d == "v":
                q.append((">", ni, nj))
            elif d == "<":
                q.append(("^", ni, nj))
        elif nxt == "|":
            if d == ">" or d == "<":
                q.append(("^", ni, nj))
                q.append(("v", ni, nj))
            else:
                q.append((d, ni, nj))
        elif nxt == "-":
            if d == "^" or d == "v":
                q.append(("<", ni, nj))
                q.append((">", ni, nj))
            else:
                q.append((d, ni, nj))

        vis.add((d, ni, nj))

    ans = len({(i, j) for _, i, j in vis})

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        g = [list(line) for line in f.read().splitlines()]

    m = len(g)
    n = len(g[0])
    dirs = {
        "^": (-1, 0),
        ">": (0, 1),
        "v": (1, 0),
        "<": (0, -1),
    }

    starts = []
    for i in range(m):
        starts.append((">", i, -1))
        starts.append(("<", i, n))

    for j in range(n):
        starts.append(("v", -1, j))
        starts.append(("^", m, j))

    ans = 0
    for start in starts:
        q = collections.deque([start])
        vis = set()

        while len(q) > 0:
            d, ci, cj = q.popleft()
            di, dj = dirs[d]
            ni, nj = ci + di, cj + dj

            if ni < 0 or ni >= m or nj < 0 or nj >= n:
                continue

            if (d, ni, nj) in vis:
                continue

            nxt = g[ni][nj]

            if nxt == ".":
                q.append((d, ni, nj))
            elif nxt == "/":
                if d == "^":
                    q.append((">", ni, nj))
                elif d == ">":
                    q.append(("^", ni, nj))
                elif d == "v":
                    q.append(("<", ni, nj))
                elif d == "<":
                    q.append(("v", ni, nj))
            elif nxt == "\\":
                if d == "^":
                    q.append(("<", ni, nj))
                elif d == ">":
                    q.append(("v", ni, nj))
                elif d == "v":
                    q.append((">", ni, nj))
                elif d == "<":
                    q.append(("^", ni, nj))
            elif nxt == "|":
                if d == ">" or d == "<":
                    q.append(("^", ni, nj))
                    q.append(("v", ni, nj))
                else:
                    q.append((d, ni, nj))
            elif nxt == "-":
                if d == "^" or d == "v":
                    q.append(("<", ni, nj))
                    q.append((">", ni, nj))
                else:
                    q.append((d, ni, nj))

            vis.add((d, ni, nj))

        e = len({(i, j) for _, i, j in vis})
        ans = max(e, ans)

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
