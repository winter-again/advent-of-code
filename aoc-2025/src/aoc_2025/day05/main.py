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
        ff, aa = f.read().split("\n\n")

    fresh_lines = ff.split()
    fresh = []
    for fr in fresh_lines:
        lb, ub = fr.split("-")
        fresh.append([int(lb), int(ub)])

    avail = [int(x) for x in aa.split()]

    for ing in avail:
        for lb, ub in fresh:
            if lb <= ing <= ub:
                ans += 1
                break

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    fresh = []
    with open(file) as f:
        for line in f:
            if line.strip() == "":
                break

            _lb, _ub = line.strip().split("-")
            fresh.append([int(_lb), int(_ub)])

    fresh.sort()
    fresh_comp = [fresh[0]]
    for lb, ub in fresh[1:]:
        _, ube = fresh_comp[-1]
        if lb <= ube:
            fresh_comp[-1][1] = max(ub, ube)
        else:
            fresh_comp.append([lb, ub])

    ans = 0
    for lb, ub in fresh_comp:
        ans += ub - lb + 1

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
