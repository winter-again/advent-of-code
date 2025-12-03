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
        data = f.read().replace("\n", "").split(",")

    pairs = []
    for d in data:
        a, b = d.split("-", maxsplit=1)
        pairs.append((int(a), int(b)))

    for a, b in pairs:
        for i in range(a, b + 1):
            id = str(i)
            n = len(id)
            if n % 2 != 0:
                continue

            mid = n // 2
            if id[:mid] == id[mid:]:
                ans += i

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    with open(file) as f:
        data = f.read().replace("\n", "").split(",")

    pairs = []
    for d in data:
        a, b = d.split("-", maxsplit=1)
        pairs.append((int(a), int(b)))

    for a, b in pairs:
        for i in range(a, b + 1):
            id = str(i)
            n = len(id)
            if n < 2:
                continue

            mid = n // 2
            for chunk_size in range(1, mid + 1):
                if n % chunk_size == 0 and id[:chunk_size] * (n // chunk_size) == id:
                    ans += i
                    break

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
