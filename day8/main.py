#!/usr/bin/env python3

with open("input.txt") as f:
    ipt = [int(i) for i in
           f.read().strip().split(' ')]

# Node(child_num: int, meta_num: int, children: []node, metas: []int, meta_val: []int)

def parse_node(l): # => (node, remaining)
    node = (l[0], l[1], [], [], [])

    # parse children
    remaining = l[2:]
    for i in range(node[0]):
        result = parse_node(remaining)
        node[2].append(result[0])
        remaining = result[1]

    # parse metadata
    for i in range(node[1]):
        node[3].append(remaining[0])
        remaining = remaining[1:]

    return node, remaining


def visit(node):
    for i in range(node[0]):
        yield from visit(node[2][i])
    yield node

result = parse_node(ipt)
root = result[0]

# part 1
print(sum(([sum(n[3]) for n in visit(root)])))

# part 2
for n in visit(root):
    if n[0] == 0:
        n[4].append(sum(n[3]))
    else:
        idx = [i - 1 for i in n[3] if i > 0 and i <= n[0]]
        children = [n[2][i] for i in idx]
        n[4].append(sum([child[4][0] for child in children]))

print(root[4][0])
