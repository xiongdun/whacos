package m_task

type Task struct {
	JobName        string `json:"jobName"`
	CronExpression string `json:"cronExpression"`
	ModelName      string `json:"modelName"`
	MethodName     string `json:"methodName"`
	ConcurrentFlag string `json:"concurrentFlag"`
	JobStatus      string `json:"jobStatus"`
	JobGroup       string `json:"jobGroup"`
	Remarks        string `json:"remarks"`
}
