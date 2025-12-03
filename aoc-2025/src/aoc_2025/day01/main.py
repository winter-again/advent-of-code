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
    cur = 50
    dial_min = 0
    dial_max = 99
    width = dial_max - dial_min + 1
    with open(file) as f:
        for line in f:
            dir = line[0]
            turns = int(line[1:])
            if dir == "L":
                turns = -turns

            prev = cur
            cur = (cur + turns) % width
            print(f"{prev=} + {turns=} -> {cur=} ({cur + turns})")
            if cur == 0:
                ans += 1

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    cur = 50
    dial_min = 0
    dial_max = 99
    width = dial_max - dial_min + 1
    with open(file) as f:
        for line in f:
            dir = line[0]
            turns = int(line[1:])

            if dir == "L":
                turns = -turns
                wraps = turns // -width
                rem = turns % -width
                ans += wraps

                print(rem)
                if cur != 0 and cur + rem <= dial_min:
                    # remainder either lands dial at 0 or enough for 1 more wrap but not a full cycle
                    # ensure not currently at 0 because the turns left would wrap but not guaranteed
                    # to reach/cross 0
                    ans += 1
            else:
                wraps = turns // width
                rem = turns % width
                ans += wraps

                print(rem)
                if cur + rem > dial_max:
                    # no need to check cur != 0 because sufficiently large move to right will result in a wrap
                    # that crosses 0, which is satisfied by this cond
                    ans += 1

            cur = (cur + turns) % width

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
