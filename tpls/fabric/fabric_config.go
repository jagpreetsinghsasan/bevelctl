package fabric

const Fabric = `
##############################################################################################
#  Copyright Accenture. All Rights Reserved.
#
#  SPDX-License-Identifier: Apache-2.0
##############################################################################################

---
# yaml-language-server: $schema=../../../../platforms/network-schema.json
# This is a sample configuration file for testing Fabric deployment which has 3 nodes.
network:
  # Network level configuration specifies the attributes required for each organization
  # to join an existing network.
  type: fabric
  version: 2.2.2                 # currently tested 1.4.4, 1.4.8 and 2.2.2

  #Environment section for Kubernetes setup
  env:
    type: "local"             # tag for the environment. Important to run multiple flux on single cluster
    proxy: none               # 'none' can be used in single-cluster networks to don´t use a proxy
    ambassadorPorts:          # Any additional Ambassador ports can be given here, this is valid only if proxy='ambassador'
      portRange:              # For a range of ports 
        from: 15010 
        to: 15020
    # ports: 15020,15021      # For specific ports
    retry_count: 50                 # Retry count for the checks
    external_dns: disabled          # Should be enabled if using external-dns for automatic route configuration
    annotations:              # Additional annotations that can be used for some pods (ca, ca-tools, orderer and peer nodes)
      service: 
       - example1: example2
      deployment: {} 
      pvc: {}

  # Docker registry details where images are stored. This will be used to create k8s secrets
  # Please ensure all required images are built and stored in this registry.
  # Do not check-in docker_password.
  docker:
    url: "index.docker.io/hyperledgerlabs"
    username: "docker_username"
    password: "docker_password"

  # Remote connection information for orderer (will be blank or removed for orderer hosting organization)
  # For RAFT consensus, have odd number (2n+1) of orderers for consensus agreement to have a majority.
  orderers:
    - orderer:
{{ $IncrementedOrdererCount := .OrdererCount | add1 | int }}
{{ range $i, $e := untilStep 1 $IncrementedOrdererCount 1 }}
      type: orderer
      name: orderer{{ $e }}
      org_name: supplychain                 # org_name should match one organization definition below in organizations: key            
      uri: orderer{{ $e }}.supplychain-net:7050    # Internal URI for orderer which should be reachable by all peers
      certificate: /home/bevel/build/orderer{{ $e }}.crt   # the directory should be writable 
{{ end }}
`
