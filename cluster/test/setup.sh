#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Creating netbox credential secret..."

cat <<EOF | ${KUBECTL} apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: provider-secret
  namespace: upbound-system
type: Opaque
stringData:
  credentials: |
   {
      "server_url": "http://netbox.netbox-system.svc.cluster.local",
      "api_token": "0123456789abcdef0123456789abcdef01234567"
    }
EOF


echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: netbox.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: upbound-system
      key: credentials
EOF