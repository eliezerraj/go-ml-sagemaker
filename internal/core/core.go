package core

type InfoPod struct {
	PodName				string `json:"pod_name"`
	Version				string `json:"version"`
	OSPID				string `json:"os_pid"`
	IPAddress			string `json:"ip_address"`
	AvailabilityZone 	string `json:"availability_zone"`
	OtelExportEndpoint	string `json:"otel_export_endpoint"`
	ConfigOTEL			*ConfigOTEL 
}

type HttpAppServer struct {
	InfoPod 			*InfoPod 		`json:"info_pod"`
	Server     			*Server     		`json:"server"`
	SageMakerEndpoint 	*SageMakerEndpoint  `json:"sagemaker_endpoints"`
}

type Server struct {
	Port 			int `json:"port"`
	ReadTimeout		int `json:"readTimeout"`
	WriteTimeout	int `json:"writeTimeout"`
	IdleTimeout		int `json:"idleTimeout"`
	CtxTimeout		int `json:"ctxTimeout"`
}
