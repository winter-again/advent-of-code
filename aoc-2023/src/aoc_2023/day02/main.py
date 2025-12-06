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

    games = collections.defaultdict(list)
    with open(file) as f:
        for i, line in enumerate(f, 1):
            cubes = line.strip().split(": ")[1]
            cube_sets = [combo.split(", ") for combo in cubes.split("; ")]
            for cs in cube_sets:
                parsed = [0] * 3
                for ct in cs:
                    freq = ct.split(" ")
                    num = int(freq[0])
                    color = freq[1]
                    if color == "red":
                        parsed[0] = num
                    elif color == "green":
                        parsed[1] = num
                    elif color == "blue":
                        parsed[2] = num

                games[i].append(parsed)

    ans = 0
    for game, sets in games.items():
        ok = True
        for cubes in sets:
            if cubes[0] > 12 or cubes[1] > 13 or cubes[2] > 14:
                ok = False
                break

        if ok:
            ans += game

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    games = collections.defaultdict(list)
    with open(file) as f:
        for i, line in enumerate(f, 1):
            cubes = line.strip().split(": ")[1]
            cube_sets = [combo.split(", ") for combo in cubes.split("; ")]
            for cs in cube_sets:
                parsed = [0] * 3
                for ct in cs:
                    freq = ct.split(" ")
                    num = int(freq[0])
                    color = freq[1]
                    if color == "red":
                        parsed[0] = num
                    elif color == "green":
                        parsed[1] = num
                    elif color == "blue":
                        parsed[2] = num

                games[i].append(parsed)

    ans = 0
    for sets in games.values():
        r_max = 0
        g_max = 0
        b_max = 0
        for cubes in sets:
            r_max = max(r_max, cubes[0])
            g_max = max(g_max, cubes[1])
            b_max = max(b_max, cubes[2])

        ans += r_max * g_max * b_max

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
