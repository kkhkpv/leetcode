def isSubsequence(s: str, t: str) -> bool:
    l1, l2 = len(s), len(t)
    if l1 > l2:
        return False
    for i in range(l1):
        if s[i] not in t:
            return False
        else:
            t = t[t.find(s[i])+1:]
    return True
