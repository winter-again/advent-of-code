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
        patterns = [
            list(map(list, lines.split("\n")))
            for lines in f.read().strip().split("\n\n")
        ]

    vert = 0
    hori = 0
    for pattern in patterns:
        pattern_T = list(map(list, zip(*pattern)))  # transpose
        v, _ = find_horizontal(pattern_T)
        vert += v
        h, _ = find_horizontal(pattern)
        hori += h

    return vert + 100 * hori


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        patterns = [
            list(map(list, lines.split("\n")))
            for lines in f.read().strip().split("\n\n")
        ]

    vert = 0
    hori = 0
    for pattern in patterns:
        pattern_T = list(map(list, zip(*pattern)))  # transpose
        v, diff = find_horizontal(pattern_T, True)
        if diff == 1:
            vert += v
            continue

        h, diff = find_horizontal(pattern, True)
        if diff == 1:
            hori += h

    return vert + 100 * hori


def find_horizontal(pattern: list[list[str]], smudge: bool = False) -> tuple[int, int]:
    m = len(pattern)
    for k in range(m - 1):
        u = k
        d = k + 1
        diff = 0
        for i in range(min(u, m - d) + 1):
            if u - i < 0 or d + i >= m:
                break

            for idx, (eu, ed) in enumerate(zip(pattern[u - i], pattern[d + i])):
                if eu != ed:
                    diff += 1

        targ = 1 if smudge else 0
        if diff == targ:
            return k + 1, diff

    return 0, 0


if __name__ == "__main__":
    raise SystemExit(main())
