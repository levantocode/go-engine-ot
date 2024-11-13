
# Speed Trade-offs

There are a few forces shaping the development of most projects:
1. Wanting a nice Architecture so the code is easier to change and understand
2. Wanting a fast Runtime Performance
3. Wanting to Ship today's Features Quickly

All of those are about some kind of speed:
1. Long-term Development Speed
2. Software Execution Speed
3. Short-term Development Speed

Good architecture improves productivity over the long term, but maintaining it means every change requires a little more effort to keep things clean.

The implementation that’s quickest to write is rarely the quickest to run. Optimization takes significant engineering time and once it’s done, it tends to calcify the codebase (highly optimized code is inflexible and very difficult to change).

There’s no simple answer here, just trade-offs.


# Simplicity

If there is any method that eases all 3 constraints, it's Simplicity.
A good solution isn’t an accretion of code, it’s a distillation of it.

"Perfection is achieved, not when there is nothing more to add, but when there is nothing left to take away."

If we can keep things simple, there’s less code overall.
- That means less code to fill one's Cognitive Load in order to change it.
- It often runs fast because there’s simply not as much Overhead and not much code to execute (not always the case though).

**Elegant Solutions**: a small bit of logic that still correctly covers a large space of Use Cases.

Finding that is a bit like Pattern Matching or Puzzle Solving. It takes effort to see through the scattering of example Use Cases to find the hidden order underlying them all. It’s a great feeling when you pull it off.