import re
from dataclasses import dataclass


@dataclass
class Map:
    src: str
    dest: str
    mappers: list[tuple[int, int, int]]

    # inputs: (start, len)
    # (79, 14) = (79-92)
    # (55, 13) = (55-67)
    def conv_range(self, to_conv: list[tuple[int, int]]) -> list[tuple[int, int]]:
        converted = []
        while len(to_conv) > 0:
            ilb, iub = to_conv.pop()
            # print(f"Input: {ilb}, {iub}")
            inp_len = iub - ilb + 1

            for mapper in self.mappers:
                # src (98, 2) = (98-99) -> dest (50-51)
                # src (50, 48) = (50-97) -> dest (52-99)
                # (55, 67) becomes (57, 69)
                # (79, 92) becomes (81, 94)
                slb = mapper[1]
                sub = mapper[1] + mapper[2] - 1
                dlb = mapper[0]

                olb = max(ilb, slb)
                oub = min(iub, sub)
                if olb < oub:
                    converted.append((dlb + (olb - slb), dlb + (oub - slb)))
                    if ilb < olb:
                        to_conv.append((ilb, olb - 1))
                        # to_conv.append((ilb, olb))
                    if iub > oub:
                        to_conv.append((oub + 1, iub))
                        # to_conv.append((oub, iub))
                    break
            else:
                # if mapping isn't successful then we just preserve the orig values
                converted.append((ilb, iub))
        return converted


def parse_almanac(inp: str) -> tuple[list[tuple[int, int]], list[Map]]:
    with open(inp) as f:
        lines = f.readlines()

    seed_mkrs = [int(s) for s in lines[0].split(": ")[1].strip().split(" ")]
    seeds = []
    for idx, mkr in enumerate(seed_mkrs):
        if idx % 2 == 0:
            seeds.append((mkr, mkr + seed_mkrs[idx + 1] - 1))
    src = ""
    dest = ""
    ranges = []
    maps = []
    # there has to be a better way of looping through the lines by chunks
    for idx, line in enumerate(lines[2:]):
        if re.search("map", line):
            src, dest = (
                line.strip().replace("-to-", " ").replace(" map:", "").split(" ")
            )
        elif re.search(r"\d+", line):
            elems = line.strip().split(" ")
            rng = tuple(int(e) for e in elems)
            ranges.append(rng)
        if not line.strip() or idx == sum(1 for _ in lines[2:]) - 1:
            maps.append(Map(src, dest, ranges))
            ranges = []
    return seeds, maps


def main(inp: str) -> int:
    seeds, maps = parse_almanac(inp)
    # print(f"Init seeds: {seeds}")
    for map in maps:
        seeds = map.conv_range(seeds)
    print(seeds)
    locs = [s[0] for s in seeds]
    return min(locs)


if __name__ == "__main__":
    print(main("./inputs/day05_a_smpl.txt"))
    print(main("./inputs/day05.txt"))
