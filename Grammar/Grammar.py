import random

class Grammar:
    def __init__(self, Vn, Vt, P, S):
        self.Vn = Vn
        self.Vt = Vt
        self.P = P
        self.S = S

    def generateString(self):
        word = self.S
        x = True
        while x:
            if not any(n in word for n in self.Vn):
                x = False
                break
            for n in self.Vn:
                if n in word:
                    transition = self.P.get(n)
                    transition = random.choice(transition)
                    word = word.replace(n, transition)
        return word



