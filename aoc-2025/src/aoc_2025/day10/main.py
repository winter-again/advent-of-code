import argparse
import itertools
from collections.abc import Sequence
from dataclasses import dataclass
from pathlib import Path

import numpy as np
import scipy

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


@dataclass
class Machine:
    target: set[int]
    buttons: list[set[int]]
    joltages: list[int]


def part_1(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    machines: list[Machine] = []
    with open(file) as f:
        for line in f:
            data = line.strip().split(" ")
            target = {i for i, light in enumerate(data[0][1:-1]) if light == "#"}
            buttons = [set(map(int, b[1:-1].split(","))) for b in data[1:-1]]
            joltage = list(map(int, data[-1][1:-1].split(",")))

            machines.append(Machine(target=target, buttons=buttons, joltages=joltage))

    ans = 0
    for m in machines:
        for n_presses in range(1, len(m.buttons) + 1):
            done = False
            for combo in itertools.combinations(m.buttons, n_presses):
                lights = set()
                for button in combo:
                    lights ^= button

                if lights == m.target:
                    ans += n_presses
                    done = True
                    break

            if done:
                break

    return ans


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    machines: list[Machine] = []
    with open(file) as f:
        for line in f:
            data = line.strip().split(" ")
            target = {i for i, light in enumerate(data[0][1:-1]) if light == "#"}
            buttons = [set(map(int, b[1:-1].split(","))) for b in data[1:-1]]
            joltages = list(map(int, data[-1][1:-1].split(",")))

            machines.append(Machine(target=target, buttons=buttons, joltages=joltages))

    ans = 0
    for m in machines:
        a = []
        for button_set in m.buttons:
            buttons = [0] * len(m.joltages)
            for button in button_set:
                buttons[button] = 1

            a.append(buttons)

        A = np.array(a).T
        b = np.array(m.joltages)
        c = [1] * len(m.buttons)
        res = scipy.optimize.linprog(c, A_eq=A, b_eq=b, integrality=1)
        ans += int(sum(res.x))

    return ans


if __name__ == "__main__":
    raise SystemExit(main())
