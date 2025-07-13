import re
from dataclasses import dataclass


@dataclass
class PartNum:
    num: int
    span: tuple[int, int]


@dataclass
class Symbol:
    span: tuple[int, int]


@dataclass
class MapLine:
    line: int
    parts: list[PartNum]
    gears: list[tuple[int, int]]


def parse_map(inp: str) -> tuple[list[MapLine], int, int]:
    with open(inp) as f:
        lines = f.readlines()

    map_len = len(lines)
    map_width = len(lines[0].strip())

    map_lines = []
    for idx, line in enumerate(lines):
        num_matches = re.finditer(r"\d+", line)
        parts = [
            PartNum(int(match.group()), (match.start(), match.end() - 1))
            for match in num_matches
        ]
        sym_matches = re.finditer(r"\*", line)  # only look for asterisks
        symbols = [(match.start(), match.end() - 1) for match in sym_matches]
        map_line = MapLine(idx, parts, symbols)
        map_lines.append(map_line)

    return map_lines, map_len, map_width


def main(inp: str) -> int:
    map_lines, map_len, map_width = parse_map(inp)

    # for l in map_lines:
    #     print(l)
    gear_ratios = []
    for idx, l in enumerate(map_lines):
        idx_abv = idx - 1 if idx > 0 else None
        idx_blw = idx + 1 if idx < map_len - 1 else None

        for gear in l.gears:
            gear_hits = []
            bl = gear[0] - 1 if gear[0] > 0 else gear[0]
            br = gear[1] + 1 if gear[1] < map_width else gear[1]

            parts = [part for part in l.parts if part]
            for part in parts:
                if any(
                    bl <= val <= br for val in range(part.span[0], part.span[1] + 1)
                ):
                    gear_hits.append(part.num)

            parts = [part for part in map_lines[idx_abv].parts if part]
            for part in parts:
                if idx_abv is not None:
                    if any(
                        bl <= val <= br for val in range(part.span[0], part.span[1] + 1)
                    ):
                        gear_hits.append(part.num)

            parts = [part for part in map_lines[idx_blw].parts if part]
            for part in parts:
                if idx_blw is not None:
                    if any(
                        bl <= val <= br for val in range(part.span[0], part.span[1] + 1)
                    ):
                        gear_hits.append(part.num)

            if len(gear_hits) == 2:
                gear_ratios.append(gear_hits[0] * gear_hits[1])

    return sum(gear_ratios)


if __name__ == "__main__":
    print(main("./inputs/day03_a_smpl.txt"))
    print(main("./inputs/day03.txt"))
