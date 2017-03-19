


def compress(uncompressed):
    count = 1
    compressed = ""
    if not uncompressed:
        return compressed
    letter = uncompressed[0]
    for nx in uncompressed[1:]:
        if letter == nx:
            count = count + 1
        else:
            compressed += "{}{}".format(letter, count)
            count = 1
        letter = nx
    compressed += "{}{}".format(letter, count)
    return compressed
            
if __name__ == "__main__":
    print(compress("aaabbbccccd"))