from collections import Counter
import re
import json

def word_count(text):
    words = re.findall(r"\S+", text)
    
    counter = Counter(words)
    
    word_dict = dict(counter)
    
    return word_dict

with open("count_words.txt", "r", encoding="utf-8") as f:
    content = f.read()

result = word_count(content)

print(dict(list(result.items())[:20]))

with open("word_count.json", "w", encoding="utf-8") as f:
    json.dump(result, f, ensure_ascii=False, indent=2)
