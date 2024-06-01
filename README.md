# GitArchived Service (Under Development)

GitArchived service is a repository that hosts the platform's new backend. This backend will replace the current one.

## What's New?

The new backend offers several advantages, including enhanced stability and easier scalability, thanks to RabbitMQ.

### Key Changes:

1. **Lister**:
   - Previously, the updater was responsible for directly checking repositories for updates. Now, this task has been assigned to a new component called the `lister`. The lister is responsible for identifying repositories that need to be updated and placing them in the appropriate RabbitMQ queues for further processing.

2. **Introduction of the Deleter**:
   - We introduced the `deleter`, a worker dedicated to performing precise tests on repositories considered unreachable by the lister. Previously, this function was integrated into the updater.

3. **Docker Compose Integration**:
   - We have integrated Docker Compose to facilitate fast and secure deployment. Previously, Docker was used alone without Compose.

These improvements ensure a more robust and scalable system.
