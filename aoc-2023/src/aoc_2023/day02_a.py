import re
from dataclasses import dataclass

RED = 12
GREEN = 13
BLUE = 14


@dataclass
class Draw:
    red: int
    green: int
    blue: int


@dataclass
class Game:
    id: int
    draws: list[Draw]

    def check_draws(self) -> bool:
        valids = []
        for draw in self.draws:
            is_valid = True
            if draw.red > RED:
                is_valid = False
            if draw.green > GREEN:
                is_valid = False
            if draw.blue > BLUE:
                is_valid = False
            valids.append(is_valid)

        if all(valid for valid in valids):
            return True
        else:
            return False


def get_games(inp: str) -> int:
    with open(inp) as f:
        lines = f.readlines()

    games = []
    for line in lines:
        game_id = int(re.search(r"^Game (\d+):", line).group(1))
        draws = line.split(": ")[1]
        draws = [d.strip() for d in draws.split("; ")]

        draws_out = []
        for draw in draws:
            num_red = re.findall(r"(\d+) red", draw)
            num_green = re.findall(r"(\d+) green", draw)
            num_blue = re.findall(r"(\d+) blue", draw)
            nums = tuple(
                int(sub_list[0]) if len(sub_list) == 1 else 0
                for sub_list in (num_red, num_green, num_blue)
            )
            draws_out.append(Draw(red=nums[0], green=nums[1], blue=nums[2]))

        games.append(Game(id=game_id, draws=draws_out))

    # print(games)
    valid_games = [game.id for game in games if game.check_draws() == True]
    return sum(valid_games)


if __name__ == "__main__":
    print(get_games("./inputs/day_02_a_smpl.txt"))
    print(get_games("./inputs/day_02.txt"))
