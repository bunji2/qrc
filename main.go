// qrc
// [使い方]
//  ●Writer モード：テキストファイルからQRコードの画像ファイルを出力する
//  % qrc -mode writer -i src.txt -o dst.png
//  メモ：
//  ・-i 引数を省略時には標準入力から読み出す。
//  ・-o 引数を省略時には "out.png" に出力する。
//
//  ●Reader モード：画像ファイル内のQRコードをデコードしてテキストファイルに出力する
//  % qrc -mode reader -i src.png -o dst.txt
//  メモ：
//  ・-i 引数を省略時には標準入力から読み出す。
//  ・-o 引数を省略時には "out.txt" に出力する。
//
//  -mode 引数を省略時にはデフォルトで "writer モード" が設定される。

// 例：
// echo で指定した文字列を持つQRコードを画像ファイル out.png に出力する。
//  % echo hogehoge | qrc

//  明示的に、入力に標準入力 "stdin"、出力には標準出力 "stdout" を指定できる。

package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	OK  = 0
	ERR = 1

	QRCODE_READER = 100
	QRCODE_WRITER = 200

	DEFAULT_MODE           = "writer"
	DEFAULT_IN_FILE        = "stdin"
	DEFAULT_OUT_FILE       = ""
	DEFAULT_OUT_IMAGE_FILE = "./out.png"
	DEFAULT_OUT_TEXT_FILE  = "stdout"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s / %s / %s\n", PROGRAM_NAME, VERSION, AUTHOR)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	mode := flag.String("mode", DEFAULT_MODE, "reader / writer")
	src := flag.String("i", DEFAULT_IN_FILE, "input file")
	dst := flag.String("o", DEFAULT_OUT_FILE, "output file")

	flag.Parse()

	//fmt.Fprintln(os.Stderr, "[DEBUG] *mode =", *mode)

	// モードチェック
	m := QRCODE_WRITER
	if *mode == "reader" {
		m = QRCODE_READER
	} else {
		if *mode != "writer" {
			flag.Usage()
			os.Exit(OK)
		}
	}

	// 出力先が省略された時のデフォルトの設定
	if *dst == DEFAULT_OUT_FILE {
		if m == QRCODE_READER {
			*dst = DEFAULT_OUT_TEXT_FILE
		} else {
			*dst = DEFAULT_OUT_IMAGE_FILE
		}
	}

	// 出力先が標準出力じゃないときはプログラム名・バージョン・作成者名を表示
	if *dst != "stdout" {
		fmt.Printf("%s / %s / %s\n", PROGRAM_NAME, VERSION, AUTHOR)
	}

	os.Exit(run(m, *src, *dst))
}

func run(mode int, src, dst string) int {

	if mode == QRCODE_READER {
		return runReader(src, dst)
	}

	if mode == QRCODE_WRITER {
		return runWriter(src, dst)
	}

	fmt.Fprintf(os.Stderr, "unknown mode: %d\n", mode)
	return ERR
}
