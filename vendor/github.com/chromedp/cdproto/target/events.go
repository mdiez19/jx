package target

// Code generated by chromedp-gen. DO NOT EDIT.

// EventAttachedToTarget issued when attached to target because of
// auto-attach or attachToTarget command.
type EventAttachedToTarget struct {
	SessionID          SessionID `json:"sessionId"` // Identifier assigned to the session used to send/receive messages.
	TargetInfo         *Info     `json:"targetInfo"`
	WaitingForDebugger bool      `json:"waitingForDebugger"`
}

// EventDetachedFromTarget issued when detached from target for any reason
// (including detachFromTarget command). Can be issued multiple times per target
// if multiple sessions have been attached to it.
type EventDetachedFromTarget struct {
	SessionID SessionID `json:"sessionId"` // Detached session identifier.
}

// EventReceivedMessageFromTarget notifies about a new protocol message
// received from the session (as reported in attachedToTarget event).
type EventReceivedMessageFromTarget struct {
	SessionID SessionID `json:"sessionId"` // Identifier of a session which sends a message.
	Message   string    `json:"message"`
}

// EventTargetCreated issued when a possible inspection target is created.
type EventTargetCreated struct {
	TargetInfo *Info `json:"targetInfo"`
}

// EventTargetDestroyed issued when a target is destroyed.
type EventTargetDestroyed struct {
	TargetID ID `json:"targetId"`
}

// EventTargetInfoChanged issued when some information about a target has
// changed. This only happens between targetCreated and targetDestroyed.
type EventTargetInfoChanged struct {
	TargetInfo *Info `json:"targetInfo"`
}