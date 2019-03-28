package 工厂模式

type Worker interface {
	Work(task *string)
}

type WorkerCreator interface {
	Create() Worker
}
