import random


class Grammar:
    def __init__(self, Vn, Vt, P, S):
        self.Vn = Vn
        self.Vt = Vt
        self.P = P
        self.S = S

    def initial_symbol(self):
        return self.S

    def productions(self):
        return self.P

# Generates a word by given grammar
    def generateString(self) -> str:
        word = self.S
        x = True
        while x:
            # checking if there are non-terminal states in the word
            # I did the negation of that statement in order to avoid nesting
            if not any(n in word for n in self.Vn):
                # exiting the while loop if the word is complete
                x = False
                break
            # for every non-terminal state
            for n in self.Vn:
                if n in word:
                    # choosing a transition for corresponding non-terminal state
                    transition = self.P.get(n)
                    # getting a random transition
                    transition = random.choice(transition)
                    # replacing Vn with the transition
                    word = word.replace(n, transition)
        return word



