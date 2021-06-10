#!/usr/bin/env python3

import collections

def place_normal(m, curr, chains):
    i = (curr + 2) % len(chains)
    if i == 0:
        chains.append(m)
        return len(chains) - 1, chains
    else:
        return i, chains[:i] + [m] + chains[i:]

def play(lm, np):
    players = collections.defaultdict(int)
    curr = 0
    chains = [0]
    for i in range(lm):
        pi = i % np # current player
        m = i + 1
        if m % 23 == 0:
            # keep current marble score
            players[pi] += m

            j = (curr - 7) % len(chains)
            n = chains[j]
            players[pi] += n

            curr = j
            chains = chains[:j] + chains[j+1:]
        else:
            curr, chains = place_normal(m, curr, chains)
    return max(players.values())

# part 1
np = 459
lm = 71790
# print(play(lm, np))

# # part 2
np = 29
lm = 394
result = []
for i in range(20):
    result.append(
        [lm * (i + 1), np, play(lm * (i + 1), np)])

diff = []
for pair in zip(result, result[1:]):
    diff = pair[1][2] - pair[0][2]
    print(diff, diff % 23, diff // 23)
    print(diff, diff % np)
