
def sliding_window(s, t):
    if s == t:
        return s
    import collections
    numt = collections.Counter(t)

    left, right = 0, 0
    num = {}
    ahead = float('inf')
    result = ""
    
    def check(m1, m2):
        if len(m1) != len(m2):
            return False
        for k, v in (m2.items()):
            if m1.get(k, 0) < v:
                return False
        return True
    
    while right < len(s):
        if s[right] in t:
            num.setdefault(s[right], 0)
            num[s[right]] += 1
        if check(num, numt):
            while left <= right and s[left] not in t:
                left += 1
            if (right - left) < ahead:
                result = s[left:right+1]
                ahead = right-left
            
            left += 1
            right = left
            num = {}

        else:
            right += 1

    return result

print("--", sliding_window("ADOBECODEBANC", "ABC"))
