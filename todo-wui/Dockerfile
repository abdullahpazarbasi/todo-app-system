FROM node:lts as builder

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# install project dependencies
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

# build app for production with minification
RUN npm run build

FROM nginx:stable as server

ENV HTTP_PORT 80
ENV VITE_TODO_WBFF_BASE_URL "http://localhost"

RUN mkdir -p /etc/nginx/templates && echo "server { listen \${HTTP_PORT}; root /usr/share/nginx/html; index index.html; location / { try_files \$uri \$uri/ /index.html; } location /status { access_log off; return 200 'OK'; add_header Content-Type text/plain; } error_page 500 502 503 504 /50x.html; location = /50x.html { root /usr/share/nginx/html; } }" > /etc/nginx/templates/default.conf.template

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE $HTTP_PORT

CMD [ "nginx", "-g", "daemon off;" ]
