import argparse
import heapq
from collections.abc import Sequence
from pathlib import Path
from typing import NamedTuple

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


class Block(NamedTuple):
    heat: int
    i: int
    j: int
    di: int
    dj: int
    steps: int


def part_1(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        g = [[int(x) for x in row] for row in f.read().splitlines()]

    m = len(g)
    n = len(g[0])
    q = []
    heapq.heappush(q, Block(heat=0, i=0, j=0, di=0, dj=0, steps=0))
    vis = set()
    dirs = [(-1, 0), (0, 1), (1, 0), (0, -1)]

    ans = 0
    while len(q) > 0:
        heat, ci, cj, di, dj, steps = heapq.heappop(q)

        if (ci, cj) == (m - 1, n - 1):
            ans = heat
            break

        if (ci, cj, di, dj, steps) in vis:
            continue

        vis.add((ci, cj, di, dj, steps))

        for ndi, ndj in dirs:
            if (ndi, ndj) == (di, dj):
                if steps < 3:
                    ni, nj = ci + di, cj + dj
                    if 0 <= ni < m and 0 <= nj < n:
                        heapq.heappush(
                            q,
                            Block(
                                heat=heat + g[ni][nj],
                                i=ni,
                                j=nj,
                                di=di,
                                dj=dj,
                                steps=steps + 1,
                            ),
                        )
            else:
                if (ndi, ndj) != (-di, -dj):
                    ni, nj = ci + ndi, cj + ndj
                    if 0 <= ni < m and 0 <= nj < n:
                        heapq.heappush(
                            q,
                            Block(
                                heat=heat + g[ni][nj],
                                i=ni,
                                j=nj,
                                di=ndi,
                                dj=ndj,
                                steps=1,
                            ),
                        )

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        g = [[int(x) for x in row] for row in f.read().splitlines()]

    m = len(g)
    n = len(g[0])
    q = []
    heapq.heappush(q, Block(heat=0, i=0, j=0, di=0, dj=0, steps=0))
    vis = set()
    dirs = [(-1, 0), (0, 1), (1, 0), (0, -1)]

    ans = 0
    while len(q) > 0:
        heat, ci, cj, di, dj, steps = heapq.heappop(q)

        if (ci, cj) == (m - 1, n - 1) and steps >= 4:
            ans = heat
            break

        if (ci, cj, di, dj, steps) in vis:
            continue

        vis.add((ci, cj, di, dj, steps))

        for ndi, ndj in dirs:
            if (ndi, ndj) == (di, dj):
                if steps < 10:
                    ni, nj = ci + di, cj + dj
                    if 0 <= ni < m and 0 <= nj < n:
                        heapq.heappush(
                            q,
                            Block(
                                heat=heat + g[ni][nj],
                                i=ni,
                                j=nj,
                                di=di,
                                dj=dj,
                                steps=steps + 1,
                            ),
                        )
            else:
                if (di, dj) == (0, 0) or ((ndi, ndj) != (-di, -dj) and steps >= 4):
                    ni, nj = ci + ndi, cj + ndj
                    if 0 <= ni < m and 0 <= nj < n:
                        heapq.heappush(
                            q,
                            Block(
                                heat=heat + g[ni][nj],
                                i=ni,
                                j=nj,
                                di=ndi,
                                dj=ndj,
                                steps=1,
                            ),
                        )

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
