package hms_errors

import (
	"errors"
)

var (
	ErrMissingAppAccessKey = errors.New("provide your app access key in the environment variables")

	ErrMissingAppSecretKey = errors.New("provide your app secret in the environment variables")

	ErrMissingBaseUrl = errors.New("provide the base url in the environment variables")

	ErrMissingRoomId = errors.New("provide a room ID")

	ErrMissingRoomIdAndRole = errors.New("provide a room ID and role")

	ErrMissingAuthCode = errors.New("provide auth code")

	ErrMissingRoomIdAndPeerId = errors.New("provide a room ID and peer ID")

	ErrMissingStreamId = errors.New("provide a stream ID")

	ErrMissingTemplateId = errors.New("provide a template ID")

	ErrMissingTemplateIdAndRoleName = errors.New("provide a template ID and a role name")

	ErrMissingPollId = errors.New("provide a poll ID")

	ErrMissingPollIdAndQuestionId = errors.New("provide a poll ID and a question ID")

	ErrMissingPollIdAndSessionId = errors.New("provide a poll ID and a session ID")

	ErrMissingPollIdAndSessionIdAndResultID = errors.New("provide a poll ID, session ID and result ID")

	ErrMissingPollIdAndQuestionIdAndOptionId = errors.New("provide a poll ID, question ID and option ID")

	ErrMissingRecordingId = errors.New("provide a recording ID")

	ErrMissingSessionId = errors.New("provide a session ID")

	ErrMissingAssetId = errors.New("provide a asset ID")
)
