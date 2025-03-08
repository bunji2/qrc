package main

import (
	"fmt"
	"io"
	"os"

	"github.com/skip2/go-qrcode"
)

const (
	QRCODE_SIZE = 256
)

func runWriter(src, dst string) (code int) {
	fmt.Printf("QRCODE WRITER: %s -> %s\n", src, dst)
	err := writer(src, dst)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		code = ERR
	}

	return
}

func writer(src, dst string) (err error) {

	var r io.Reader
	var w io.Writer

	if src == "stdin" {
		r = os.Stdin
	} else {
		f, e := os.Open(src)
		if e != nil {
			err = e
			return
		}
		defer f.Close()
		r = f
	}

	if dst == "stdout" {
		w = os.Stdout
	} else {
		f2, e := os.Create(dst)
		if e != nil {
			return
		}
		defer f2.Close()
		w = f2
	}

	err = qrcWriter(r, w)

	return
}

func qrcWriter(src io.Reader, dst io.Writer) (err error) {

	bb, e := io.ReadAll(src)
	if e != nil {
		err = e
		return
	}

	contents := string(bb)
	if len(contents) < 1 {
		err = fmt.Errorf("empty text")
		return
	}

	q, e := qrcode.New(contents, qrcode.Medium)
	if e != nil {
		err = e
		return
	}

	err = q.Write(QRCODE_SIZE, dst)
	return
}
