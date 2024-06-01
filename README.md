# GitArchived Service

GitArchived service is a repository that hosts the platform's new backend. This backend will replace the current one.

## What's New?

The new backend offers several advantages, including enhanced stability and easier scalability, thanks to RabbitMQ.

### Key Changes:

1. **Cronjob to Lister**:
   - We have moved the cronjob responsible for checking repositories to a new component called `lister`. The lister places repositories in the appropriate queue for further processing.

2. **Introduction of the Deleter**:
   - We introduced the `deleter`, a worker dedicated to performing precise tests on repositories considered unreachable by the lister. Previously, this function was integrated into the updater.

3. **Docker Compose Integration**:
   - We have integrated Docker Compose to facilitate fast and secure deployment. Previously, Docker was used alone without Compose.

These improvements ensure a more robust and scalable system.
