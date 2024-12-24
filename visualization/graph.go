// visualization/graph.go

package visualization

import (
    "github.com/go-echarts/go-echarts/charts"
    "github.com/pk1151222/bug-scanners/internal/models"
    "os"
)

func GenerateVulnerabilityGraph(scanResults []models.ScanResult) {
    bar := charts.NewBar()
    bar.SetGlobalOptions(
        charts.TitleOpts{Title: "Vulnerability Severity Distribution"},
    )

    severities := []string{"Critical", "High", "Medium", "Low"}
    counts := []int{5, 15, 8, 3} // Example data

    bar.AddXAxis(severities).
        AddYAxis("Vulnerabilities", counts).
        SetSeriesOptions(
            charts.BarOpts{Label: charts.LabelOpts{Show: true}},
        )

    f, err := os.Create("vulnerability_graph.html")
    if err != nil {
        panic(err)
    }
    bar.Render(f)
}
