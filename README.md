Title: Refactor and Improve Error Handling in Multiverse API Service
Issue Description:
After reviewing the code in our api package, several areas could benefit from refactoring and improvements, particularly for maintainability and robustness. Here are some points for discussion:

1. Error Handling
The current error handling in functions like getUniverse and idHttp is basic and could be enhanced for better debuggability.
2. Code Comments
Although some functions are self-explanatory, adding comments that explain the purpose and intricacies of our logic would aid future development and onboarding of new developers.
3. Function Naming
Method names like idHttp and getUniverse could be more descriptive. For instance, idHttp could be renamed to something like setupUniverseIdRoute.
4. Code Reusability
The logic to check if a universe exists appears multiple times. This could be abstracted into its own function.
5. Context Timeout
The BigBang function has a hardcoded context timeout of 10 seconds. It might be better to make this configurable.
6. Dependency Management
There are several dependencies in the code. We should ensure they are well-managed, and possibly look into simplifying the dependency graph.
Proposed Changes:
Refactor error handling to use a standardized approach across all methods.
Add more detailed comments to the code.
Rename methods to more descriptive names.
Extract repeated logic into reusable functions.
Make context timeout in BigBang configurable through an environment variable or configuration file.
Evaluate and potentially refactor dependencies.

//////////////////////////////////////////
Acceptance Criteria:
 Standardized error handling implemented.
 Comments added to the codebase.
 Method names updated.
 Repeated logic extracted into functions.
 Timeout setting in BigBang is configurable.
 Dependencies evaluated and refactored if necessary.
