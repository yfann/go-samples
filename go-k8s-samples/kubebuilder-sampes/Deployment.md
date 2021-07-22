```go
import (
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *CustomReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	// Get CustomResource
	var customResource myApi.CustomResource
	if err := r.Get(ctx, req.NamespacedName, &customResource); err != nil {
		if apierrs.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, errors.Wrap(err, "unable to fetch CustomResource")
	}

	// CreateOrUpdate SERVICE
	var svc corev1.Service
	svc.Name = customResource.Name
	svc.Namespace = customResource.Namespace
	_, err := ctrl.CreateOrUpdate(ctx, r, &svc, func() error {
		ModifyService(customResource, &svc)
		return controllerutil.SetControllerReference(&customResource, &svc, r.Scheme)
	})
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "unable to CreateOrUpdate Service")
	}

	// CreateOrUpdate DEPLOYMENT
	var app appsv1.Deployment
	app.Name = customResource.Name + "-app"
	app.Namespace = customResource.Namespace
	_, err = ctrl.CreateOrUpdate(ctx, r, &app, func() error {
		ModifyDeployment(customResource, &app)
		return controllerutil.SetControllerReference(&customResource, &app, r.Scheme)
	})
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "unable to CreateOrUpdate Deployment")
	}

	return ctrl.Result{}, nil
}

func ModifyDeployment(cr myApi.CustomResource, deployment *appsv1.Deployment) {
	labels := generateLabels(cr.Name)
	if deployment.Labels == nil {
		deployment.Labels = make(map[string]string)
	}
	for k, v := range labels {
		deployment.Labels[k] = v
	}
	replicas := cr.Spec.Replicas
	deployment.Spec.Replicas = &replicas
	deployment.Spec.Template.Labels = labels

	templateSpec := &deployment.Spec.Template.Spec

	if len(templateSpec.Containers) == 0 {
		templateSpec.Containers = make([]corev1.Container, 1)
	}
	container := &templateSpec.Containers[0]

	container.Name = "myapp"
	container.Args = []string{"/opt/myapp/bin/myapp"}
	container.Image = "myrepo/myapp:v1.0"
	container.Resources = corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    cr.Spec.CPU,
			corev1.ResourceMemory: cr.Spec.Memory,
		},
	}
}
```


## ref
+ [source](https://engineering.pivotal.io/post/gp4k-kubebuilder-lessons/)