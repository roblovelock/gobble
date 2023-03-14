package parser

import (
	"fmt"
	"io"
	"time"
)

type (
	parserDebug[R Reader, T any] struct {
		parser             Parser[R, T]
		callCount          int
		successCount       int
		errorCount         int
		totalBytesConsumed int64
		totalTime          time.Duration
	}

	callDebug[R Reader, T any] struct {
		*parserDebug[R, T]
		callTime      time.Duration
		startOffset   int64
		endOffset     int64
		bytesConsumed int64
		result        T
		err           error
	}
)

func (d *parserDebug[R, T]) String() string {
	return fmt.Sprintf(
		"calls: %d\ttime: %s\tbytes: %d\tsuccess: %d\terror: %d",
		d.callCount, d.totalTime.String(), d.totalBytesConsumed, d.successCount, d.errorCount,
	)
}

func (d *callDebug[R, T]) String() string {
	return fmt.Sprintf(
		"start: %d\tend: %d\ttime: %d\tbytes: %d\tsuccess: '%v'\terror: '%v'\nTotals(%s)",
		d.startOffset, d.endOffset, d.callTime.Nanoseconds(), d.bytesConsumed, d.result, d.err, d.parserDebug.String(),
	)
}

func (d *parserDebug[R, T]) Parse(in R) (T, error) {
	call := callDebug[R, T]{
		parserDebug: d,
	}
	d.callCount++
	call.startOffset, _ = in.Seek(0, io.SeekCurrent)
	startTime := time.Unix(0, time.Now().UnixNano())
	call.result, call.err = d.parser.Parse(in)
	call.callTime = time.Since(startTime)
	call.endOffset, _ = in.Seek(0, io.SeekCurrent)
	d.totalTime += call.callTime
	call.bytesConsumed = call.endOffset - call.startOffset
	d.totalBytesConsumed += call.bytesConsumed
	if call.err != nil {
		d.errorCount++
	} else {
		d.successCount++
	}

	fmt.Printf("%s\n", call.String())
	return call.result, call.err
}

func (d *parserDebug[R, T]) ParseBytes(in []byte) (T, []byte, error) {
	var out []byte
	call := callDebug[R, T]{
		parserDebug: d,
	}
	d.callCount++
	call.startOffset = 0
	startTime := time.Unix(0, time.Now().UnixNano())
	call.result, out, call.err = d.parser.ParseBytes(in)
	call.callTime = time.Since(startTime)
	call.endOffset = int64(len(in) - len(out))
	d.totalTime += call.callTime
	call.bytesConsumed = call.endOffset - call.startOffset
	d.totalBytesConsumed += call.bytesConsumed
	if call.err != nil {
		d.errorCount++
	} else {
		d.successCount++
	}

	fmt.Printf("%s\n", call.String())
	return call.result, out, call.err
}

func Debug[R Reader, T any](p Parser[R, T]) Parser[R, T] {
	return &parserDebug[R, T]{parser: p}
}
