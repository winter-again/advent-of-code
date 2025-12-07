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

    ans = 0
    with open(file) as f:
        g = [list(line.strip()) for line in f]

    si, sj = 0, 0
    for i, row in enumerate(g):
        for j, c in enumerate(row):
            if c == "S":
                si, sj = i + 1, j

    m = len(g)
    n = len(g[0])
    q = collections.deque([(si, sj)])
    beams = set()
    while len(q) > 0:
        ci, cj = q.popleft()
        ni, nj = ci + 1, cj
        if ni < 0 or ni >= m or nj < 0 or nj >= n:
            continue

        if g[ni][nj] == "^":
            lb = (ci + 1, cj - 1)
            rb = (ci + 1, cj + 1)

            if lb not in beams:
                q.append(lb)
                beams.add(lb)
            if rb not in beams:
                q.append(rb)
                beams.add(rb)

            ans += 1
        else:
            if (ni, nj) not in beams:
                q.append((ni, nj))
                beams.add((ni, nj))

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        g = [list(line.strip()) for line in f]

    si, sj = 0, 0
    for i, row in enumerate(g):
        for j, c in enumerate(row):
            if c == "S":
                si, sj = i + 1, j

    cache = {}
    return timelines((si, sj), g, cache)


def timelines(
    cur: tuple[int, int],
    g: list[list[str]],
    cache: dict[tuple[int, int], int],
) -> int:
    if cur in cache:
        return cache[cur]

    ci, cj = cur
    ni, nj = ci + 1, cj
    m = len(g)
    n = len(g[0])
    if ni < 0 or ni >= m or nj < 0 or nj >= n:
        return 1

    if g[ni][nj] == "^":
        left = timelines((ci + 1, cj - 1), g, cache)
        right = timelines((ci + 1, cj + 1), g, cache)
        cache[cur] = left + right
        return left + right
    else:
        a = timelines((ci + 1, cj), g, cache)
        cache[cur] = a
        return a


if __name__ == "__main__":
    raise SystemExit(main())
