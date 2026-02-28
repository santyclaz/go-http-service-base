ARG GO_VERSION=1.25

#
# [Stage] dev
# For local devcontainer
#
FROM golang:${GO_VERSION} AS dev

ARG USERNAME=eng-user
RUN adduser $USERNAME

ARG TZ
ENV TZ="$TZ"

# For persisting bash history
RUN SNIPPET="export PROMPT_COMMAND='history -a' && export HISTFILE=/cmdhistory/.bash_history" \
	&& mkdir /cmdhistory \
	&& touch /cmdhistory/.bash_history \
	&& chown -R $USERNAME /cmdhistory \
	&& echo "$SNIPPET" >> "/home/$USERNAME/.bashrc"

WORKDIR "/workspace"

#
# [Stage] dev-claude
# For local devcontainer with claude included
#
FROM dev AS dev-claude

USER $USERNAME

# Install Claude
ARG CLAUDE_CODE_VERSION=latest
RUN curl -fsSL https://claude.ai/install.sh | bash -s ${CLAUDE_CODE_VERSION}

# For persisting claude config
RUN mkdir -p /home/$USERNAME/.claude

#
# [Stage] build
# For building image
#
FROM golang:${GO_VERSION} AS build

WORKDIR "/workspace"
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the static binary, explicitly setting CGO_ENABLED=0 for compatibility with minimal images like scratch or alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o go-service .

#
# [Stage] final image
#
FROM scratch AS final

# Copy the compiled binary from the build stage
COPY --from=build /workspace/go-service /go-service

EXPOSE 8080

# Set the entry point to run the application
ENTRYPOINT ["/go-service"]
