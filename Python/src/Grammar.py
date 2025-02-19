from __future__ import annotations

import random
import string

import Automaton

class Grammar:
    def __init__(self, non_terminals: set, terminals: set, productions: dict, initial_symbol: str):
        self._non_terminals = non_terminals
        self._terminals = terminals
        self._productions = productions
        self._initial_symbol = initial_symbol
        # this fixes if dict parameter was not initialized with values as set
        for non_terminal, prod in self._productions.items():
            if not isinstance(prod, set):
                self._productions[non_terminal] = {prod}

    def __repr__(self):
        return f"<Grammar:\n\tnon_terminal: {self._non_terminals}\n\tterminal: {self._terminals}\n\tproductions: {self._productions}\n\tinitial_symbol: {self._initial_symbol}>"

    def __str__(self):
        return f"non_terminal: {self._non_terminals}\nterminal: {self._terminals}\nproductions: {self._productions}\ninitial_symbol: {self._initial_symbol}"

    def initial_symbol(self) -> str:
        return self._initial_symbol

    def productions(self) -> dict:
        return self._productions

# Generates a word by given grammar
    def generateString(self) -> str:
        word = self._initial_symbol
        while True:
            # checking if there are non-terminal states in the word
            # I did the negation of that statement in order to avoid nesting
            if not any(n in word for n in self._non_terminals):
                # exiting the while loop if the word is complete
                break
            # for every non-terminal state
            for n in self._non_terminals:
                if n in word:
                    # choosing a transition for corresponding non-terminal state
                    transition = self._productions.get(n)
                    # getting a random transition
                    transition = random.choice(list(transition))
                    # replacing Vn with the transition
                    # print(isinstance(n, str) + isinstance(transition, str))
                    word = word.replace(str(n), str(transition))
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
    def FA_to_RG(fa: Automaton.FiniteAutomaton) -> Grammar:
        initial_symbol = 'S'
        non_terminals = set(initial_symbol)
        terminals = fa.alphabet()
        productions = dict()

        alphabet_upper = list(string.ascii_uppercase)
        alphabet_upper.remove('S')
        name_changes = dict()
        for state in fa.states():
            if state == fa.initial_state():
                name_changes[state] = 'S'
            else:
                name_changes[state] = alphabet_upper[0]
                alphabet_upper.pop(0)

        for transition, next_state in fa.transitions().items():
            non_terminal = name_changes.get(transition[0])
            terminal = transition[1]
            new_non_terminal = name_changes.get(next_state) if next_state not in fa.final_states() else ""
            if new_non_terminal != "":
                non_terminals.add(new_non_terminal)
            productions.setdefault(
                non_terminal, set()
            ).add(terminal + new_non_terminal)

        return Grammar(non_terminals, terminals, productions, initial_symbol)
