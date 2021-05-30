Proxy
Cert

## Goals
- Manage *** and dependent resources using Gateway API

## Definitions
- A *** instance or an instance of *** defines a fully functional *** deployment that runs on a k8s cluster

### Gateway API 
  open source project that exposes more general API for proxies used for more protocols than just HTTP; models more infrastructure component provides better deployment and management

### BackGround
  ** Operator manages

  
###### Design IDEA
  *** operator will manage ** using the GatewayClass, Gateway, and ** custom resources. Gateway API xRoute resourcs, i.e HTTPRoute, TCPRoute, TLSRoute, will be manged by ** and not ** Operator. The Gateway API targets different user personas that are responsible for defining and exposing an application in a kubernetes cluster.
  
  - platform provider : responsible for the overall environment the cluster runs in, i.e clouder provider. Platform provider interacts with Gateway Class and ** resources
  - platform operator:  responsible for cluster administration, manage network polcies, network access, interact with gateway resources
  - service operator: responsible for defining application configuration, service composition, interact with xRoute resources and generic k8s resources

  The Gateway API provides extension-points to linking for resources at different level.

Penetration Testing
Software Audit

Due to the open source nature, vulnerabilities are exposed and easy to located.
