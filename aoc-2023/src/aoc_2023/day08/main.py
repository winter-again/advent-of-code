import argparse
import math
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
    # NOTE: there's actually 2 sample inputs given; testing against 2nd
    if sample:
        file = parent / "input_smpl.txt"

    nodes = {}
    with open(file) as f:
        moves, lines = f.read().split("\n\n")
        for line in lines.splitlines():
            node, elems = line.split(" = ")
            left, right = (
                elems.strip().replace("(", "").replace(")", "").strip().split(", ")
            )
            nodes[node] = (left, right)

    i = 0
    cur = "AAA"
    ans = 0
    n = len(moves)
    while cur != "ZZZ":
        if moves[i] == "L":
            nxt = nodes[cur][0]
        else:
            nxt = nodes[cur][1]

        cur = nxt
        ans += 1
        i = ans % n

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl_2.txt"

    nodes = {}
    with open(file) as f:
        moves, lines = f.read().split("\n\n")
        for line in lines.splitlines():
            node, elems = line.split(" = ")
            left, right = (
                elems.strip().replace("(", "").replace(")", "").strip().split(", ")
            )
            nodes[node] = (left, right)

    starts = [n for n in nodes if n.endswith("A")]
    n = len(moves)

    steps = []
    for start in starts:
        i = 0
        cur = start
        ans = 0
        n = len(moves)
        while not cur.endswith("Z"):
            if moves[i] == "L":
                nxt = nodes[cur][0]
            else:
                nxt = nodes[cur][1]

            cur = nxt
            ans += 1
            i = ans % n

        steps.append(ans)

    return math.lcm(*steps)


if __name__ == "__main__":
    raise SystemExit(main())
