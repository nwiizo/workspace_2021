# argo submit -n argo --watch https://raw.githubusercontent.com/argoproj/argo-workflows/master/examples/hello-world.yaml
# argo list -n argo
# argo get -n argo @latest
# argo logs -n argo @latest
# Download the binary
curl -sLO https://github.com/argoproj/argo/releases/download/v3.0.0-rc6/argo-linux-amd64.gz

# Unzip
gunzip argo-linux-amd64.gz

# Make binary executable
chmod +x argo-linux-amd64

# Move binary to path
mv ./argo-linux-amd64 /usr/local/bin/argo

# Test installation
argo version
