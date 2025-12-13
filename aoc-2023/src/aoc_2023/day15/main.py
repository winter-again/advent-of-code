import argparse
import collections
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


def part_1(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        steps = f.read().strip().split(",")

    ans = 0
    for step in steps:
        ans += hash_algo(step)

    return ans


def hash_algo(step: str) -> int:
    cur = 0
    for c in step:
        cur += ord(c)
        cur *= 17
        cur = cur % 256

    return cur


class Step(NamedTuple):
    label: str
    op: str
    focal: int = -1


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    with open(file) as f:
        s = f.read().strip().split(",")

    steps: list[Step] = []
    for step in s:
        if "=" in step:
            steps.append(Step(label=step[:-2], op=step[-2], focal=int(step[-1])))
        elif "-" in step:
            steps.append(Step(label=step[:-1], op=step[-1]))

    boxes = collections.defaultdict(list)
    lenses = {}
    for step in steps:
        box = hash_algo(step.label)
        lenses[step.label] = step.focal

        if step.op == "-":
            if step.label in boxes[box]:
                boxes[box].remove(step.label)
        else:
            if step.label in boxes[box]:
                continue
            else:
                boxes[box].append(step.label)

    ans = 0
    for i, box in boxes.items():
        for j, label in enumerate(box, 1):
            ans += (i + 1) * j * lenses[label]

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
