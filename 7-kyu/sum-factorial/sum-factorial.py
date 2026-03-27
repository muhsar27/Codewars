def sum_factorial(lst):
    factorial = 1
    sum = 0 
    for i in range(len(lst)):
        for j in range(1,lst[i]+1,1):
            factorial = factorial * j
        sum += factorial 
        factorial = 1               
    return sum