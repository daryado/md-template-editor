FROM node:lts-alpine as build-stage
ARG NODE_ENV
ARG API_AUTH_URL
ARG API_TEMPLATE_URL
ENV NODE_ENV "$NODE_ENV"
ENV API_AUTH_URL "$API_AUTH_URL"
ENV API_TEMPLATE_URL "$API_TEMPLATE_URL"

WORKDIR /app
COPY ["package.json", "package-lock.json*", "./"]
RUN npm install

COPY . .
RUN npm run build

FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY --from=build-stage /app/replace_api_url.sh /


RUN chmod +x /replace_api_url.sh
EXPOSE 80
CMD ["sh", "replace_api_url.sh"]
