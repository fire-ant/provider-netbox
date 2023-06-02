allow_k8s_contexts('default')
load('ext://helm_remote', 'helm_remote')
load('ext://namespace', 'namespace_create')
watch_settings(ignore='demo/.*')


def netbox():
    hvals=[
        "postgresql.auth.postgresPassword=mypgsecret!",
        "postgresql.auth.Password=mypgsecret!",
        "redis.auth.password=myredissecret!",
        "image.tag=v3.5.2",
        "timeZone=GMT",
        "dateTimeFormat=H:i:s",
    ]
        # "superuser.name=admin",
        # "superuser.password=admin",
        # "superuser.apiToken=0123456789abcdef0123456789abcdef01234567"
    helm_remote('netbox',
        namespace="netbox-system",
        create_namespace=True,
        set=hvals,
        repo_name='bootc',
        version="4.0.1",
        repo_url='https://charts.boo.tc'
        )
    k8s_resource(workload='netbox',
        port_forwards='8888:8080',
        links=[
                'netbox',
                link('localhost:8888', 'http://localhost:8888')
            ]
            )
    k8s_yaml('patch/manifests/netbox-import-job.yaml')
    k8s_resource('netbox-import',resource_deps=['netbox'])


def zot():
    zvals=[]
    helm_remote('zot',
        namespace="zot-system",
        create_namespace=True,
        set=zvals,
        repo_name='project-zot',
        repo_url='http://zotregistry.io/helm-charts'
        )
    k8s_resource(workload='zot',
        port_forwards='5000:5000',
        links=[
                'zot',
                link('localhost:5000', 'zot portal ui')
            ]
            )
def crossplane():
    zvals=[]
    local_resource('crossplane',cmd='up uxp install -n upbound-system || true')
#   install crds
    p_resources = listdir("package/crds", recursive=True)
    for resource in p_resources:
        k8s_yaml(resource)
#  build and install provider
    local_resource(
        "build and install netbox-provider",
        cmd="make build local.xpkg.deploy.provider.provider-netbox",
        allow_parallel=True,
        resource_deps=['crossplane']
        )

def flux_install(
    base_components="source-controller,kustomize-controller,helm-controller",
    add_components="",
    all_namespaces=True):
    """Install Flux locally for GitOps-sync
    Args:
        all_namespaces: whether to watch all namespaces globally
        add_components:

    """
    if all_namespaces == "":
        fail("this must be set to a boolean value of true|false")
    if base_components == []:
        fail("this must be a list of strings, accepts comma-separated value")

    cmd = "flux install --components={base_components} --components-extra={add_components} --watch-all-namespaces={all_namespaces} ".format(
        base_components = base_components,
        add_components = add_components,
        all_namespaces = all_namespaces
    )
    namespace_create('flux-system')
    local(cmd)
def gitops_ui():
# PASSWORD=WWGit0ps!
    wgvalues=[
            "service.type=NodePort",
            "adminUser.create=true",
            "adminUser.username=admin",
            "adminUser.createClusterRole=true",
            "adminUser.createSecret=true",
            "adminUser.passwordHash=$2a$10$iBSzHfPmZzNYy1WDrC4WNu4OfJDcJfDNh801BwoqA2GgwqjNb9lsu"
            ]
    helm_remote('weave-gitops',
        namespace="flux-system",
        set=wgvalues,
        repo_name='ww-gitops',
        version="4.0.19",
        repo_url='https://helm.gitops.weave.works'
        )
    k8s_resource(workload='weave-gitops',
        port_forwards='9001:9001',
        links=[
                'weave-gitops',
                link('localhost:9001', 'weave-gitops ui')
            ]
        )
crossplane()
netbox()
#  test
# zot()
# flux_install()
# gitops_ui()