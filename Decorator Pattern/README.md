**Problem Definition:**

Imagine you are developing a library for a logging system in a software application. This logging system needs to log messages to different outputs such as the console, a file, or a remote server. Additionally, the system needs to support various functionalities like formatting the log messages, filtering messages based on severity, and timestamping each log message.

Over time, the requirements evolve:
1. The logs need to be encrypted before being written to a file.
2. You need to add contextual information (e.g., user ID, session ID) to each log entry.
3. Certain log entries need to be duplicated and sent to multiple outputs simultaneously.
4. You want the flexibility to enable or disable these features dynamically at runtime.

**Hint:**
- Think about how you would implement this system without making the logger classes overly complex or tightly coupled to these features.
- Consider how you can add these functionalities independently and in a combinable way without modifying the core logging class.

This problem can be effectively addressed by using the **decorator pattern**.