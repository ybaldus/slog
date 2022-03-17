package slog

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/valyala/bytebufferpool"
)

func TestLogger_newRecord_AllocTimes(t *testing.T) {
	l := Std()
	l.Output = ioutil.Discard
	defer l.Reset()

	// output: 0 times
	fmt.Println("Alloc Times:", int(testing.AllocsPerRun(100, func() {
		// logger.Info("rate", "15", "low", 16, "high", 123.2, msg)
		r := l.newRecord()
		// do something...
		l.releaseRecord(r)
	})))
}

func Test_formatArgsWithSpaces_oneElem_AllocTimes(t *testing.T) {
	l := Std()
	l.Output = ioutil.Discard
	defer l.Reset()

	// output: 1 times -> 0 times
	fmt.Println("Alloc Times:", int(testing.AllocsPerRun(100, func() {
		// logger.Info("rate", "15", "low", 16, "high", 123.2, msg)
		formatArgsWithSpaces([]interface{}{
			"msg", // 2343, -23, 123.2,
		})
	})))
}

func Test_AllocTimes_formatArgsWithSpaces_manyElem(t *testing.T) {
	l := Std()
	l.Output = ioutil.Discard
	defer l.Reset()

	// TIP:
	// `float` will alloc 2 times memory
	// `int <0`, `int > 100` will alloc 1 times memory
	fmt.Println("Alloc Times:", int(testing.AllocsPerRun(100, func() {
		// logger.Info("rate", "15", "low", 16, "high", 123.2, msg)
		formatArgsWithSpaces([]interface{}{
			"rate", -23, true, 106, "high", 123.2,
		})
	})))
}

func TestRecord_logBytes_AllocTimes(t *testing.T) {
	l := Std()
	l.Output = ioutil.Discard
	defer l.Reset()

	// use buffer pool
	bb := bytebufferpool.Get()

	// output: 50 times
	fmt.Println("Alloc Times:", int(testing.AllocsPerRun(100, func() {
		// logger.Info("rate", "15", "low", 16, "high", 123.2, msg)
		r := l.newRecord()

		_, _ = bb.Write([]byte("info message"))
		r.logBytes(InfoLevel, bb.B)

		l.releaseRecord(r)
	})))

	bytebufferpool.Put(bb)
}

func Test_AllocTimes_stringsPool(t *testing.T) {
	l := Std()
	l.Output = ioutil.Discard
	l.LowerLevelName = true
	defer l.Reset()

	var ln, cp int
	// output: 0 times
	fmt.Println("Alloc Times:", int(testing.AllocsPerRun(100, func() {
		// logger.Info("rate", "15", "low", 16, "high", 123.2, msg)

		// oldnew := stringsPool.Get().([]string)
		// defer stringsPool.Put(oldnew)

		oldnew := make([]string, 0, len(map[string]string{"a": "b"})*2+1)

		oldnew = append(oldnew, "a")
		oldnew = append(oldnew, "b")
		oldnew = append(oldnew, "c")
		// oldnew = append(oldnew, "d")

		ln = len(oldnew)
		cp = cap(oldnew)
	})))

	dump.P(ln, cp)
}

func TestLogger_Info_AllocTimes(t *testing.T) {
	l := Std()
	l.Output = ioutil.Discard
	l.LowerLevelName = true
	defer l.Reset()

	// l.Info("msg")
	// return

	// output: 0 times
	fmt.Println("Alloc Times:", int(testing.AllocsPerRun(100, func() {
		// logger.Info("rate", "15", "low", 16, "high", 123.2, msg)
		l.Info("msg")
	})))
}