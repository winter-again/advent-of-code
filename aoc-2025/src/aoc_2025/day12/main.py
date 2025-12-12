import argparse
from collections.abc import Sequence
from dataclasses import dataclass
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
    elif args.part == 1:
        print(f"Part 1 (sample): {part_1(True)}")
        if not args.sample_only:
            print(f"Part 1: {part_1()}")
    elif args.part == 2:
        print("No part 2")

    return 0


@dataclass
class Present:
    id: int
    shape: list[list[int]]


@dataclass
class Region:
    dims: tuple[int, ...]
    counts: list[int]


def part_1(sample: bool = False) -> int:
    """
    Troll solution
    """
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        data = f.read().strip().split("\n\n")

    pres = data[:-1]
    presents: list[Present] = []
    for p in pres:
        id, s = p.split(":\n")
        shape = [[1 if c == "#" else 0 for c in row] for row in s.splitlines()]
        presents.append(Present(id=int(id), shape=shape))

    regs = data[-1].splitlines()
    regions: list[Region] = []
    for reg in regs:
        d, c = reg.split(": ")
        dims = tuple(map(int, d.split("x")))
        counts = list(map(int, c.split(" ")))
        regions.append(Region(dims=dims, counts=counts))

    skip = 0
    poss = 0
    for i, region in enumerate(regions):
        area_avail = region.dims[0] * region.dims[1]
        area_need = 0
        for id, ct in enumerate(region.counts):
            area_need += ct * sum(map(sum, presents[id].shape))

        if area_need > area_avail:
            skip += 1
            continue

        poss += 1

    return poss


if __name__ == "__main__":
    raise SystemExit(main())
