{{ define "cf.common" }}
version: '1.0'
stages:{{ template "stages" .Stages }}
steps: {{ template "steps" .Steps }}
{{ end }}
{{define "steps" }}
  {{ range . }}
    - {{printf "%s" . -}}
  {{ end }}
{{ end }}
{{define "stages" }}
  {{ range . }}
    - {{printf "%s" . -}}
  {{ end }}
{{ end }}
 

{{template "cf.common" .}}