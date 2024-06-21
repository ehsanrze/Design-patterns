**Problem Definition:**

Imagine you are developing a real-time monitoring system for a fleet of delivery vehicles. Each vehicle sends updates about its current location, speed, and status at regular intervals. These updates need to be processed and displayed in real-time on a dashboard for dispatchers to efficiently manage and track the fleet.

Real-Time Updates: Vehicles send updates continuously and asynchronously.
Multiple Observers: The dashboard needs to display the updated information in various formats (maps, tables, statistics) simultaneously.
Decoupling of Components: The vehicle update handling logic (subject) should be decoupled from the dashboard display logic (observers) to ensure flexibility and maintainability of the system.
Efficient Updates: Observers should only receive updates when relevant (e.g., based on vehicle ID or geographical area).


**Hint:**
* Subject (Vehicle Update Source): Implement a subject interface or abstract class (VehicleUpdateSubject) that allows vehicles to register as observers and notifies them of any updates.
* Observers (Dashboard Components): Define observer interfaces (DashboardObserver) that the dashboard components implement to receive updates. Each observer can decide how to process and display the received updates.
* Registration and Notification: Vehicles (subjects) register themselves with the subject manager, and each dashboard component (observer) subscribes to updates from relevant vehicles or categories.
* Event Filtering: Optionally implement filtering mechanisms within the subject to notify only those observers interested in specific types of updates (e.g., location changes, status updates).
* Dynamic Updates: Allow observers to subscribe and unsubscribe dynamically during runtime to adapt to changing requirements or system configurations.

* By applying the Observer Method Pattern in this scenario, you achieve a flexible and efficient solution where vehicle updates are processed and displayed in real-time on the dashboard without tightly coupling the vehicle update handling logic with the dashboard display logic. This enhances scalability, maintainability, and adaptability of your fleet monitoring system.
