import argparse
from collections.abc import Sequence
from pathlib import Path
from typing import NamedTuple

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


class Map(NamedTuple):
    dst: int
    src: int
    len: int


def read_chunk(file):
    chunk = []
    for line in file:
        if line.strip():
            if line[0].isdigit():
                data = [int(d) for d in line.strip().split()]
                chunk.append(Map(data[0], data[1], data[2]))
        else:
            yield chunk
            chunk = []

    if chunk:
        yield chunk


def part_1(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        seeds = f.readline().strip().split(": ")
        seeds = [int(s) for s in seeds[1].split(" ")]

        for chunk in read_chunk(f):
            seeds_mapped = []
            for seed in seeds:
                for dst, src, ln in chunk:
                    if seed >= src and seed < src + ln:
                        offset = seed - src
                        seeds_mapped.append(dst + offset)
                        break
                else:
                    seeds_mapped.append(seed)

            seeds = seeds_mapped

    return min(seeds)


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        seeds = f.readline().strip().split(": ")
        seeds = [int(s) for s in seeds[1].split(" ")]
        seeds = [(s, s + e) for s, e in zip(seeds[::2], seeds[1::2])]

        for chunk in read_chunk(f):
            seeds_mapped = []
            while len(seeds) > 0:
                start, end = seeds.pop()
                for dst, src, ln in chunk:
                    overlap_start = max(src, start)
                    overlap_end = min(src + ln, end)
                    if overlap_start < overlap_end:
                        seeds_mapped.append(
                            (overlap_start - src + dst, overlap_end - src + dst)
                        )
                        if overlap_start > start:
                            seeds.append((start, overlap_start))
                        if end > overlap_end:
                            seeds.append((overlap_end, end))
                        break
                else:
                    seeds_mapped.append((start, end))

            seeds = seeds_mapped

    return min([s[0] for s in seeds])


if __name__ == "__main__":
    raise SystemExit(main())
