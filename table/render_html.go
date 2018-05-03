package table

import "strings"

// RenderHTML renders the Table in HTML format. Example:
//  <table class="go-pretty-table">
//    <thead>
//    <tr>
//      <th align="right">#</th>
//      <th>First Name</th>
//      <th>Last Name</th>
//      <th align="right">Salary</th>
//      <th>&nbsp;</th>
//    </tr>
//    </thead>
//    <tbody>
//    <tr>
//      <td align="right">1</td>
//      <td>Arya</td>
//      <td>Stark</td>
//      <td align="right">3000</td>
//      <td>&nbsp;</td>
//    </tr>
//    <tr>
//      <td align="right">20</td>
//      <td>Jon</td>
//      <td>Snow</td>
//      <td align="right">2000</td>
//      <td>You know nothing, Jon Snow!</td>
//    </tr>
//    <tr>
//      <td align="right">300</td>
//      <td>Tyrion</td>
//      <td>Lannister</td>
//      <td align="right">5000</td>
//      <td>&nbsp;</td>
//    </tr>
//    </tbody>
//    <tfoot>
//    <tr>
//      <td align="right">&nbsp;</td>
//      <td>&nbsp;</td>
//      <td>Total</td>
//      <td align="right">10000</td>
//      <td>&nbsp;</td>
//    </tr>
//    </tfoot>
//  </table>
func (t *Table) RenderHTML() string {
	t.init()

	var out strings.Builder
	out.WriteString("<table class=\"")
	out.WriteString(t.htmlCSSClass)
	out.WriteString("\">\n")
	t.htmlRenderRows(&out, t.rowsHeader, true, false)
	t.htmlRenderRows(&out, t.rows, false, false)
	t.htmlRenderRows(&out, t.rowsFooter, false, true)
	out.WriteString("</table>")
	return t.render(&out)
}

func (t *Table) htmlRenderRow(out *strings.Builder, row Row, isHeader bool, isFooter bool) {
	out.WriteString("  <tr>\n")
	for idx := range t.maxColumnLengths {
		colStr := ""
		if idx < len(row) {
			colStr = row[idx].(string)
		}

		// header uses "th" instead of "td"
		colTagName := "td"
		if isHeader {
			colTagName = "th"
		}

		// determine the HTML "align"/"valign" property values
		align := t.getAlign(idx).HTMLProperty()
		vAlign := t.getVAlign(idx).HTMLProperty()

		// write the row
		out.WriteString("    <")
		out.WriteString(colTagName)
		if align != "" {
			out.WriteRune(' ')
			out.WriteString(align)
		}
		if vAlign != "" {
			out.WriteRune(' ')
			out.WriteString(vAlign)
		}
		out.WriteString(">")
		if len(colStr) > 0 {
			out.WriteString(strings.Replace(colStr, "\n", "<br/>", -1))
		} else {
			out.WriteString("&nbsp;")
		}
		out.WriteString("</")
		out.WriteString(colTagName)
		out.WriteString(">\n")
	}
	out.WriteString("  </tr>\n")
}

func (t *Table) htmlRenderRows(out *strings.Builder, rows []Row, isHeader bool, isFooter bool) {
	if len(rows) > 0 {
		// determine that tag to use based on the type of the row
		rowsTag := "tbody"
		if isHeader {
			rowsTag = "thead"
		} else if isFooter {
			rowsTag = "tfoot"
		}

		// render all the rows enclosed by the "rowsTag"
		out.WriteString("  <")
		out.WriteString(rowsTag)
		out.WriteString(">\n")
		for _, row := range rows {
			t.htmlRenderRow(out, row, isHeader, isFooter)
		}
		out.WriteString("  </")
		out.WriteString(rowsTag)
		out.WriteString(">\n")
	}
}