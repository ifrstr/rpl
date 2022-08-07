// Package rpl, the Remote Procedure Logger.
//
// # Introduction
//
// RPL is a lite logger utility which supports serialize its Log object,
// making it easy to show log in remote procedure.
//
// # Overview
//
// Here's a typical RPL logging workflow:
//
//     ---------------------------------------------------------------
//     |                            |                                |
//     |     R  E  M  O  T  E       |          L  O  C  A  L         |
//     |                            |                                |
//     |  Logger ----- Sender ------|--- Receiver ----- FileTarget   |
//     |        \                   |       \               |        |
//     |         \                  |        \              |        |
//     |          \                 |         \          (Write to   |
//     |     EventLogTarget         |      FileTarget    os.Stdout)  |
//     |          |                 |          |                     |
//     |          |                 |          |                     |
//     |    (Write to Windows       |    (Write to log               |
//     |       Event Log)           |        file)                   |
//     |                            |                                |
//     ---------------------------------------------------------------
//
// RPL consists of 3 basic types:
// Log, Target and Source.
//
//  - Log transfers as plain old data, between remote and local.
//  - Target is the output (local) or sender (remote) of Source.
//  - Source is the log producer (remote), or receiver (local).
package rpl

// Log transfers as plain old data, between remote and local.
type Log struct {
	Ch    uint16 `json:"ch" mapstructure:"ch"`
	Level int8   `json:"level" mapstructure:"level"`
	Value string `json:"value" mapstructure:"value"`
}

// Target is the output (local) or sender (remote) of Source.
type Target interface {
	// Writer returns the Log channel.
	Writer() chan<- *Log

	// Close this [rpl.Target].
	Close()
}

// Source is the log producer (remote), or receiver (local).
type Source interface {
	// Register a Target.
	Register(target Target)

	// Close this [rpl.Source] and all registered [rpl.Target]s.
	Close()
}
