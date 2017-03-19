
def merge(a, b):
    combined = []
    while a and b:
        if a < b:
            combined.append(a.pop(0))
        else:
            combined.append(b.pop(0))
    if a:
        combined = combined + a
    if b:
        combined = combined + b
    return combined

if __name__ == "__main__":
    print(merge([1, 3, 4], []))