package goalpdf

import (
	"github.com/Hari-Kiri/gofpdf"
)

/*
New returns a pointer to a new "Jung-Kurt goFpdf" instance. Its methods are subsequently called to produce a single PDF
document.

pageOrientation specifies the default page orientation. For portrait mode, specify "P" or "Portrait". For landscape mode,
specify "L" or "Landscape". An empty string will be replaced with "P".

measurementUnit specifies the unit of length used in size parameters for elements other than fonts, which are always measured
in points. Specify "pt" for point, "mm" for millimeter, "cm" for centimeter, or "in" for inch. An empty string will be
replaced with "mm".

pagesSize specifies the page size. Acceptable values are "A3", "A4", "A5", "Letter", "Legal", or "Tabloid". An empty string
will be replaced with "A4".

fontDirectory specifies the file system location in which font resources will be found. An empty string is replaced with
".". This argument only needs to reference an actual directory if a font other than one of the core fonts is used. The
core fonts are "courier", "helvetica" (also called "arial"), "times", and "zapfdingbats" (also called "symbol").

marginLeft set the left margin.

marginTop set the top margin.

marginRight set the right margin.

fontName specifies the font name. It can be either a name defined by (gofpdf.Fpdf).AddFont(),
(gofpdf.Fpdf).AddFontFromReader() or one of the standard families (case insensitive): "Courier" for fixed-width, "Helvetica"
or "Arial" for sans serif, "Times" for serif, "Symbol" or "ZapfDingbats" for symbolic.

fontStyle can be "B" (bold), "I" (italic), "U" (underscore), "S" (strike-out) or any combination. The default value
(specified with an empty string) is regular. Bold and italic styles do not apply to Symbol and ZapfDingbats.

fontSize is the font size measured in points. The default value is the current size. If no size has been specified since
the beginning of the document, the value taken is 12.
*/
func New(pageOrientation string, measurementUnit string, pagesSize string, fontDirectory string, marginLeft, marginTop,
	marginRight float64, fontName string, fontStyle string, fontSize float64) (f *gofpdf.Fpdf) {
	goFpdfPointer := gofpdf.New(pageOrientation, measurementUnit, pagesSize, fontDirectory)
	goFpdfPointer.SetMargins(marginLeft, marginTop, marginRight)
	goFpdfPointer.SetFont(fontName, fontStyle, fontSize)
	return goFpdfPointer
}

/*
AddPage adds a new page to the document.

goFpdfInstance an instance of gofpdf. Its compatible with "Jung-Kurt goFpdf" but unfortunately the repository was in
public archived.
*/
func AddPage(goFpdfInstance *gofpdf.Fpdf) {
	goFpdfInstance.AddPage()
}

/*
AddSingleRow Adds one or more columns to a row. In each column there will be three cells, namely: title cells,
subtitle cells and text cells.

goFpdfInstance an instance of gofpdf. Its compatible with "Jung-Kurt goFpdf" but unfortunately the repository was in
public archived.

columnWidth will set width of every column in a row. Every column will have same width with this method.

newlineSpacing set the space between row inside cell. Every column will have same newline spacing with this method.

xCoordinate set the starting point of the first cell to rendered according to the left of the screen.

yCoordinate set the starting point of the cell to rendered according to the top of the screen.

columnCount total of the column of the row.

cellsMargin set the margin of cells in a column. Every column will have same cells margin with this method.

backgroundColor set the background of the row. Fill this parameter with RGB (Red Green Blue) value inside 3 length slice
data type.

drawLineBorder set columns have outline border or not. The border will rendered in black color.

title is column title string in slice.

titleFontSize set the size of the title font. Size is specified in points (1/ 72 inch).

titleFontStyle can be "B" (bold), "I" (italic), "U" (underscore), "S" (strike-out) or any combination. The default value
(specified with an empty string) is regular. Bold and italic styles do not apply to Symbol and ZapfDingbats.

titleAlignment specifies how the title text is to be positioned within the cell. Horizontal alignment is controlled by
including "L", "C" or "R" (left, center, right) in alignStr. Vertical alignment is controlled by including "T", "M", "B" or
"A" (top, middle, bottom, baseline) in alignStr. The default alignment is left middle.

subtitle is column subtitle string in slice.

subtitleFontSize set the size of the subtitle font. Size is specified in points (1/ 72 inch).

subtitleFontStyle can be "B" (bold), "I" (italic), "U" (underscore), "S" (strike-out) or any combination. The default value
(specified with an empty string) is regular. Bold and italic styles do not apply to Symbol and ZapfDingbats.

subtitleAlignment specifies how the subtitle text is to be positioned within the cell. Horizontal alignment is controlled by
including "L", "C" or "R" (left, center, right) in alignStr. Vertical alignment is controlled by including "T", "M", "B" or
"A" (top, middle, bottom, baseline) in alignStr. The default alignment is left middle.

text is column text string in slice.

textFontSize set the size of the text font. Size is specified in points (1/ 72 inch).

textFontStyle can be "B" (bold), "I" (italic), "U" (underscore), "S" (strike-out) or any combination. The default value
(specified with an empty string) is regular. Bold and italic styles do not apply to Symbol and ZapfDingbats.

textAlignment specifies how the subtitle text is to be positioned within the cell. Horizontal alignment is controlled by
including "L", "C" or "R" (left, center, right) in alignStr. Vertical alignment is controlled by including "T", "M", "B" or
"A" (top, middle, bottom, baseline) in alignStr. The default alignment is left middle.

Notes: title, subtitle and text must have same length or it will be panic.

rowLineBreak performs a line break. The current abscissa goes back to the left margin and the ordinate increases by the
amount passed in this parameter. A negative value indicates the height of the last printed cell.
*/
func AddSingleRow(goFpdfInstance *gofpdf.Fpdf, columnWidth, newlineSpacing, xCoordinate, yCoordinate float64,
	columnCount int, cellsMargin float64, backgroundColor [3]int, drawLineBorder bool,
	title []string, titleFontSize float64, titleFontStyle, titleAlignment string,
	subtitle []string, subtitleFontSize float64, subtitleFontStyle, subtitleAlignment string,
	text []string, textFontSize float64, textFontStyle, textAlignment string,
	rowLineBreak float64) {
	type cellType struct {
		list   [][]byte
		height float64
	}
	var (
		titleList    []cellType = make([]cellType, columnCount)
		subtitleList []cellType = make([]cellType, columnCount)
		textList     []cellType = make([]cellType, columnCount)
		titleCell    cellType
		subtitleCell cellType
		textCell     cellType
	)
	var rectStyle string
	if drawLineBorder {
		rectStyle = "FD"
	}
	if !drawLineBorder {
		rectStyle = "F"
	}
	for row := 0; row < 1; row++ {
		columnsHeight := newlineSpacing
		// Cell height calculation loop
		for columnCounting := 0; columnCounting < columnCount; columnCounting++ {
			// Calculate height & make bytes array of word title cell
			goFpdfInstance.SetFontSize(titleFontSize)
			goFpdfInstance.SetFontStyle(titleFontStyle)
			titleCell.list = goFpdfInstance.SplitLines([]byte(title[columnCounting]), columnWidth-cellsMargin-cellsMargin)
			titleCell.height = float64(len(titleCell.list)) * newlineSpacing
			if columnsHeight < titleCell.height {
				columnsHeight += titleCell.height
			}
			titleList[columnCounting] = titleCell
			// Calculate height & make bytes array of word subtitle cell
			goFpdfInstance.SetFontSize(subtitleFontSize)
			goFpdfInstance.SetFontStyle(subtitleFontStyle)
			subtitleCell.list = goFpdfInstance.SplitLines([]byte(subtitle[columnCounting]),
				columnWidth-cellsMargin-cellsMargin)
			subtitleCell.height = float64(len(subtitleCell.list)) * newlineSpacing
			if columnsHeight < subtitleCell.height {
				columnsHeight += subtitleCell.height
			}
			subtitleList[columnCounting] = subtitleCell
			// Calculate height & make bytes array of word text cell
			goFpdfInstance.SetFontSize(textFontSize)
			goFpdfInstance.SetFontStyle(textFontStyle)
			textCell.list = goFpdfInstance.SplitLines([]byte(text[columnCounting]), columnWidth-cellsMargin-cellsMargin)
			textCell.height = float64(len(textCell.list)) * newlineSpacing
			if columnsHeight < textCell.height {
				columnsHeight += textCell.height
			}
			textList[columnCounting] = textCell
		}
		// Cell render loop
		for columnRendering := 0; columnRendering < columnCount; columnRendering++ {
			// Set cell background color
			goFpdfInstance.SetFillColor(backgroundColor[0], backgroundColor[1], backgroundColor[2])
			// Set border
			goFpdfInstance.Rect(xCoordinate, yCoordinate, columnWidth,
				newlineSpacing+newlineSpacing+columnsHeight+cellsMargin+cellsMargin, rectStyle)
			titleCell = titleList[columnRendering]
			subtitleCell = subtitleList[columnRendering]
			textCell = textList[columnRendering]
			cellY := yCoordinate + cellsMargin + (columnsHeight-textCell.height)/2
			// Render title cell
			for splitH := 0; splitH < len(titleCell.list); splitH++ {
				goFpdfInstance.SetXY(xCoordinate+cellsMargin, cellY)
				goFpdfInstance.SetFontSize(titleFontSize)
				goFpdfInstance.SetFontStyle(titleFontStyle)
				goFpdfInstance.CellFormat(columnWidth-cellsMargin-cellsMargin, newlineSpacing, string(titleCell.list[splitH]),
					"", 0, titleAlignment, false, 0, "")
				cellY += newlineSpacing
			}
			// Render subtitle cell
			for splitI := 0; splitI < len(subtitleCell.list); splitI++ {
				goFpdfInstance.SetXY(xCoordinate+cellsMargin, cellY)
				goFpdfInstance.SetFontSize(subtitleFontSize)
				goFpdfInstance.SetFontStyle(subtitleFontStyle)
				goFpdfInstance.CellFormat(columnWidth-cellsMargin-cellsMargin, newlineSpacing,
					string(subtitleCell.list[splitI]), "", 0, subtitleAlignment, false, 0, "")
				cellY += newlineSpacing
			}
			// Render text cell
			for splitJ := 0; splitJ < len(textCell.list); splitJ++ {
				goFpdfInstance.SetXY(xCoordinate+cellsMargin, cellY)
				goFpdfInstance.SetFontSize(textFontSize)
				goFpdfInstance.SetFontStyle(textFontStyle)
				goFpdfInstance.CellFormat(columnWidth-cellsMargin-cellsMargin, newlineSpacing, string(textCell.list[splitJ]),
					"", 0, textAlignment, false, 0, "")
				cellY += newlineSpacing
			}
			xCoordinate += columnWidth
		}
		yCoordinate += columnsHeight + cellsMargin + cellsMargin
	}
	goFpdfInstance.Ln(rowLineBreak + newlineSpacing)
}
