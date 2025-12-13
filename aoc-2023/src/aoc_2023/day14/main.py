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

    with open(file) as f:
        g = tuple(tuple(row.strip()) for row in f.readlines())

    print(g)

    gg = tilt(g, "north")

    ans = 0
    m = len(gg)
    for i, row in enumerate(gg):
        n_rocks = row.count("O")
        ans += n_rocks * (m - i)

    return ans


def tilt(g: tuple[tuple[str, ...], ...], dir: str) -> tuple[tuple[str, ...], ...]:
    dirs = {
        "north": (-1, 0),
        "south": (1, 0),
        "east": (0, 1),
        "west": (0, -1),
    }
    di, dj = dirs[dir]

    m = len(g)
    n = len(g[0])

    start_i = 0
    stop_i = m
    start_j = 0
    stop_j = n
    step_i = 1
    step_j = 1
    if dir == "south":
        start_i = m - 1
        stop_i = -1
        step_i = -1
    elif dir == "east":
        start_j = n - 1
        stop_j = -1
        step_j = -1

    gl = list(list(e) for e in g)

    for i in range(start_i, stop_i, step_i):
        for j in range(start_j, stop_j, step_j):
            if gl[i][j] != "O":
                continue

            pi, pj = i, j
            ni, nj = i + di, j + dj
            while ni >= 0 and ni < m and nj >= 0 and nj < n and gl[ni][nj] == ".":
                gl[ni][nj] = "O"
                gl[pi][pj] = "."

                pi, pj = pi + di, pj + dj
                ni, nj = ni + di, nj + dj

    gt = tuple(tuple(e) for e in gl)

    return gt


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        g = tuple(tuple(row.strip()) for row in f.readlines())

    cycles = 1_000_000_000
    seen = {g}
    g_seq = [g]
    iter = 0  # tracks iter of first encounter of repeated state in the cycle
    while True:
        iter += 1
        for dir in ["north", "west", "south", "east"]:
            g = tilt(g, dir)

        if g in seen:
            break

        seen.add(g)
        g_seq.append(g)

    cycle_offset = g_seq.index(g)
    # cycle every iter - cycle_offset iterations
    final = g_seq[cycle_offset + (cycles - cycle_offset) % (iter - cycle_offset)]

    ans = 0
    m = len(final)
    for i, row in enumerate(final):
        n_rocks = row.count("O")
        ans += n_rocks * (m - i)

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
