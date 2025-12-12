import argparse
import itertools
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


class Tile(NamedTuple):
    x: int
    y: int


def part_1(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    red_tiles: list[Tile] = []
    with open(file) as f:
        for line in f:
            x, y = map(int, line.strip().split(","))
            red_tiles.append(Tile(x=x, y=y))

    ans = 0
    for t1, t2 in itertools.combinations(red_tiles, 2):
        di = abs(t2.x - t1.x) + 1
        dj = abs(t2.y - t1.y) + 1
        area = di * dj
        ans = max(area, ans)

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    red_tiles: list[Tile] = []
    with open(file) as f:
        for line in f:
            x, y = map(int, line.strip().split(","))
            red_tiles.append(Tile(x=x, y=y))

    rects: list[tuple[Tile, Tile, int]] = []
    for t1, t2 in itertools.combinations(red_tiles, 2):
        di = abs(t2.x - t1.x) + 1
        dj = abs(t2.y - t1.y) + 1
        area = di * dj
        rects.append((t1, t2, area))

    edges: set[tuple[Tile, Tile]] = set()
    for t1, t2 in zip(red_tiles, red_tiles[1:] + red_tiles[:1]):
        edges.add((t1, t2))

    ans = 0
    for t1, t2, area in sorted(rects, key=lambda r: r[2], reverse=True):
        if rect_inside(t1, t2, edges):
            ans = area
            break

    return ans


def tile_inside(t: Tile, edges: set[tuple[Tile, Tile]]) -> bool:
    intersections = 0
    for e1, e2 in edges:
        if (t.x == e1.x == e2.x and min(e1.y, e2.y) <= t.y <= max(e1.y, e2.y)) or (
            t.y == e1.y == e2.y and min(e1.x, e2.x) <= t.x <= max(e1.x, e2.x)
        ):
            return True

        if ((t.y < e1.y) != (t.y < e2.y)) and (
            t.x < (e2.x - e1.x) * (t.y - e1.y) / (e2.y - e1.y) + e1.x
        ):
            intersections += 1

    return intersections % 2 != 0


def rect_inside(t1: Tile, t2: Tile, edges: set[tuple[Tile, Tile]]) -> bool:
    x_min = min(t1.x, t2.x)
    x_max = max(t1.x, t2.x)
    y_min = min(t1.y, t2.y)
    y_max = max(t1.y, t2.y)

    # easy case of corners
    for x, y in itertools.product([x_min, x_max], [y_min, y_max]):
        if not tile_inside(Tile(x=x, y=y), edges):
            return False

    # sides of rectangle cross any edge
    for e1, e2 in edges:
        if rect_crosses(x_min, y_min, x_max, y_max, e1, e2):
            return False

    return True


def rect_crosses(x_min, y_min, x_max, y_max, e1, e2) -> bool:
    # horizontal edge must be crossed by vertical rect edge
    if e1.y == e2.y:
        # rect edge starts below poly edge and ends above it
        if y_min < e1.y < y_max:
            # check if at least one of the vertical rect edges crosses
            if min(e1.x, e2.x) < x_min < max(e1.x, e2.x) or min(
                e1.x, e2.x
            ) < x_max < max(e1.x, e2.x):
                return True
    # vertical edge must be crossed by horizontal rect edge
    else:
        # rect edge starts to left of poly edge and ends to the right
        if x_min < e1.x < x_max:
            # check if at least one of the horizontal rect edges crosses
            if min(e1.y, e2.y) < y_min < max(e1.y, e2.y) or min(
                e1.y, e2.y
            ) < y_max < max(e1.y, e2.y):
                return True

    return False


if __name__ == "__main__":
    raise SystemExit(main())
