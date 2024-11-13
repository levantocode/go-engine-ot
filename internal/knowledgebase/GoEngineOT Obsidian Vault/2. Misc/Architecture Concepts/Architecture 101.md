
Every program has **Some Organization**, even if it’s just:
> “jam the whole thing into main() and see what happens”.

The questions that emerge are:
- What makes for Good Organization?
- How do we tell a Good Architecture from a bad one?

# Change, Flexibility & Abstraction

Ultimately, Architecture is about **Change** (considering the future).
Someone has to be modifying the Codebase.
If no one is ever touching the code, it's design is irrelevant.

A lot of Software Architecture is about making the program more **Flexible**, so it takes less **Time & Effort** to Change. Good Architecture makes a huge difference in **Productivity**.

Making the code more Flexible, ultimately means encoding fewer assumptions into it, which comes in some form or another of Abstraction.

Examples:
- You use interfaces so that your code works with *any* class that implements it instead of just the one that does today.
- You use design patterns like Observers & Messaging to let two parts of the game talk to each other so that tomorrow, it can easily be three or four.

# Iteration & Experimentation

It is virtually impossible to come up with a balanced game design on paper.
It demands **Iteration & Experimentation**. And for that we have Good Architecture.

The faster you can try out ideas and see how they feel, the more you can try and the more likely you are to find something great. Even after you’ve found the right mechanics, you need plenty of time for Tuning.

A poorly designed game collapses to the one winning tactic played over and over until you get bored and quit.

# Complexity & Cognitive Load

Much of Software Architecture is about that learning phase every developer have to go through when reading a new codebase for them. Loading code into neurons is so painfully slow that it pays to find strategies to reduce the volume of it.

Key goal of Software Architecture minimize the amount of knowledge you need to have in Short/Medium Term Memory (in your brain) before you can make progress.

A major element of that is Managing Complexity (usually reducing), and a major element of that is the concept of Coupling & Decoupling. As such, Architectural Concepts tends to flow down into Design Patterns.

# Costs of Architecture

Although Good Architecture sounds great, it doesn’t come free. Good Architecture takes real Effort & Discipline. You have to take great care to both organize the code well and keep it organized throughout the thousands of little changes that make up a development cycle. Many programs start out beautifully and then die a death of a thousand cuts as programmers add “just one tiny little hack” over and over again. 

Much like gardening, it’s not enough to put in new plants. You must also Weed and Prune. Writing well-architected code takes careful thought, and that translates to time. Maintaining a Good Architecture over the life of a project takes a lot of effort. You have to treat your codebase like a good camper does their campsite: always try to leave it a little better than you found it.

The problem with encoding fewer assumptions is that performance is all about assumptions. The practice of Optimization thrives on concrete limitations. For example:
- "Can we safely assume we’ll never have more than 256 enemies? Great, we can pack an ID into a single byte"
- "Are all of the entities going to be the same class? Great, we can make a nice contiguous array of them"
- "Will we only call a method on one concrete type here? Good, we can statically dispatch or inline it"

One compromise is to keep the code flexible until the design settles down and then tear out some of the abstraction later to improve your performance.

Think about and design for performance throughout your development cycle, but put off the low-level, nitty-gritty optimizations that lock assumptions into your code until as late as possible.

# Prototyping (When Not to Architect)

**Prototyping** is a perfectly legitimate programming practice: slapping together code that’s just barely functional enough to answer a design question.

Game design requires a lot of experimentation and exploration. Especially early on, it’s common to write code that you know you’ll throw away.

If you write throwaway code, you must ensure you’re able to throw it away.

If you just want to find out if some gameplay idea plays right at all, architecting it beautifully means burning more time before you actually get it on screen and get some feedback. If it ends up not working, that time spent making the code elegant goes to waste when you delete it.

Make sure the people using the throwaway code understand that even though it kind of looks like it works, it cannot be maintained and must be rewritten.

Move quickly to explore your game’s design space, but don’t go so fast that you leave a mess behind you. You’ll have to live with it, after all.