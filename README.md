# 100ms Server API

This is a Golang implementation of the 100ms Server v2 REST API.

# Getting Started

# Contributing Guide

General

- [x] Generating a management token
- [x] Create a token for joining a room

[Rooms](https://www.100ms.live/docs/server-side/v2/api-reference/Rooms/overview)

- [x] Create a new room
- [x] Get details of a single room
- [x] Get details of a list of rooms in your account
- [x] Update the details of a room
- [x] Enable and Disable a room

[Room Codes](https://www.100ms.live/docs/server-side/v2/api-reference/room-codes/room-code-overview)

- [x] Create a Room Code for every Role in the Room at once
- [x] Create a Room Code for a specific Role in a Room
- [x] Get Room Codes for all Roles in a Room
- [x] Update the current state for a given Room Code.
- [x] Create the auth token for a given short code

[Active Rooms](https://www.100ms.live/docs/server-side/v2/api-reference/active-rooms/overview)

- [x] Get details of a specific Active Room
- [x] Get details of a specific Peer in an active Room
- [x] List details of the Active Peers in a Room
- [x] Update the details of a connected Peer
- [x] Send Message
- [x] Remove/Disconnect a connected Peer from an Active Room
- [x] End an Active Room

[Recordings](https://www.100ms.live/docs/server-side/v2/api-reference/recordings/overview)

- [x] Start a recording for a room
- [x] Stop all recordings running a room
- [x] Get details of a recording
- [x] Get recording jobs of a workspace
- [x] Stop a recording using its unique identifier
- [x] Get the configuration of a recording

[External Streams](https://www.100ms.live/docs/server-side/v2/api-reference/external-streams/overview)

- [ ] Start an external stream for a room
- [ ] Stop all external streams running in the room
- [ ] Get the details of an external stream
- [ ] Get details of all external streams of a workspace
- [ ] Stop an external stream using its unique identifier

[Live Streams](https://www.100ms.live/docs/server-side/v2/api-reference/live-streams/overview)

- [ ] Start a livestream for room
- [ ] Stop all livestreams in a room
- [ ] Get details of a livestream
- [ ] Get details of all livestreams of a workspace
- [ ] Stop a specific live stream by its unique identifier
- [ ] Send timed metadata for a running live stream
- [ ] Pause recording of a running livestream
- [ ] Resume a paused livestream recording

[Recording Assets](https://www.100ms.live/docs/server-side/v2/api-reference/recording-assets/overview)

- [ ] Get details of a Recording Asset
- [ ] Generate a short-lived pre-signed URL for a recording asset
- [ ] Get details of all recording assets of a workspace

[Sessions](https://www.100ms.live/docs/server-side/v2/api-reference/Sessions/object)

- [ ] Get details of a specific session
- [ ] Get details of all sessions in an account

[Policy](https://www.100ms.live/docs/server-side/v2/api-reference/policy/template-object)

- [ ] Create a template.
- [ ] Get the details of a specific template
- [ ] Get the details of all templates in an account
- [ ] Update the details of a template
- [ ] Create or Modify a Role in a template
- [ ] Get details of a specific Role in a template
- [ ] Delete a Role in a template
- [ ] Get settings of a template
- [ ] Update the settings of a template
- [ ] Get the list of destinations in a template
- [ ] Update the destinations in a template

[Analytics](https://www.100ms.live/docs/server-side/v2/api-reference/analytics/overview)

- [ ] Track events of a participant in a room
- [ ] Track recording events in a room

[Stream Key](https://www.100ms.live/docs/server-side/v2/api-reference/stream-key/overview)

- [ ] Create RTMP Stream Key and URL
- [ ] Get the RTMP stream key and URL for a specific room
- [ ] Disable the RTMP stream key for a specific room.
