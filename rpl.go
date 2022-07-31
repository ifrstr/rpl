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
//     |  Logger --- SenderTarget --|-- Receiver ------ FileTarget   |
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
	Level int8   `json:"level"`
	Value string `json:"value"`
}

// Target is the output (local) or sender (remote) of Source.
type Target interface {
	Writer() chan<- Log
}

// Source is the log producer (remote), or receiver (local).
type Source interface {
	Register(target Target)
}
