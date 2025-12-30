#!/bin/bash
set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <version>"
  echo "Example: $0 0.1.3"
  exit 1
fi

VERSION=$1
REPO="Doarakko/draw"
TMP_DIR=$(mktemp -d)

echo "Downloading release assets for v${VERSION}..."
gh release download "v${VERSION}" --repo "$REPO" --pattern "*.tar.gz" --dir "$TMP_DIR"

echo ""
echo "Calculating sha256..."
X86_64_DARWIN=$(shasum -a 256 "$TMP_DIR/draw-x86_64-apple-darwin.tar.gz" | awk '{print $1}')
AARCH64_DARWIN=$(shasum -a 256 "$TMP_DIR/draw-aarch64-apple-darwin.tar.gz" | awk '{print $1}')
X86_64_LINUX=$(shasum -a 256 "$TMP_DIR/draw-x86_64-unknown-linux-gnu.tar.gz" | awk '{print $1}')

echo "x86_64-apple-darwin:      $X86_64_DARWIN"
echo "aarch64-apple-darwin:     $AARCH64_DARWIN"
echo "x86_64-unknown-linux-gnu: $X86_64_LINUX"

echo ""
echo "Generating formula..."

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
OUTPUT_FILE="$SCRIPT_DIR/../Formula/draw.rb"
mkdir -p "$(dirname "$OUTPUT_FILE")"

cat > "$OUTPUT_FILE" << EOF
class Draw < Formula
  desc "Display Yu-Gi-Oh! card image on the terminal"
  homepage "https://github.com/Doarakko/draw"
  version "${VERSION}"
  license "MIT"

  on_macos do
    on_intel do
      url "https://github.com/Doarakko/draw/releases/download/v#{version}/draw-x86_64-apple-darwin.tar.gz"
      sha256 "${X86_64_DARWIN}"
    end

    on_arm do
      url "https://github.com/Doarakko/draw/releases/download/v#{version}/draw-aarch64-apple-darwin.tar.gz"
      sha256 "${AARCH64_DARWIN}"
    end
  end

  on_linux do
    on_intel do
      url "https://github.com/Doarakko/draw/releases/download/v#{version}/draw-x86_64-unknown-linux-gnu.tar.gz"
      sha256 "${X86_64_LINUX}"
    end
  end

  def install
    bin.install "draw"
  end

  test do
    assert_match "draw", shell_output("#{bin}/draw --help", 2)
  end
end
EOF

echo ""
echo "Generated: $OUTPUT_FILE"
cat "$OUTPUT_FILE"

echo ""
echo "Cleaning up..."
rm -rf "$TMP_DIR"

echo ""
echo "Done! Copy Formula/draw.rb to homebrew-tap/Formula/draw.rb"
