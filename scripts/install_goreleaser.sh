#!/usr/bin/env sh
# SPDX-License-Identifier: MIT
# Copyright (c) 2016-2025 Carlos Alexandro Becker
# Copyright (c) 2018 Márk Sági-Kazár <mark.sagikazar@gmail.com>
# Copyright (c) 2025 Geza Corp authors
set -e

readonly VERSION=${1:-}
readonly GIT_ROOT="$(git rev-parse --show-toplevel)"

if test "$DISTRIBUTION" = "pro"; then
	echo "Using Pro distribution..."
	RELEASES_URL="https://github.com/goreleaser/goreleaser-pro/releases"
	FILE_BASENAME="goreleaser-pro"
	LATEST="$(curl -sf https://goreleaser.com/static/latest-pro)"
else
	echo "Using the OSS distribution..."
	RELEASES_URL="https://github.com/goreleaser/goreleaser/releases"
	FILE_BASENAME="goreleaser"
	LATEST="$(curl -sf https://goreleaser.com/static/latest)"
fi

test -z "$VERSION" && VERSION="$LATEST"

test -z "$VERSION" && {
	echo "Unable to get goreleaser version." >&2
	exit 1
}

if [ -x ${GIT_ROOT}/bin/goreleaser-${VERSION} ]; then
    ln -sf goreleaser-${VERSION} bin/goreleaser
    echo "GoReleaser ${VERSION} is already installed"
    exit 0
fi

TMP_DIR="$(mktemp -d)"
# shellcheck disable=SC2064 # intentionally expands here
trap "rm -rf \"$TMP_DIR\"" EXIT INT TERM

OS="$(uname -s)"
ARCH="$(uname -m)"
test "$ARCH" = "aarch64" && ARCH="arm64"
TAR_FILE="${FILE_BASENAME}_${OS}_${ARCH}.tar.gz"

(
	cd "$TMP_DIR"
	echo "Downloading GoReleaser $VERSION..."
	curl -sfLO "$RELEASES_URL/download/$VERSION/$TAR_FILE"
	curl -sfLO "$RELEASES_URL/download/$VERSION/checksums.txt"
	echo "Verifying checksums..."
	sha256sum --ignore-missing --quiet --check checksums.txt
	if command -v cosign >/dev/null 2>&1; then
		echo "Verifying signatures..."
		REF="refs/tags/$VERSION"
		if test "$VERSION" = "nightly"; then
			REF="refs/heads/main"
		fi
		cosign verify-blob \
			--certificate-identity-regexp "https://github.com/goreleaser/goreleaser.*/.github/workflows/.*.yml@$REF" \
			--certificate-oidc-issuer 'https://token.actions.githubusercontent.com' \
			--cert "$RELEASES_URL/download/$VERSION/checksums.txt.pem" \
			--signature "$RELEASES_URL/download/$VERSION/checksums.txt.sig" \
			checksums.txt
	else
		echo "Could not verify signatures, cosign is not installed."
	fi
)

tar -xf "$TMP_DIR/$TAR_FILE" -C "$TMP_DIR"

mkdir -p ${GIT_ROOT}/bin bin
mv ${TMP_DIR}/goreleaser ${GIT_ROOT}/bin/goreleaser-${VERSION}
ln -sf goreleaser-${VERSION} bin/goreleaser

echo "GoReleaser ${VERSION} is installed"
