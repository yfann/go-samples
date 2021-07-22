```go
func main() {
  a, err := builder.SimpleController()
    // ReplicaSet is the Application type that
    // is Reconciled Respond to ReplicaSet events.
    ForType(&appsv1.ReplicaSet{}).
    // ReplicaSet creates Pods. Trigger
    // ReplicaSet Reconciles for Pod events.
    Owns(&corev1.Pod{}).
    // Call ReplicaSetController with the
    // Namespace / Name of the ReplicaSet
    Build(&ReplicaSetController{})
  if err != nil {
    log.Fatal(err)
  }
  log.Fatal(mrg.Start(signals.SetupSignalHandler()))
}

// ReplicaSetController is a simple Controller example implementation.
type ReplicaSetController struct {
  client.Client
}
```


```go
// InjectClient is called by the application.Builder
// to provide a client.Client
func (a *ReplicaSetController) InjectClient(
  c client.Client) error {
  a.Client = c
  return nil
}

// Reconcile reads the Pods for a ReplicaSet and writes
// the count back as an annotation
func (a *ReplicaSetController) Reconcile(
  req reconcile.Request) (reconcile.Result, error) {
  // Read the ReplicaSet
  rs := &appsv1.ReplicaSet{}
  err := a.Get(context.TODO(), req.NamespacedName, rs)
  if err != nil {
    return reconcile.Result{}, err
  }

  // List the Pods matching the PodTemplate Labels
  pods := &corev1.PodList{}
  err = a.List(context.TODO(),
    client.InNamespace(req.Namespace).
        MatchingLabels(rs.Spec.Template.Labels),
    pods)
  if err != nil {
    return reconcile.Result{}, err
  }

  // Update the ReplicaSet
  rs.Labels["selector-pod-count"] =
    fmt.Sprintf("%v", len(pods.Items))
  err = a.Update(context.TODO(), rs)
  if err != nil {
    return reconcile.Result{}, err
  }

  return reconcile.Result{}, nil
}
```


## ref
+ [source](kubebulder/docs/book/getting_started/hello_world.md)