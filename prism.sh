#!/bin/bash

echo "📦 Scaffolding universal entities into prism/..."

mkdir -p prism

create_entity() {
  file="prism/entity_$1.go"
  touch "$file"
  cat <<EOF > "$file"
package prism

type $2 struct {
  // TODO: Define fields for $2
}
EOF
}

# Entities
create_entity "user" "User"
create_entity "payout" "Payout"
create_entity "subscription" "Subscription"
create_entity "audit_log" "AuditLog"
create_entity "feature_flag" "FeatureFlag"
create_entity "session" "Session"

echo "✅ Done. Entities stubbed into prism/."
