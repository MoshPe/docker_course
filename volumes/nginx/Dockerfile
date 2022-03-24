# Use an official Python runtime as a parent image
FROM node:latest
WORKDIR .
ADD . .
RUN npm install express body-parser cors mongoose nodemon
CMD ["node", "index.js"]
