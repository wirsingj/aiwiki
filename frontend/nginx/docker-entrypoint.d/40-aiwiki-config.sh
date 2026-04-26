#!/bin/sh
set -eu

api_base_url="${AIWIKI_API_BASE_URL:-}"
escaped_api_base_url="$(printf '%s' "$api_base_url" | sed 's/\\/\\\\/g; s/"/\\"/g')"

cat > /usr/share/nginx/html/config.js <<EOF
window.__AIWIKI_CONFIG__ = {
  API_BASE_URL: "${escaped_api_base_url}"
};
EOF
