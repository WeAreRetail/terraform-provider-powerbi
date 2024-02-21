# Inspiration: https://github.com/devcontainers/templates/blob/main/src/go-postgres/.devcontainer/Dockerfile
ARG GO_VARIANT=1.21-bullseye
FROM mcr.microsoft.com/devcontainers/go:1-${GO_VARIANT}

ARG REMOTE_USER=vscode
ARG DELVE_VERSION=1.21.2

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Install additional packages
RUN apt-get update \
    && apt-get -y install --no-install-recommends procps lsb-release unzip gpg \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts

# Install Customizations
COPY customization ./customization
RUN bash ./customization/install_eza.sh \
    && bash ./customization/install_starship.sh \
    && bash ./customization/install_task.sh \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts && rm -rf ./customization

# Configure terraformrc
COPY terraform/.terraformrc /root/.terraformrc

# Install Go tools
RUN \
    # --> Delve for debugging
    go install github.com/go-delve/delve/cmd/dlv@v${DELVE_VERSION} \
    # --> Go language server
    && go install golang.org/x/tools/gopls@latest \
    # --> GolangCI-lint
    && curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sed 's/tar -/tar --no-same-owner -/g' | sh -s -- -b $(go env GOPATH)/bin \
    # --> Staticcheck
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    # --> GoReleaser
    && go install github.com/goreleaser/goreleaser@latest \
    #
    # Clean up
    && rm -rf /go/src/ && rm -rf /go/pkg

# Switching to the user that will be used in the Dev Container
USER ${REMOTE_USER}

COPY terraform/.terraformrc /tmp/.terraformrc

# Customize bash
RUN \
    # Add alias and PATH
    echo "alias tf=terraform" >> "$HOME/.bashrc" \
    && echo "alias ll=\"eza -la --icons\"" >> "$HOME/.bashrc" \
    && echo 'eval "$(starship init bash)"' >> "$HOME/.bashrc" \
    && echo "export PATH=$PATH:/usr/bin/"  >> "$HOME/.bashrc" \
    && cp "/tmp/.terraformrc" "$HOME/.terraformrc"

USER root
# Switch back to dialog for any ad-hoc use of apt-get
ENV DEBIAN_FRONTEND=dialog
