package ports

import (
	linechart "conso-tracker/src/external/components/line-chart"
	"mime/multipart"
)

type PuissanceProcessor interface {
	ProcessChart(file multipart.File) linechart.Chart
}
