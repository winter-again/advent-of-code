import argparse
import collections
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
    if sample:
        file = parent / "input_smpl.txt"

    hands = []
    with open(file) as f:
        for line in f:
            hand, bid = line.split()
            hands.append((hand, int(bid)))

    hands.sort(key=lambda h: sort_hands_p1(h))
    ans = 0
    for r, (_, bid) in enumerate(hands, 1):
        ans += r * bid

    return ans


cards_p1 = ["A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"]
strength_p1 = {card: s for s, card in enumerate(reversed(cards_p1), 1)}


def sort_hands_p1(t: tuple[str, int]) -> tuple[int, list[int]]:
    hand = t[0]
    hand_type = hand_type_p1(hand)
    strengths = [strength_p1[card] for card in hand]
    return (hand_type, strengths)


def hand_type_p1(hand: str) -> int:
    d = collections.Counter(hand)

    if 5 in d.values():
        # five of a kind (5)
        return 7
    elif 4 in d.values():
        # four of a kind (4, 1)
        return 6
    elif 3 in d.values():
        # full house (3, 2)
        if 2 in d.values():
            return 5
        # three of a kind (3, 1, 1)
        return 4
    elif len(d) == 3 and 2 in d.values() and 1 in d.values():
        # two pair (2, 2, 1)
        return 3
    elif len(d) == 4:
        # one pair (2, 1, 1, 1)
        return 2
    # high card (1, 1, 1, 1, 1)
    return 1


def part_2(sample: bool = False) -> int:
    file = parent / "input.txt"
    if sample:
        file = parent / "input_smpl.txt"

    hands = []
    with open(file) as f:
        for line in f:
            hand, bid = line.split()
            hands.append((hand, int(bid)))

    hands.sort(key=lambda h: sort_hands_p2(h))
    ans = 0
    for r, (_, bid) in enumerate(hands, 1):
        ans += r * bid

    return ans


def hand_type_p2(hand: str) -> int:
    d = collections.Counter(hand)

    if 5 in d.values():
        # five of a kind (5)
        return 7
    elif 4 in d.values():
        # four of a kind (4, 1)
        if "J" in d:
            return 7
        return 6
    elif 3 in d.values():
        # full house (3, 2)
        if 2 in d.values():
            if "J" in d:
                return 7
            return 5
        # three of a kind (3, 1, 1)
        if 1 in d.values():
            if "J" in d:
                return 6
            return 4
    elif len(d) == 3 and 2 in d.values() and 1 in d.values():
        # two pair (2, 2, 1)
        if "J" in d:
            if d["J"] == 2:
                return 6
            if d["J"] == 1:
                return 5
        return 3
    elif len(d) == 4:
        # one pair (2, 1, 1, 1)
        if "J" in d:
            return 4
        return 2
    # high card (1, 1, 1, 1, 1)
    if "J" in d:
        return 2
    return 1


cards_p2 = ["A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"]
strength_p2 = {card: s for s, card in enumerate(reversed(cards_p2), 1)}


def sort_hands_p2(t: tuple[str, int]) -> tuple[int, list[int]]:
    hand = t[0]
    hand_type = hand_type_p2(hand)
    strengths = [strength_p2[card] for card in hand]
    return (hand_type, strengths)


if __name__ == "__main__":
    raise SystemExit(main())
