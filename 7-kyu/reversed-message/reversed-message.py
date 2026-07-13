def reverse_message(text):
    reverse = text[::-1]
    words = reverse.split()
    
    for i in range(len(words)-1,-1 , -1):
        words[i] = words[i].lower()
        
        words[i] = words[i].capitalize()
        
    return " ".join(words)