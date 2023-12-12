import re
from dataclasses import dataclass


@dataclass
class Card:
    winning: list[int]
    draw: list[int]

    def get_score(self) -> int:
        winners = 0
        for d in self.draw:
            if d in self.winning:
                winners += 1
        if winners > 0:
            return 2 ** (winners - 1)
        else:
            return 0


def parse_cards(inp: str) -> list[Card]:
    cards = []
    with open(inp) as f:
        for line in f:
            card = re.sub(r"Card\s+\d+:", "", line).strip().split(" | ")
            win_str = list(filter(lambda x: x != "", card[0].split(" ")))
            draw_str = list(filter(lambda x: x != "", card[1].split(" ")))
            win_num = [int(w) for w in win_str]
            draw_num = [int(d) for d in draw_str]
            cards.append(Card(win_num, draw_num))
    return cards


def main(inp: str) -> int:
    cards = parse_cards(inp)
    scores = [card.get_score() for card in cards]
    return sum(scores)


if __name__ == "__main__":
    print(main("./inputs/day04_a_smpl.txt"))
    print(main("./inputs/day04.txt"))
