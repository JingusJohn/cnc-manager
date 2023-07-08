FROM node

WORKDIR /dockerBuild

COPY . .

RUN npm install

EXPOSE 3000

RUN npm run build

ENTRYPOINT [ "node", "build" ]
