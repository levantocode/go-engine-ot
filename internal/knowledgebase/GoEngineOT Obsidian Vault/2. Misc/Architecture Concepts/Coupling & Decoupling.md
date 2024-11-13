# Coupling

2+ Pieces of Code are Coupled when you can’t understand one piece without understanding the other. The level of Abstraction that this concept starts to matter most are Classes & Components.


# Decoupling

Allows to reason about a piece of code independently from another.
Change to one piece of code doesn’t necessitate a change to another.

This improves Extensibility.

Design Concepts related to Decoupling are: **Abstraction & Modularity**.

You can think about which parts of the program should be Decoupled and introduce abstractions at those points. This is where it starts to get tricky. Whenever you add a Layer of Abstraction (or a place where extensibility is supported), you’re speculating that you will need that flexibility later. You’re adding code to your game that takes time to develop, debug, and maintain.

All decoupling made means you have less code to understand before you can extend it, but the Layers of Abstraction themselves end up filling your mind.

Predicting the future is hard, and when that modularity doesn’t end up being helpful, it quickly becomes actively harmful. Some people coined the term **YAGNI** (You Aren't Gonna Need It) as a mantra to use to fight the urge to speculate about future extensibility. It’s easy to get so wrapped up in the code itself that you lose sight of the fact that you’re trying to ship the project.

Decoupling & Abstraction make evolving your program faster and easier, but don’t waste time doing them unless you’re confident the code in question needs that flexibility.