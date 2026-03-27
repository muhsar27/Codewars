def solution(number):
    count = 0 
    for i in range(number):
        if i % 5 == 0 or i % 3 == 0:
            count += i
    return count 