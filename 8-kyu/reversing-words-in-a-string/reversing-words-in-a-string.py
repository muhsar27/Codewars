def reverse(st):
    new_str = ""
    
    words = st.split()
    
    for i in range(len(words)-1, -1 , -1):
        if i != 0:
            new_str += words[i] + " " 
        else: 
            new_str += words[i]
        
    return new_str