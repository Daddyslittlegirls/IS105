package oppgaver

//////////////////////////////////
// 	ICA04, IS-105	 	//
// 	2017, 	Erlend Wiklem
//	https://github.com/Zwirc//
//////////////////////////////////
import (
	"./huffman"

	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

var huffstring []string

func Oppgave4a() {
	/*
	 Helse og Idrettsdag		1420	1829
	 Humanoira og pedagogikk	1182	1525
	 Kunstfag			 293	 420
	 Teknologi og realfag		1337	2166
	 Lærerutdanning			1158	1506
	 Økonomi og samfunnsvitenskap	2398	3098
	 Antall mulige utfall:		7788	10544
	*/
	UIAlist()

}
func Oppgave4b() {
	UIAinfo()
	fmt.Println(" Fakulitetet Kunstfag får minst informasjon som sett over.")

}
func Oppgave4c() {
	UIAhuffman()

}
func Oppgave4d() {
	UIAaverage()
}

func Oppgave4e() {
	EncodeDecore()
}

// Oppgave 4a
var TOT float64 = 10544
var HI float64 = 1829
var HP float64 = 1525
var K float64 = 420
var TR float64 = 2166
var L float64 = 1506
var OS float64 = 3098
var hundred float64 = 100

func UIAlist() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Fakulitet \t år 2008 \t år 2015 \t")
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintln(w, "Helse og Idrettsdag \t 1420 \t 1829 \t")
	fmt.Fprintln(w, "Humanoira og pedagogikk \t 1182 \t 1525 \t")
	fmt.Fprintln(w, "Kunstfag \t 293 \t 420 \t")
	fmt.Fprintln(w, "Teknologi og realfag \t 1337 \t 2166 \t")
	fmt.Fprintln(w, "Lærerutdanning \t 1158 \t 1506 \t")
	fmt.Fprintln(w, "Økonomi og samfunnsvitenskap \t 2398 \t 3098 \t")
	w.Flush()

	// Sannsynlighet = Muligheter / total * 100 %
	fmt.Println()
	fmt.Fprintln(w, "Fakulitet \t Sannsynlighet 2015 \t")
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintln(w, "Helse og Idrettsdag \t ", HI/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Humaniora og pedagogikk \t ", HP/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Kunstfag \t", K/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Teknologi og realfag \t ", TR/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Lærerutdanning \t ", L/TOT*hundred, "% \t")
	fmt.Fprintln(w, "Økonomi og samfunnsvitenskap \t ", OS/TOT*hundred, "% \t")
	w.Flush()
}

// Oppgave 4b
func UIAinfo() {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Fakulitet \t Informasjon \t")
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintln(w, "Helse og Idrettsdag \t", strconv.FormatInt(int64(HI), 2), " \t")
	fmt.Fprintln(w, "Humaniora og pedagogikk \t ", strconv.FormatInt(int64(HP), 2), " \t")
	fmt.Fprintln(w, "Kunstfag \t ", strconv.FormatInt(int64(K), 2), " \t")
	fmt.Fprintln(w, "Teknologi og realfag \t ", strconv.FormatInt(int64(TR), 2), " \t")
	fmt.Fprintln(w, "Lærerutdanning \t ", strconv.FormatInt(int64(L), 2), " \t")
	fmt.Fprintln(w, "Økonomi og samfunnsvitenskap \t", strconv.FormatInt(int64(OS), 2), " \t")
	w.Flush()
}

// Oppgave 4c
// Skriver så ut resultatet i table
func UIAhuffman() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Fakulitet \t Informasjon \t")
	fmt.Fprintln(w, "\t\t")
	fmt.Fprintln(w, "Helse og Idrettsdag \t", "111", " \t")
	fmt.Fprintln(w, "Humaniora og pedagogikk \t ", "110", " \t")
	fmt.Fprintln(w, "Kunstfag \t ", "000", " \t")
	fmt.Fprintln(w, "Teknologi og realfag \t ", "01", " \t")
	fmt.Fprintln(w, "Lærerutdanning \t ", "001", " \t")
	fmt.Fprintln(w, "Økonomi og samfunnsvitenskap \t", "10", " \t")
	w.Flush()

}

// Oppgave 4d
func UIAaverage() {

	// Av bit lengden
	fmt.Println("Antall * (Lengden på melding * sansynlighet) + (Lengden på ......")
	fmt.Print("Lengde på melding til 100 stk fra huffman bit lengde: ")
	tott := 100*(3*0.1734) + (3 * 0.1446) + (3 * 0.0398) + (2 * 0.2054) + (3 * 0.1428) + (2 * 0.2938)
	fmt.Print(tott)
	fmt.Print("\nTil binary:")
	fmt.Print(strconv.FormatInt(int64(tott), 2))
	fmt.Println("")
}

// Oppgave 4e
func EncodeDecore() {

	huffmancode := make(map[string]int)
	huffmancode["Helse og Idrettsdag "] = int(HI)
	huffmancode["Humaniora og pedagogikk"] = int(HP)
	huffmancode["Kunstfag"] = int(K)
	huffmancode["Teknologi og realfag"] = int(TR)
	huffmancode["Lærerutdanning"] = int(L)
	huffmancode["Økonomi og samfunnsvitenskap"] = int(OS)

	tree := huffman.BuildTree(huffmancode)

	// print out results
	w := tabwriter.NewWriter(os.Stdout, 0, 100, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Verdi \t Huffman \t Fakulitet \t")
	w.Flush()
	huffman.PrintCodes(tree, []byte{})
}
