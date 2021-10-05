# csi-driver-crds

## Generate CRD

```bash
$ kubebuilder init --domain csi.appscode.com --skip-go-version-check
$ kubebuilder edit --multigroup=true
$ kubebuilder create api --group cacerts --version v1alpha1 --kind CAProviderClass
```

## Deploy and Run

```bash
$ kubectl apply -f ./config/crd/bases/cacerts.csi.appscode.com_caproviderclasses.yaml
customresourcedefinition.apiextensions.k8s.io/caproviderclasses.cacerts.csi.appscode.com created

$ make run
```
