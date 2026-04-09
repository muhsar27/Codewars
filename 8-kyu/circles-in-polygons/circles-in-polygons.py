import math 
​
def circle_diameter(sides, side_length): 
    
    d = side_length/math.tan(math.pi/sides)
        
    return d 
    