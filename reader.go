package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"os"

	"github.com/MordFustang21/gozbar"
)

func runReader(src, dst string) (code int) {
	fmt.Printf("QRCODE READER: %s -> %s\n", src, dst)
	err := reader(src, dst)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		code = ERR
	}

	return
}

func reader(src, dst string) (err error) {

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

	err = qrcReader(r, w)
	return
}

func qrcReader(src io.Reader, dst io.Writer) (err error) {

	i, e := png.Decode(src)
	if e != nil {
		err = e
		return
	}

	img := gozbar.FromImage(i)
	s := gozbar.NewScanner()
	//e = s.SetConfig(0, gozbar.CFG_ENABLE, 1)
	e = s.SetConfig(gozbar.QRCODE, gozbar.CFG_ENABLE, 1)
	if e != nil {
		err = e
		return
	}

	e = s.Scan(img)
	if e != nil {
		err = e
		return
	}

	var b bytes.Buffer

	img.First().Each(func(str string) {
		_, e = b.WriteString(str)
		if e != nil {
			err = e
			return
		}
	})

	_, e = b.WriteTo(dst)
	if e != nil {
		err = e
	}

	return
}
