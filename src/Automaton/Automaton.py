import sys

sys.path.insert(1, "../FLFA-Labs/src/Grammar")
import Grammar


class FiniteAutomaton:
    def __init__(self,
                 states,
                 alphabet,
                 transitions,
                 initial_state,
                 final_sates) -> None:
        self.states = states
        self.alphabet = alphabet
        self.transitions = transitions
        self.initial_state = initial_state
        self.final_states = final_sates

    def RG_to_NFA(self, grammar) -> 'FiniteAutomaton':
        productions = grammar.productions
        states = set(productions.keys()) | {"X"}
        alphabet = set()
        transitions = {}
        initial_state = grammar.initial_symbol()
        final_states = set("X") | \
                       ({initial_state} if "&" in productions[initial_state] else set())
        for non_terminal, prods in productions.items():
            for production in prods:
                if production == "&":
                    continue
                new_transition = "X" if len(production) == 1 else production[1]
                transitions.setdefault(
                    (non_terminal, production[0]), set()
                ).add(new_transition)
                alphabet.add(production[0])
        return FiniteAutomaton(states, alphabet, transitions, initial_state, final_states)

    def check_word(self, word) -> bool:
        current_state = {self.initial_state}
        for symbol in word:
            next_state = set()
            for state in current_state:
                next_state.update(
                    self.transitions.get((state, symbol), set())
                )
            current_state = next_state
        return bool(current_state.intersection(self.final_states))

    def determine_FA(self) -> str:
        transition_count = []
        for k, states in self.transitions.items():
            for state in states:
                if k[1] == '&':
                    return "&-NFA"
                if k[0] in transition_count:
                    return "NFA"
                else:
                    transition_count.append(k[0])
        return "DFA"

    def __closure(self, q) -> list:
        closure = list()
        closure.append(q)
        for element in closure:
            for k, states in self.transitions.items():
                for state in states:
                    if k[0] == element and k[1] == '&':
                        closure.append(state)
        return closure

    def NFA_to_DFA(self):
        # checking if it is already a DFA
        if self.determine_FA() == 'DFA':
            raise ValueError("This is already a DFA")

        # additional functions to avoid nesting
        def add_new_transition():
            nonlocal dfa_state, letter, new_state
            # values in dictionaries are not always traversed in order
            # to avoid duplicates with different order of elements
            # sort new_state before adding it in dfa_states
            new_state.sort()
            if len(new_state) > 0:
                # making a transition if there is a state
                dfa_transitions.setdefault(
                    (tuple(dfa_state), letter), set()
                ).add(tuple(new_state))
                update_final_states()
                # adding the state into dfa_states if it's a new one
                if new_state not in dfa_states:
                    dfa_states.append(new_state)

        def update_final_states():
            nonlocal new_state
            if any(item in new_state for item in self.final_states)\
                    and tuple(new_state) not in dfa_final_states:
                dfa_final_states.append(tuple(new_state))

        def find_new_state():
            nonlocal dfa_state, letter, new_state
            for element in dfa_state:
                for key, states in self.transitions.items():
                    # finding the necessary state in dict key
                    if key[0] == element and letter == key[1]:
                        for state in states:
                            if state not in new_state:
                                new_state = append_closure(state)

        def append_closure(state):
            nonlocal new_state, closures
            for closure in closures:
                if state == closure[0]:
                    new_state = new_state + closure
                    break
            new_state = list(set(new_state))
            return new_state

        # initializing variables of the DFA
        dfa_states = list()
        dfa_transitions = {}
        dfa_final_states = list()
        closures = list()
        # making a list of &-closure
        for key, states in self.transitions.items():
            if key[0] in [item[0] for item in closures]:
                continue
            closures.append(self.__closure(key[0]))
        # dfa_states[0] is the initial state of DFA
        dfa_states.append(closures[0])
        self.initial_state = dfa_states[0]

        for dfa_state in dfa_states:
            for letter in self.alphabet:
                new_state = list()
                find_new_state()
                add_new_transition()

        self.states = dfa_states
        self.transitions = dfa_transitions
        self.final_states = dfa_final_states



