ARG GO_VERSION=1.25

#
# [Stage] dev
# For local devcontainer
#
FROM golang:${GO_VERSION} AS dev

ARG USERNAME=vscode
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
