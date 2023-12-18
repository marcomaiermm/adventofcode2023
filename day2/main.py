import re

total_1, total_2 = 0, 0
for line in open("data").readlines():
    game = int(re.findall("Game (\d+)", line)[0])
    mred, mgreen, mblue = [
        max(map(int, re.findall(f"(\d+) {colour}", line)))
        for colour in ["red", "green", "blue"]
    ]
    if mred < 13 and mgreen < 14 and mblue < 15:
        total_1 += game
    total_2 += mred * mgreen * mblue
print(total_1, total_2)
