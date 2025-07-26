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

    g = []
    start = (0, 0, 0)
    # NOTE: turns out don't need to identify S to solve this part,
    # but vis set won't be useful for part 2 since it will consider neighbors
    # of S that shouldn't be possible
    with open(file) as f:
        for i, line in enumerate(f):
            for j, c in enumerate(line.strip()):
                if c == "S":
                    start = (i, j, 0)

            row = line.strip()
            g.append(row)

    ans, _ = fill(start, g)
    return ans


dirs = [(-1, 0), (0, 1), (1, 0), (0, -1)]


def fill(
    s: tuple[int, int, int], g: list[list[str]]
) -> tuple[int, set[tuple[int, int]]]:
    si, sj, _ = s
    q = collections.deque([s])
    vis = set()
    vis.add((si, sj))

    res = 0
    while len(q) > 0:
        i, j, steps = q.popleft()
        res = steps
        for dd in dirs:
            di, dj = dd
            ni = i + di
            nj = j + dj
            if (
                ni >= 0
                and ni < len(g)
                and nj >= 0
                and nj < len(g[0])
                and (ni, nj) not in vis
                and g[ni][nj] != "."
            ):
                if dd == (-1, 0) and not (g[i][j] in "S|JL" and g[ni][nj] in "|7F"):
                    continue
                if dd == (0, 1) and not (g[i][j] in "S-LF" and g[ni][nj] in "-J7"):
                    continue
                if dd == (1, 0) and not (g[i][j] in "S|F7" and g[ni][nj] in "|LJ"):
                    continue
                if dd == (0, -1) and not (g[i][j] in "S-J7" and g[ni][nj] in "-LF"):
                    continue

                vis.add((ni, nj))
                q.append((ni, nj, steps + 1))

    return res, vis


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        # NOTE: there's actually 2 sample inputs given; testing against last sample
        file = parent / "input_smpl_2.txt"

    g = []
    start = (0, 0, 0)
    with open(file) as f:
        for i, line in enumerate(f):
            for j, c in enumerate(line.strip()):
                if c == "S":
                    start = (i, j, 0)

            row = line.strip()
            g.append([c for c in row])

    si, sj, _ = start
    s_real = {"|", "-", "L", "J", "7", "F"}
    for d in dirs:
        di, dj = d
        ni, nj = si + di, sj + dj
        nc = g[ni][nj]
        if d == (-1, 0) and nc in "|F7":
            s_real &= {"|", "L", "J"}
        elif d == (0, 1) and nc in "-J7":
            s_real &= {"-", "L", "F"}
        elif d == (1, 0) and nc in "|LJ":
            s_real &= {"|", "F", "7"}
        elif d == (0, -1) and nc in "-LF":
            s_real &= {"-", "7", "J"}

    assert len(s_real) == 1
    s_real = list(s_real)[0]
    g[start[0]][start[1]] = s_real
    _, vis = fill(start, g)

    for i, row in enumerate(g):
        for j in range(len(row)):
            if (i, j) not in vis:
                g[i][j] = "."

    ans = 0
    for i, row in enumerate(g):
        for j in range(len(row)):
            if (i, j) not in vis:
                if ray_cast((i, j), g, vis):
                    ans += 1
                    g[i][j] = "#"

    # print("\n".join(["".join(row) for row in g]))

    return ans


def ray_cast(start: tuple[int, int], g: list[str], vis: set[tuple[int, int]]) -> bool:
    """
    True if inside, else False
    | -> cross 1
    F-...-J -> cross 1
    L-...-7 -> cross 1
    """
    si, sj = start
    n = len(g[0])

    cross_r = 0
    # NOTE: return `inside` after ray cast to right is actually sufficient
    inside = False
    enter_F = None
    for j in range(sj, n):
        if (si, j) not in vis:
            continue

        c = g[si][j]
        if c == "|":
            assert enter_F is None
            inside = not inside
            cross_r += 1
        elif c == "-":
            assert enter_F is not None
        elif c in "LF":
            # enter pipe
            assert enter_F is None, f"{enter_F=}"
            enter_F = c == "F"
        elif c in "7J":
            # exit pipe
            assert enter_F is not None
            if (enter_F and c == "J") or (not enter_F and c == "7"):
                cross_r += 1
                inside = not inside
            enter_F = None

    cross_d = 0
    inside = False
    enter_F = None
    m = len(g)
    for i in range(si, m):
        if (i, sj) not in vis:
            continue

        c = g[i][sj]
        if c == "-":
            assert enter_F is None
            inside = not inside
            cross_d += 1
        elif c == "|":
            assert enter_F is not None
        elif c in "F7":
            assert enter_F is None, f"{enter_F=}"
            enter_F = c == "F"
        elif c in "JL":
            assert enter_F is not None
            if (enter_F and c == "J") or (not enter_F and c == "L"):
                cross_d += 1
                inside = not inside
            enter_F = None

    # NOTE: odd = inside, even and > 0 = outside
    inside_r = cross_r > 0 and cross_r % 2 != 0
    inside_d = cross_d > 0 and cross_d % 2 != 0
    if inside_r and inside_d:
        return True
    return False


if __name__ == "__main__":
    raise SystemExit(main())
