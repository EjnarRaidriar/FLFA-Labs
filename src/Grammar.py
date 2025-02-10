import random
import string

import Automaton


class Grammar:
    def __init__(self, non_terminal: list, terminal: list, productions: dict, initial_symbol: chr):
        self._non_terminal = non_terminal
        self._terminal = terminal
        self._productions = productions
        self._initial_symbol = initial_symbol
        # this fixes if dict parameter wan not initialized with values as lists
        for non_terminal, prod in self._productions.items():
            if not isinstance(prod, list):
                self._productions[non_terminal] = [prod]

    def __repr__(self):
        return f"<Grammar:\n\tnon_terminal: {self._non_terminal}\n\tterminal: {self._terminal}\n\tproductions: {self._productions}\n\tinitial_symbol: {self._initial_symbol}>"

    def __str__(self):
        return f"non_terminal: {self._non_terminal}\nterminal: {self._terminal}\nproductions: {self._productions}\ninitial_symbol: {self._initial_symbol}"

    def initial_symbol(self) -> chr:
        return self._initial_symbol

    def productions(self) -> dict:
        return self._productions

# Generates a word by given grammar
    def generateString(self) -> str:
        word = self._initial_symbol
        x = True
        while x:
            # checking if there are non-terminal states in the word
            # I did the negation of that statement in order to avoid nesting
            if not any(n in word for n in self._non_terminal):
                # exiting the while loop if the word is complete
                x = False
                break
            # for every non-terminal state
            for n in self._non_terminal:
                if n in word:
                    # choosing a transition for corresponding non-terminal state
                    transition = self._productions.get(n)
                    # getting a random transition
                    transition = random.choice(transition)
                    # replacing Vn with the transition
                    word = word.replace(n, transition)
        return word

    def define_grammar(self) -> str:
        type_3 = True
        right_regular = False
        type_2 = True
        for non_terminal, strings in self._productions.items():
            # type 3 and 2 have single non_terminal on the left side
            if len(non_terminal) > 1:
                type_3 = False
                type_2 = False
            for string in strings:
                # it is type 0 if left is longer than right
                if len(non_terminal) > len(string):
                    return 'Type_0'
                # type 1 can not have & unless it derives from starting symbol
                if non_terminal != self._initial_symbol:
                    if not (type_3 or type_2) and string == '&':
                        return 'Type_0'
                # making checks from type 2 to see if it is not type 3
                # skipping this check if right side is made from only terminals
                if type_2 and not string.islower():
                    if string.isupper():
                        type_3 = False
                    # if there are two non-terminals on right side
                    # then it can't be of type 3
                    elif sum(1 for c in string if c.isupper()) > 1:
                        type_3 = False
                    # if it is both left and right regular then it can't be type 3
                    elif string[0].islower() and string[-1].islower():
                        type_3 = False
                if type_3:
                    if string[-1].isupper():
                        right_regular = True
                    if right_regular and string[0].isupper():
                        type_3 = False
        if type_3 and right_regular:
            return 'Right_Regular'
        elif type_3:
            return 'Left_Regular'
        if type_2:
            return 'Type_2'
        return 'Type_1'

    @staticmethod
    def FA_to_RG(fa: 'Automaton') -> 'Grammar':
        # store changing from states into non-terminal through dictionary
        name_changes = dict()

        # loop through fa transitions and match each state with a non-terminal
        alphabet_upper = list(string.ascii_uppercase)
        letter_index = 0
        for key, states in fa.transitions.items():
            if key[0] not in name_changes.values():
                name_changes.update({alphabet_upper[letter_index]: key[0]})
                letter_index += 1
                # by looping over the values we will cover dead states too
                # dead states usually are not keys
                for state in states:
                    if state not in name_changes.values():
                        name_changes.update({alphabet_upper[letter_index]: state})

        initial_symbol = ''
        for prod, state in name_changes.items():
            if state == fa.initial_state:
                initial_symbol = prod
                break

        productions = dict()
        # loop over keys in fa transitions and name matching
        for key, states in fa.transitions.items():
            for prod, fa_state in name_changes.items():
                # now we have transition start state matched with its new name
                if key[0] == fa_state:
                    for state in states:
                        for rg_letter, fa_transition in name_changes.items():
                            # now match state with its new name
                            if state == fa_transition:
                                # if it's not epsilon transition append terminal with non-terminal
                                if key[1] != '&':
                                    productions.setdefault(
                                        prod, set()
                                    ).add(str(key[1]) + str(rg_letter))
                                else:
                                    # for epsilon transition use only terminal symbol
                                    productions.setdefault(
                                        prod, set()
                                    ).add(str(rg_letter))
        return Grammar(list(name_changes.keys()), fa.alphabet, productions, initial_symbol)
