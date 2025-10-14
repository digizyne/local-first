package dockerfiles

import (
	"fmt"
	"os"
)

func CreateDockerfile() error {
	const fileContent = `
	FROM node:24-alpine AS build
	WORKDIR /app
	COPY package*.json ./
	RUN npm install
	COPY . .
	RUN npm run build
	
	FROM nginx:alpine AS production
	COPY --from=build /app/dist /usr/share/nginx/html
	EXPOSE 80
	CMD ["nginx", "-g", "daemon off;"]
	`

	const fileName = "Dockerfile"

	err := os.WriteFile(fileName, []byte(fileContent), 0644)

	if err != nil {
		return fmt.Errorf("Failed to write to file %s: %v", fileName, err)
	}

	return nil
}

func CreateDockerComposeFile() error {
	const fileContent = `services:
	vue-app:
	  image: node:24
	  working_dir: /app
	  volumes:
	  - ./:/app
	  command: sh -c "npm install && npm run dev --host"
	`

	const fileName = "docker-compose.yml"

	err := os.WriteFile(fileName, []byte(fileContent), 0644)

	if err != nil {
		return fmt.Errorf("Failed to write to file %s: %v", fileName, err)
	}

	return nil
}
