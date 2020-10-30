from copy import deepcopy

result = []

def backtrack(path, select):
    if not select:
        result.append(deepcopy(path))
        return
    for se in select:
        next_select = deepcopy(select)
        next_select.remove(se)
        path.append(se)
        backtrack(path, next_select)
        path.remove(se)
    
def permetu(nums):
    select = nums
    path = []
    backtrack(path, select)
    print(result)

permetu([1,2,3])
