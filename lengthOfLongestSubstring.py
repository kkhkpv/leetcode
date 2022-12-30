def lengthOfLongestSubstring(s: str) -> int:
    l = len(s)
    if l <= 1:
        return l

    left, right, res = 0, 0, 0
    chars = set()

    while right < l:
        if s[right] not in chars:
            chars.add(s[right])
            right += 1
            res = max(res, len(chars))
        else:
            chars.remove(s[left])
            left += 1
    return res
