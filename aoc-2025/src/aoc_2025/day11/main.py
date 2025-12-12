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

    devs = {}
    with open(file) as f:
        for line in f:
            dev, outs = line.strip().split(": ")
            devs[dev] = outs.split()

    cache = {}
    start = "you"
    return search(start, True, True, devs, cache)


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl_2.txt"

    devs = collections.defaultdict(list)
    with open(file) as f:
        for line in f:
            dev, outs = line.strip().split(": ")
            devs[dev] = outs.split()

    cache = {}
    start = "svr"
    return search(start, False, False, devs, cache)


def search(
    cur: str,
    dac: bool,
    fft: bool,
    devs: dict[str, list[str]],
    cache: dict[tuple[str, bool, bool], int],
) -> int:
    if cur == "out":
        if dac and fft:
            return 1

        return 0

    if (cur, dac, fft) in cache:
        return cache[(cur, dac, fft)]

    paths = 0
    for nxt in devs[cur]:
        paths += search(nxt, dac or nxt == "dac", fft or nxt == "fft", devs, cache)

    cache[(cur, dac, fft)] = paths

    return paths


if __name__ == "__main__":
    raise SystemExit(main())
