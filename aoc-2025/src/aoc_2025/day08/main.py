import argparse
import collections
import math
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


class Box(NamedTuple):
    id: int
    x: int
    y: int
    z: int


def dist(p: Box, q: Box) -> float:
    return math.sqrt((p.x - q.x) ** 2 + (p.y - q.y) ** 2 + (p.z - q.z) ** 2)


class UnionFind:
    def __init__(self, n: int):
        self.parent = {}
        self.size = collections.defaultdict(lambda: 1)

    def find(self, p: Box) -> Box:
        if p not in self.parent:
            self.parent[p] = p

        rep = p
        while self.parent[rep] != rep:
            rep = self.parent[rep]

        while p != rep:
            tmp = self.parent[p]
            self.parent[p] = rep
            p = tmp

        return rep

    def union(self, p: Box, q: Box) -> None:
        p_par = self.find(p)
        q_par = self.find(q)

        if p_par == q_par:
            return

        if self.size[p_par] > self.size[q_par]:
            self.parent[q_par] = p_par
            self.size[p_par] += self.size[q_par]
        else:
            self.parent[p_par] = q_par
            self.size[q_par] += self.size[p_par]


def part_1(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    boxes: list[Box] = []
    with open(file) as f:
        for i, line in enumerate(f):
            x, y, z = line.strip().split(",")
            boxes.append(Box(id=i, x=int(x), y=int(y), z=int(z)))

    edges: list[tuple[Box, Box, float]] = []
    for i, p in enumerate(boxes):
        for q in boxes[i:]:
            if p == q:
                continue

            dd = dist(p, q)
            edges.append((p, q, dd))

    edges.sort(key=lambda x: x[2])

    n = len(boxes)
    n_pairs = 1000 if not sample else 10
    uf = UnionFind(n)
    for pair in edges[:n_pairs]:
        p, q, dd = pair
        uf.union(p, q)

    sizes = sorted(list(uf.size.values()), reverse=True)[:3]

    ans = 1
    for s in sizes:
        ans *= s

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    boxes: list[Box] = []
    with open(file) as f:
        for i, line in enumerate(f):
            x, y, z = line.strip().split(",")
            boxes.append(Box(id=i, x=int(x), y=int(y), z=int(z)))

    edges: list[tuple[Box, Box, float]] = []
    for i, p in enumerate(boxes):
        for q in boxes[i:]:
            if p == q:
                continue

            dd = dist(p, q)
            edges.append((p, q, dd))

    edges.sort(key=lambda x: x[2])

    n = len(boxes)
    uf = UnionFind(n)
    px, qx = 0, 0
    for p, q, dd in edges:
        p_par = uf.find(p)
        q_par = uf.find(q)

        if p_par != q_par:
            px, qx = p.x, q.x
            uf.union(p, q)

    ans = px * qx
    return ans


if __name__ == "__main__":
    raise SystemExit(main())
