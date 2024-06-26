FROM node:18-alpine
WORKDIR /webapp/
COPY . .
RUN npm install && npm run build
CMD ["npm", "start"]
