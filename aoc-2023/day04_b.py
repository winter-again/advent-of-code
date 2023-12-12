import re
from dataclasses import dataclass


@dataclass
class Card:
    id: int
    winners: list[int]
    draw: list[int]
    wins: list[int]

    def is_winning(self):
        if len(self.wins) != 0:
            return True
        return False


def parse_cards(inp: str) -> list[Card]:
    cards = []
    with open(inp) as f:
        # parsing looks ugly and long
        for line in f:
            # can just get the id from this split instead of doing the regex
            card_dat = line.split(": ")
            card_id = int(re.search(r"\d+$", card_dat[0]).group(0))
            card_nums = card_dat[1].strip().split(" | ")

            win_str = list(filter(lambda x: x != "", card_nums[0].split(" ")))
            draw_str = list(filter(lambda x: x != "", card_nums[1].split(" ")))
            win_num = [int(w) for w in win_str]
            draw_num = [int(d) for d in draw_str]
            tot_wins = len([n for n in draw_num if n in win_num])
            win_matches = list(range(card_id + 1, card_id + tot_wins + 1))
            cards.append(Card(card_id, win_num, draw_num, win_matches))
    return cards


# prob feels like recursion but idk
def main(inp: str) -> int:
    cards = parse_cards(inp)
    cards_collec = [c.id for c in cards]
    card_dict = {card.id: card.wins for card in cards}
    # this seems weird but it works...
    # maybe there's a better way keeping track of two sep lists
    for i in cards_collec:
        if len(card_dict[i]) == 0:
            pass
        else:
            cards_collec.extend(card_dict[i])
    return len(cards_collec)


if __name__ == "__main__":
    print(main("./inputs/day04_a_smpl.txt"))
    print(main("./inputs/day04.txt"))
