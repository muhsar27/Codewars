def better_than_average(class_points, your_points):
    sum = 0 
    for i in range(0,len(class_points),1):
        sum += class_points[i]
        
    average = sum // len(class_points)
    
    if your_points > average:
        return True
    else:
        return False 
    
    return False