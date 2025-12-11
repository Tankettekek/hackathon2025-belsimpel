# Use a small nginx image
FROM nginx:alpine

# Remove default nginx static content
RUN rm -rf /usr/share/nginx/html/*

# Copy our static site into nginx's web root
COPY . /usr/share/nginx/html

# Expose port 80 (nginx default)
EXPOSE 80

# Use default nginx start command
CMD ["nginx", "-g", "daemon off;"]
