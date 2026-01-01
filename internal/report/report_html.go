package report

import (
	"html/template"
	"os"

	"github.com/ismailtsdln/socialrecon/internal/models"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>SocialRecon Report - {{.Target}}</title>
    <style>
        body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; line-height: 1.6; color: #333; max-width: 1000px; margin: 0 auto; padding: 20px; background: #f4f7f6; }
        .header { background: #1a202c; color: white; padding: 30px; border-radius: 8px; margin-bottom: 30px; }
        .dashboard { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 20px; margin-bottom: 30px; }
        .card { background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); text-align: center; }
        .card h3 { margin: 0; color: #718096; font-size: 0.9em; text-transform: uppercase; }
        .card .value { font-size: 2em; font-weight: bold; margin: 10px 0; }
        .severity-CRITICAL { color: #e53e3e; }
        .severity-HIGH { color: #dd6b20; }
        .severity-MEDIUM { color: #d69e2e; }
        .severity-LOW { color: #38a169; }
        .severity-INFO { color: #3182ce; }
        table { width: 100%; border-collapse: collapse; background: white; border-radius: 8px; overflow: hidden; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        th, td { padding: 15px; text-align: left; border-bottom: 1px solid #edf2f7; }
        th { background: #2d3748; color: white; font-weight: normal; }
        tr:hover { background: #f7fafc; }
        .badge { padding: 4px 8px; border-radius: 4px; font-size: 0.8em; font-weight: bold; }
        .badge-exists { background: #ebf8ff; color: #2b6cb0; }
        .badge-available { background: #f0fff4; color: #2f855a; }
    </style>
</head>
<body>
    <div class="header">
        <h1>SocialRecon Executive Report</h1>
        <p>Target: <strong>{{.Target}}</strong></p>
        <p>Scan completed on {{.EndTime.Format "2006-01-02 15:04:05"}}</p>
    </div>

    <div class="dashboard">
        <div class="card">
            <h3>Risk Score</h3>
            <div class="value">{{printf "%.1f" .RiskScore}}/100</div>
        </div>
        <div class="card">
            <h3>Total Findings</h3>
            <div class="value">{{len .Findings}}</div>
        </div>
        <div class="card">
            <h3>Duration</h3>
            <div class="value">{{.EndTime.Sub .StartTime}}</div>
        </div>
    </div>

    <h2>Findings Detail</h2>
    <table>
        <thead>
            <tr>
                <th>Platform</th>
                <th>Indicator</th>
                <th>Status</th>
                <th>Severity</th>
                <th>Description</th>
            </tr>
        </thead>
        <tbody>
            {{range .Findings}}
            <tr>
                <td><strong>{{.PluginName}}</strong></td>
                <td>{{.Indicator}}</td>
                <td><span class="badge badge-{{.Status}}">{{.Status}}</span></td>
                <td><span class="severity-{{.Severity}}">{{.Severity}}</span></td>
                <td>{{.Description}}</td>
            </tr>
            {{end}}
        </tbody>
    </table>
</body>
</html>
`

// ExportHTML generates a nice dashboard report
func (r *Reporter) ExportHTML(result *models.ScanResult, filename string) error {
	tmpl, err := template.New("report").Parse(htmlTemplate)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, result)
}
