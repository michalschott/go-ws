{{- define "go-ws.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "go-ws.labels" -}}
app.kubernetes.io/name: {{ include "go-ws.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/managed-by: Helm
{{- end -}}

{{- define "go-ws.name" -}}
{{- .Chart.Name -}}
{{- end -}}
