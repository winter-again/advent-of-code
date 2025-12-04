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
        for bank in f:
            s = [int(b) for b in list(bank.strip()[::-1])]
            n1 = s.pop()
            max_jolt = 0

            while len(s) > 0:
                n2 = s.pop()
                max_jolt = max(n1 * 10 + n2, max_jolt)
                if n2 > n1:
                    n1 = n2

            ans += max_jolt

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    with open(file) as f:
        for bank in f:
            s = [int(b) for b in list(bank.strip())]
            max_jolt = 0
            for idx in range(11):
                n = max(s[: idx - 11])
                s = s[s.index(n) + 1 :]
                max_jolt += n * 10 ** (11 - idx)

            n = max(s)
            max_jolt += n

            ans += max_jolt

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
