{{- define "ddnsclient.matchLabels" -}}
app.kubernetes.io/name: {{ include "common.names.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "ddnsclient.labels" -}}
{{- range $name, $value := .Values.commonLabels -}}
{{ $name }}: {{ tpl $value $ }}
{{ end -}}
helm.sh/chart: {{ include "common.names.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | replace "+" "_" | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{ include "ddnsclient.matchLabels" . }}
{{- end }}

{{- define "ddnsclient.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "common.names.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
