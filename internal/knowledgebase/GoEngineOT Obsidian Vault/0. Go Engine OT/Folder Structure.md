
## cmd/

In Go `src` are used by libraries to be imported by other projects.
The `cmd` folder is where the `main` package resides.

## resources/

Contains the Assets of the game, like:
- Sprites (.spr)
- Items & Scenario Attributes (.dat)
- Map Information & Items Positions (.otbm)

## internal/

Encapsulates implementation details that are specific to a project, not intended for public use, so internal APIs are not exposed for example.

Although it's mainly about code, for the moment we are storing Documentation (Knowledge Base) in this folder as well.
