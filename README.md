# CNC Manager

My uncle has a CNC in his garage. It'd be neat if he could remotely monitor its status securely. Eventually, actions could be automated and remotely controlled this way.

A streamed video feed would also be a great feature.

# Project Structure

- Client (SvelteKit)
- Backend (golang app built with Echo) (API Gateway)
  - bin (untracked binaries)
  - cmd (entry point)
  - controllers (http handlers)
  - models (types/structs)
  - storage (database connections)
    - PostgreSQL (main persistence)
    - Redis (sessions and cache)
- Services (individual functions needing separate runtimes)
  - stream (planned)
  - machine interface (planned)
  - automation (planned)

# Research

For status and stuff, something like Kafka may be great for event-streaming changes and updates.
