import re
from dataclasses import dataclass


@dataclass
class Map:
    src: str
    dest: str
    mappers: list[tuple[int, int, int]]

    def conv_seed(self, seed: int) -> int:
        for mapper in self.mappers:
            if mapper[1] <= seed <= mapper[1] + mapper[2] - 1:
                return mapper[0] + (seed - mapper[1])
        return seed


def parse_maps(inp: str) -> tuple[list[int], list[Map]]:
    with open(inp) as f:
        lines = f.readlines()

    seeds = [int(s) for s in lines[0].split(": ")[1].strip().split(" ")]
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
    seeds, maps = parse_maps(inp)
    locs = []
    for seed in seeds:
        for map in maps:
            seed = map.conv_seed(seed)
        locs.append(seed)
    return min(locs)


if __name__ == "__main__":
    print(main("./inputs/day05_a_smpl.txt"))
    print(main("./inputs/day05.txt"))
