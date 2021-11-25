cat <<'EOL' > /usr/local/bin/kubectl-get_all

#!/usr/bin/env bash

set -e -o pipefail; [[ -n "$DEBUG" ]] && set -x

exec kubectl get "$(kubectl api-resources --namespaced=true --verbs=list --output=name | tr "\n" "," | sed -e 's/,$//')" "$@"
EOL
chmod +x /usr/local/bin/kubectl-get_all
