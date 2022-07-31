// Package rpl, the Remote Procedure Logger.
//
// # Introduction
//
// RPL is a lite logger utility which supports serialize its Log object
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
// Log, Target and Logger.
//
//  - Log transfers as plain old data, between remote and local.
//  - Target is the output (local) or sender (remote) of Logger.
//  - Logger is the log producer (remote), or receiver (local).
package rpl

// Log transfers as plain old data, between remote and local.
type Log struct {
	Value string `json:"value"`
}

// Target is the output (local) or sender (remote) of Logger.
type Target interface {
	Log(log Log)
}

// Logger is the log producer (remote), or receiver (local).
type Logger interface {
	Register(target Target)
}
