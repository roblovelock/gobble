package parser

import (
	"fmt"
	"io"
	"time"
)

type (
	parserDebug struct {
		callCount          int
		successCount       int
		errorCount         int
		totalBytesConsumed int64
		totalTime          time.Duration
	}

	callDebug[T any] struct {
		*parserDebug
		callTime      time.Duration
		startOffset   int64
		endOffset     int64
		bytesConsumed int64
		result        T
		err           error
	}
)

func (d *parserDebug) String() string {
	return fmt.Sprintf(
		"calls: %d\ttime: %s\tbytes: %d\tsuccess: %d\terror: %d",
		d.callCount, d.totalTime.String(), d.totalBytesConsumed, d.successCount, d.errorCount,
	)
}

func (d *callDebug[T]) String() string {
	return fmt.Sprintf(
		"start: %d\tend: %d\ttime: %d\tbytes: %d\tsuccess: '%v'\terror: '%v'\nTotals(%s)",
		d.startOffset, d.endOffset, d.callTime.Nanoseconds(), d.bytesConsumed, d.result, d.err, d.parserDebug.String(),
	)
}

func Debug[R Reader, T any](p Parser[R, T]) Parser[R, T] {
	debug := parserDebug{}
	return func(in R) (T, error) {
		call := callDebug[T]{
			parserDebug: &debug,
		}
		debug.callCount++
		call.startOffset, _ = in.Seek(0, io.SeekCurrent)
		startTime := time.Unix(0, time.Now().UnixNano())
		call.result, call.err = p(in)
		call.callTime = time.Since(startTime)
		call.endOffset, _ = in.Seek(0, io.SeekCurrent)
		debug.totalTime += call.callTime
		call.bytesConsumed = call.endOffset - call.startOffset
		debug.totalBytesConsumed += call.bytesConsumed
		if call.err != nil {
			debug.errorCount++
		} else {
			debug.successCount++
		}

		fmt.Printf("%s\n", call.String())
		return call.result, call.err
	}
}
