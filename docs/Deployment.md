# Deployment

VPP Agent can run everywhere where VPP is installed. It can run either in VM or container.
 
Benefits of putting VPP to a container
 * simplifies: upgrade, start/top, potentially also scaling
 * introducing microservices takes advantage of small & reusable apps
 * supports container healing 
 
## K8s integration
Following diagram depics VPP deployed in:
- Data Plane vSwitch
- Control Plane vSwitch (TBD [Contive](http://contiv.github.io/) integration)
- VPP VNF Container
- Non VPP Container

![K8s integration](imgs/k8s_deployment.png "VPP Agent - K8s integration")

K8s:
- starts/stops the containers on multiple hosts
- checks containers health (using probes - HTTP calls)

## NB (Nort-bound) configuration vs. deployment
VPP Agent can be deployed to different environments. In following sub-chapters there are briefly 
described alternative deployments. Independent on the deployment the VPP Agent can be configured
using same Client v1 interface. There are three different implementations of the interface:
 - local client
 - remote client using Data Broker
 - remote client using GRPC

### Key Value Data Store for NB
The Control Plane using remote client writes configuration to the Data Store (tested with ETCD, Redis).
VPP Agent watches particular key prefixes in Data Store using dbsync package.

![deployment with data store](imgs/deployment_with_data_store.png)
TBD links to the code

### GRPC 
The Control Plane using remote client sends configuration to the Data Store (tested with ETCD, Redis).
VPP Agent watches particular key prefixes in Data Store using grpcsync package.

![grpc northbound](imgs/deployment_nb_grpc.png)
TBD links to the code

### Embeded deployment
VPP Agent can be embedded in different project. For integration with Contiv we use ebmedded deployment.
In this case VPP Agent get's the configuration from Local clivent v1 through in memory calls.

![embeded deployment](imgs/deployment_embeded.png)
TBD links to the code