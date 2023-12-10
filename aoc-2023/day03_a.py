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
    symbols: list[int]


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
        sym_matches = re.finditer(r"[^\.\w\s]", line)
        symbols = [match.start() for match in sym_matches]
        map_line = MapLine(idx, parts, symbols)
        map_lines.append(map_line)

    return map_lines, map_len, map_width


def main(inp: str) -> int:
    map_lines, map_len, map_width = parse_map(inp)

    parts = []
    for idx, l in enumerate(map_lines):
        idx_abv = idx - 1 if idx > 0 else None
        idx_blw = idx + 1 if idx < map_len - 1 else None
        # print(idx_abv, idx_blw)
        for part in l.parts:
            bl = part.span[0] - 1 if part.span[0] > 0 else part.span[0]
            br = part.span[1] + 1 if part.span[1] < map_width else part.span[1]

            if any(bl <= val <= br for val in l.symbols):
                parts.append(part.num)

            if idx_abv is not None:
                if any(bl <= val <= br for val in map_lines[idx_abv].symbols):
                    parts.append(part.num)
            if idx_blw is not None:
                if any(bl <= val <= br for val in map_lines[idx_blw].symbols):
                    parts.append(part.num)

    return sum(parts)


if __name__ == "__main__":
    print(main("./inputs/day03_a_smpl.txt"))
    print(main("./inputs/day03.txt"))
