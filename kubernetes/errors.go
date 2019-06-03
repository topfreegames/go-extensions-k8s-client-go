package kubernetes

type UnexpectedRestInterfaceImplError struct{}

func (e *UnexpectedRestInterfaceImplError) Error() string {
	return "Not a *rest.RESTClient or a *restWrapper.Client instance"
}

type NotClientsetError struct{}

func (e *NotClientsetError) Error() string {
	return "Not a *Clientset instance"
}
