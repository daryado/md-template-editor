#!/usr/bin/env sh

#/js
find '/usr/share/nginx/html' -name '*.js' -exec sed -i -e 's|VUE_APP_AUTH_SERVER_URL|'${API_AUTH_URL}'|g' {} \;
find '/usr/share/nginx/html' -name '*.js' -exec sed -i -e 's|VUE_APP_TEMPLATE_SERVER_URL|'${API_TEMPLATE_URL}'|g' {} \;
nginx -g "daemon off;"

