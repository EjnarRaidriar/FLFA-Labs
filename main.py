from Grammar import Grammar
#from Automaton import Automaton

G = Grammar.Grammar( ['S', 'B', 'D'],
            ['a', 'b', 'c', 'd'],
            {
                'S': ['aS', 'bB'],
                'B': ['cB', 'd', 'aD'],
                'D': ['aB', 'b']
            },
            'S')

print(G.generateString())
