package logs

import (
	"github.com/go-test/deep"
	"reflect"

	"testing"

	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/stretchr/testify/assert"
)

func TestTemplateLog(t *testing.T) {
	p := NewTemplateLogPlugin("https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logEventViewer:group=/flyte-production/kubernetes;stream=var.log.containers.{{.podName}}_{{.namespace}}_{{.containerName}}-{{.containerId}}.log", core.TaskLog_JSON)
	tl, err := p.GetTaskLog(
		"f-uuid-driver",
		"flyteexamples-production",
		"spark-kubernetes-driver",
		"cri-o://abc",
		"main_logs")
	assert.NoError(t, err)
	assert.Equal(t, tl.GetName(), "main_logs")
	assert.Equal(t, tl.GetMessageFormat(), core.TaskLog_JSON)
	assert.Equal(t, "https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logEventViewer:group=/flyte-production/kubernetes;stream=var.log.containers.f-uuid-driver_flyteexamples-production_spark-kubernetes-driver-abc.log", tl.Uri)
}

func Test_templateLogPlugin_Regression(t *testing.T) {
	type fields struct {
		templateUri   string
		messageFormat core.TaskLog_MessageFormat
	}
	type args struct {
		podName       string
		namespace     string
		containerName string
		containerID   string
		logName       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    core.TaskLog
		wantErr bool
	}{
		{
			"cloudwatch",
			fields{
				templateUri:   "https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logEventViewer:group=/flyte-production/kubernetes;stream=var.log.containers.{{.podName}}_{{.namespace}}_{{.containerName}}-{{.containerId}}.log",
				messageFormat: core.TaskLog_JSON,
			},
			args{
				podName:       "f-uuid-driver",
				namespace:     "flyteexamples-production",
				containerName: "spark-kubernetes-driver",
				containerID:   "cri-o://abc",
				logName:       "main_logs",
			},
			core.TaskLog{
				Uri:           "https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logEventViewer:group=/flyte-production/kubernetes;stream=var.log.containers.f-uuid-driver_flyteexamples-production_spark-kubernetes-driver-abc.log",
				MessageFormat: core.TaskLog_JSON,
				Name:          "main_logs",
			},
			false,
		},
		{
			"stackdriver",
			fields{
				templateUri:   "https://console.cloud.google.com/logs/viewer?project=test-gcp-project&angularJsUrl=%2Flogs%2Fviewer%3Fproject%3Dtest-gcp-project&resource=aws_ec2_instance&advancedFilter=resource.labels.pod_name%3D{{.podName}}",
				messageFormat: core.TaskLog_JSON,
			},
			args{
				podName:       "podName",
				namespace:     "flyteexamples-production",
				containerName: "spark-kubernetes-driver",
				containerID:   "cri-o://abc",
				logName:       "main_logs",
			},
			core.TaskLog{
				Uri:           "https://console.cloud.google.com/logs/viewer?project=test-gcp-project&angularJsUrl=%2Flogs%2Fviewer%3Fproject%3Dtest-gcp-project&resource=aws_ec2_instance&advancedFilter=resource.labels.pod_name%3DpodName",
				MessageFormat: core.TaskLog_JSON,
				Name:          "main_logs",
			},
			false,
		},
		{
			"kubernetes",
			fields{
				templateUri:   "https://dashboard.k8s.net/#!/log/{{.namespace}}/{{.podName}}/pod?namespace={{.namespace}}",
				messageFormat: core.TaskLog_JSON,
			},
			args{
				podName:       "flyteexamples-development-task-name",
				namespace:     "flyteexamples-development",
				containerName: "ignore",
				containerID:   "ignore",
				logName:       "main_logs",
			},
			core.TaskLog{
				Uri:           "https://dashboard.k8s.net/#!/log/flyteexamples-development/flyteexamples-development-task-name/pod?namespace=flyteexamples-development",
				MessageFormat: core.TaskLog_JSON,
				Name:          "main_logs",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TemplateLogPlugin{
				templateUri:   tt.fields.templateUri,
				messageFormat: tt.fields.messageFormat,
			}

			got, err := s.GetTaskLog(tt.args.podName, tt.args.namespace, tt.args.containerName, tt.args.containerID, tt.args.logName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTaskLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("GetTaskLog() got = %v, want %v, diff: %v", got, tt.want, diff)
			}
		})
	}
}

func TestTemplateLogPlugin_NewTaskLog(t *testing.T) {
	type fields struct {
		templateUri   string
		messageFormat core.TaskLog_MessageFormat
	}
	type args struct {
		input TaskLogInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    core.TaskLog
		wantErr bool
	}{
		{
			"splunk",
			fields{
				templateUri:   "https://{{.hostname}}/en-US/app/search/search?q=search%20index%3Diks%20kubernetes.namespace_name%3A%3A{{.namespace}}%20host%3D%22{{.podName}}%22&display.page.search.mode=verbose&dispatch.sample_ratio=1&earliest=-24h%40h&latest=now&sid=1596585766.253090_41E7FA1D-27F9-438E-850B-381E85A46081",
				messageFormat: core.TaskLog_JSON,
			},
			args{
				input: TaskLogInput{
					HostName:      "my-host",
					PodName:       "my-pod",
					Namespace:     "my-namespace",
					ContainerName: "ignore",
					ContainerID:   "ignore",
					LogName:       "main_logs",
				},
			},
			core.TaskLog{
				Uri:           "https://my-host/en-US/app/search/search?q=search%20index%3Diks%20kubernetes.namespace_name%3A%3Amy-namespace%20host%3D%22my-pod%22&display.page.search.mode=verbose&dispatch.sample_ratio=1&earliest=-24h%40h&latest=now&sid=1596585766.253090_41E7FA1D-27F9-438E-850B-381E85A46081",
				MessageFormat: core.TaskLog_JSON,
				Name:          "main_logs",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TemplateLogPlugin{
				templateUri:   tt.fields.templateUri,
				messageFormat: tt.fields.messageFormat,
			}
			got, err := s.NewTaskLog(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTaskLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskLog() got = %v, want %v", got, tt.want)
			}
		})
	}
}
