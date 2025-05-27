/*
	https://leetcode.com/problems/logger-rate-limiter

	Design a logger system that receives a stream of messages along with their timestamps. Each unique message should only be printed at most every 10 seconds (i.e. a message printed at timestamp t will prevent other identical messages from being printed until timestamp t + 10).

	All messages will come in chronological order. Several messages may arrive at the same timestamp.

	Implement the Logger class:

		Logger() Initializes the logger object.
		bool shouldPrintMessage(int timestamp, string message) Returns true if the message should be printed in the given timestamp, otherwise returns false.

*/

package solutions

type Logger struct {
	m map[string]int
}

// NewLogger must call Constructor to pass Leetcode tests
func NewLogger() Logger {
	return Logger{make(map[string]int)}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if next, ok := l.m[message]; ok {
		if next <= timestamp {
			l.m[message] = timestamp + 10
			return true
		}
		return false
	}
	l.m[message] = timestamp + 10

	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
