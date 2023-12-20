# 100ms Golang API Service

This is a Golang implementation of the 100ms Server v2 REST API.

# Getting Started

## Running Locally

Download the repository

```
git clone git@github.com:Mupati/100ms-golang-api.git
```

Copy the environment variables template and set the values.

```
cd 100ms-golang-api
cp .env.example .env
```

Install packages

```
go mod download
```

Activate the environment variables

```
source .env
```

Start the server

```
go run main.go
```

## Running on Docker

Build docker image

```
docker build --tag hms-api .
```

Run docker container from the image.

We will use the `.env` file as our environment variables file when running the docker file.

Remove the `export` which precedes each environment variable name so that it can be used with the docker run command.

eg. `export BASE_URL` becomes `BASE_URL`

```
docker run --env-file .env -p 8080:8080 hms-api
```

# Endpoints Implemented

[Auth Token For Client SDKs](https://www.100ms.live/docs/get-started/v2/get-started/security-and-tokens#auth-token-for-client-sdks)

| Description                       | Verb | Path   |
| --------------------------------- | ---- | ------ |
| Create a token for joining a room | POST | /token |

[Rooms](https://www.100ms.live/docs/server-side/v2/api-reference/Rooms/overview)

| Description                  | Verb | Path                   |
| ---------------------------- | ---- | ---------------------- |
| Get the list of rooms        | GET  | /rooms                 |
| Get details of a single room | GET  | /rooms/:roomId         |
| Create a new room            | POST | /rooms                 |
| Update a room                | POST | /rooms/:roomId         |
| Enable a room                | POST | /rooms/:roomId/enable  |
| Disable a room               | POST | /rooms/:roomId/disable |

[Room Codes](https://www.100ms.live/docs/server-side/v2/api-reference/room-codes/room-code-overview)

| Description                                           | Verb | Path                           |
| ----------------------------------------------------- | ---- | ------------------------------ |
| Get Room Codes for all Roles in a Room                | GET  | /room-codes/:roomId            |
| Create a Room Code for every Role in the Room at once | POST | /room-codes/:roomId            |
| Create a Room Code for a specific Role in a Room      | POST | /room-codes/:roomId/role/:role |
| Update the current state for a given Room Code.       | POST | /room-codes/update             |
| Create the auth token for a given short code          | POST | /room-codes/code/:code         |

[Active Rooms](https://www.100ms.live/docs/server-side/v2/api-reference/active-rooms/overview)

| Description                                            | Verb | Path                                |
| ------------------------------------------------------ | ---- | ----------------------------------- |
| Get details of a specific Active Room                  | GET  | /active-rooms/:roomId               |
| Get details of a specific Peer in an active Room       | GET  | /active-rooms/:roomId/peers/:peerId |
| List details of the Active Peers in a Room             | GET  | /active-rooms/:roomId/peers         |
| Update the details of a connected Peer                 | POST | /active-rooms/:roomId/peers/:peerId |
| Send Message to the room                               | POST | /active-rooms/:roomId/send-message  |
| Remove/Disconnect a connected Peer from an Active Room | POST | /active-rooms/:roomId/remove-peers  |
| End an Active Room                                     | POST | /active-rooms/:roomId/end-room      |

[Recordings](https://www.100ms.live/docs/server-side/v2/api-reference/recordings/overview)

| Description                           | Verb | Path                            |
| ------------------------------------- | ---- | ------------------------------- |
| Get recording jobs of a workspace     | GET  | /recordings                     |
| Get details of a recording            | GET  | /recordings/:recordingId        |
| Get the configuration of a recording  | GET  | /recordings/:recordingId/config |
| Start a recording for a room          | POST | /recordings/room/:roomId/start  |
| Stop all recordings running in a room | POST | /recordings/room/:roomId/stop   |
| Stop a specific recording             | POST | /recordings/:recordingId/stop   |

[Sessions](https://www.100ms.live/docs/server-side/v2/api-reference/Sessions/object)

| Description                               | Verb | Path                 |
| ----------------------------------------- | ---- | -------------------- |
| Get details of all sessions in an account | GET  | /sessions            |
| Get details of a specific session         | GET  | /sessions/:sessionId |

[Recording Assets](https://www.100ms.live/docs/server-side/v2/api-reference/recording-assets/overview)

| Description                                                 | Verb | Path                           |
| ----------------------------------------------------------- | ---- | ------------------------------ |
| Get details of all recording assets of a workspace          | GET  | /recording-assets              |
| Get details of a Recording Asset                            | GET  | /recording-assets/:assetId     |
| Generate a short-lived pre-signed URL for a recording asset | GET  | /recording-assets/:assetId/url |

[External Streams](https://www.100ms.live/docs/server-side/v2/api-reference/external-streams/overview)

| Description                                         | Verb | Path                                 |
| --------------------------------------------------- | ---- | ------------------------------------ |
| Get details of all external streams of a workspace  | GET  | /external-streams                    |
| Get the details of an external stream               | GET  | /external-streams/:streamId          |
| Start an external stream for a room                 | POST | /external-streams/room/:roomId/start |
| Stop all external streams running in the room       | POST | /external-streams/room/:roomId/stop  |
| Stop an external stream using its unique identifier | POST | /external-streams/:streamId/stop     |

[Polls](https://www.100ms.live/docs/server-side/v2/api-reference/polls/overview)

| Description          | Verb   | Path                                                     |
| -------------------- | ------ | -------------------------------------------------------- |
| Get a Poll           | GET    | /polls/:pollId                                           |
| Get Poll Sessions    | GET    | /polls/:pollId/sessions/:sessionId                       |
| List Poll Results    | GET    | /polls/:pollId/sessions/:sessionId/results               |
| Get Poll Result      | GET    | /polls/:pollId/sessions/:sessionId/results/:resultId     |
| List Poll Responses  | GET    | /polls/:pollId/sessions/:sessionId/responses             |
| Get Poll Response    | GET    | /polls/:pollId/sessions/:sessionId/responses/:responseId |
| Create a Poll        | POST   | /polls                                                   |
| Update a Poll        | POST   | /polls/:pollId                                           |
| Update Poll Question | POST   | /polls/:pollId/questions/:questionId                     |
| Update Poll Option   | POST   | /polls/:pollId/questions/:questionId/options/:optionId   |
| Delete Poll Question | DELETE | /polls/:pollId/questions/:questionId                     |
| Delete Poll Option   | DELETE | /polls/:pollId/questions/:questionId/options/:optionId   |

\*Link room with polls. Use the endpoint to create/update a room and pass an array of poll IDs i.e polls: ["id1", "id2"]

[Live Streams](https://www.100ms.live/docs/server-side/v2/api-reference/live-streams/overview)

| Description                                   | Verb | Path                                     |
| --------------------------------------------- | ---- | ---------------------------------------- |
| Get details of all livestreams of a workspace | GET  | /live-streams                            |
| Get details of a livestream                   | GET  | /live-streams/:streamId                  |
| Start a livestream for room                   | POST | /live-streams/room/:roomId/start         |
| Stop all livestreams in a room                | POST | /live-streams/room/:roomId/stop          |
| Send timed metadata for a running live stream | POST | /live-streams/:streamId/metadata         |
| Pause recording of a running livestream       | POST | /live-streams/:streamId/pause-recording  |
| Resume a paused livestream recording          | POST | /live-streams/:streamId/resume-recording |

[Policy](https://www.100ms.live/docs/server-side/v2/api-reference/policy/template-object)

| Description                                    | Verb   | Path                                   |
| ---------------------------------------------- | ------ | -------------------------------------- |
| Get the details of all templates in an account | GET    | /templates                             |
| Get the details of a specific template         | GET    | /templates/:templateId                 |
| Get details of a specific Role in a template   | GET    | /templates/:templateId/roles/:roleName |
| Get settings of a template                     | GET    | /templates/:templateId/settings        |
| Get the list of destinations in a template     | GET    | /templates/:templateId/destinations    |
| Create a template                              | POST   | /templates                             |
| Update the details of a template               | POST   | /templates/:templateId                 |
| Create or Modify a Role in a template          | POST   | /templates/:templateId/roles/:roleName |
| Update the settings of a template              | POST   | /templates/:templateId/settings        |
| Update the destinations in a template          | POST   | /templates/:templateId/destinations    |
| Delete a Role in a template                    | DELETE | /templates/:templateId/roles/:roleName |

[Analytics](https://www.100ms.live/docs/server-side/v2/api-reference/analytics/overview)

| Description          | Verb | Path       |
| -------------------- | ---- | ---------- |
| Get analytics events | GET  | /analytics |

[Stream Key](https://www.100ms.live/docs/server-side/v2/api-reference/stream-key/overview)

| Description                                         | Verb | Path                         |
| --------------------------------------------------- | ---- | ---------------------------- |
| Get the RTMP stream key and URL for a specific room | GET  | /stream-keys/:roomId         |
| Create RTMP Stream Key and URL                      | POST | /stream-keys/:roomId         |
| Disable the RTMP stream key for a specific room.    | POST | /stream-keys/:roomId/disable |

# Contributing Guide
