def minimumRounds(tasks: list[int]) -> int:
    rounds = 0
    d = dict()
    for task in tasks:
        d[task] = d.get(task, 0) + 1
    for task in d:
        if d[task] == 1:
            return -1
        rounds += (d[task] + 2) // 3
    return rounds
