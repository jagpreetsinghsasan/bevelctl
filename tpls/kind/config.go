package kind

const Kind = `
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
{{ $IncrementedControlPaneCount := .ControlPlaneCount | add1 | int }}
{{ range $i, $e := untilStep 1 $IncrementedControlPaneCount 1 }}
- role: control-plane
  image: kindest/node:v1.21.1@sha256:fae9a58f17f18f06aeac9772ca8b5ac680ebbed985e266f711d936e91d113bad
{{ end }}
{{ $IncrementedWorkerCount := .WorkerNodeCount | add1 | int }}
{{ range $i, $e := untilStep 1 $IncrementedWorkerCount 1 }}
- role: worker
  image: kindest/node:v1.21.1@sha256:fae9a58f17f18f06aeac9772ca8b5ac680ebbed985e266f711d936e91d113bad
{{ end }}
`
