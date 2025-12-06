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
        for line in f:
            springs, c = line.split()
            counts = tuple([int(x) for x in c.split(",")])
            ans += find(springs, counts)

    return ans


cache = {}


def find(springs: str, counts: tuple[int, ...]) -> int:
    """
    functools.lru_cache works too
    """
    if springs == "":
        if len(counts) == 0:
            return 1
        return 0

    if len(counts) == 0:
        if "#" in springs:
            return 0
        return 1

    k = (springs, counts)
    if k in cache:
        return cache[k]

    c = 0
    if springs[0] == "." or springs[0] == "?":
        # treat first pos as non-broken (? = non-broken);
        c += find(springs[1:], counts)
    if springs[0] == "#" or springs[0] == "?":
        # treat first pos as broken
        if (
            counts[0] <= len(springs)
            and "." not in springs[: counts[0]]
            and (counts[0] == len(springs) or springs[counts[0]] != "#")
        ):
            c += find(springs[counts[0] + 1 :], counts[1:])

    cache[k] = c

    return c


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    ans = 0
    with open(file) as f:
        for line in f:
            springs_folded, c_folded = line.split()
            springs = "?".join([springs_folded] * 5)
            counts = tuple([int(x) for x in c_folded.split(",")]) * 5
            ans += find(springs, counts)

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
