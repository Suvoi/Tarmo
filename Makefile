# Configuration
VERSION=0.0.1
REPO_URL=git@github.com:Suvoi/Tarmo.git
IMAGE_NAME=ghcr.io/Suvoi/tarmo:$(VERSION)

.PHONY: release

release:
	@# Check branch
	@BRANCH=$$(git rev-parse --abbrev-ref HEAD); \
	if [ "$$BRANCH" != "develop" ]; then \
		echo "Error: you can only release from develop branch (current: $$BRANCH)"; \
		exit 1; \
	fi; \
	echo "✅ On develop branch, continuing..."

	@# Interactive commit
	@echo "Adding changes..."; \
	git add .; \
	read -p "Commit message: " MSG; \
	git commit -m "$$MSG"; \
	git push origin develop

	@# Docker build
	@echo "Building Docker image $(IMAGE_NAME)..."; \
	docker build -t $(IMAGE_NAME) .

	@# GHCR login
	@echo "Logging in to GHCR..."; \
	echo $${GHCR_TOKEN} | docker login ghcr.io -u Suvoi --password-stdin

	@# Docker push
	@echo "Pushing Docker image to GHCR..."; \
	docker push $(IMAGE_NAME); \
	echo "✅ Image pushed: $(IMAGE_NAME)"
