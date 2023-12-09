import re


def get_calibs(inp: str) -> int:
    calibs = []
    with open(inp) as f:
        lines = f.readlines()

    for line in lines:
        # could've avoided the regex with
        # m = [int(d) for d in line if d.isdigit()]
        m = re.findall(r"\d", line)
        num = m[0] + m[-1]
        calibs.append(int(num))

    return sum(calibs)


if __name__ == "__main__":
    print(get_calibs("./inputs/day01_a_smpl.txt"))
    print(get_calibs("./inputs/day01.txt"))
