import re

digit_map = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}


def get_calibs(inp):
    with open(inp) as f:
        lines = f.readlines()

    calibs = []
    for line in lines:
        m = re.findall(
            # regex uses "(?=)" to allow capturing of overlapping expressions;
            # e.g. "eightwo" has "eight" and "two" in it that we probably want to capture before filtering
            r"(?=((\d)|one|two|three|four|five|six|seven|eight|nine))",
            line,
        )
        interm_digits = []
        for tup in m:
            for elem in tup:
                if elem != "":
                    interm_digits.append(elem)
        digits = [interm_digits[0], interm_digits[-1]]

        # this seems ugly; maybe there's a cleaner way
        # one alt is a sep converter func that returns either int() of the string or the result of mapping the string digit to the int equivalent
        for idx, d in enumerate(digits):
            for k, v in digit_map.items():
                if d == k:
                    digits[idx] = v

        calibs.append(int(digits[0] + digits[1]))

    return sum(calibs)


if __name__ == "__main__":
    print(get_calibs("./inputs/day01_b_smpl.txt"))
    print(get_calibs("./inputs/day01.txt"))
