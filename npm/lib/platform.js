'use strict';
// Shared platform/asset resolution for the install script and the launcher.
// Asset names must match .goreleaser.yaml's archive name_template:
//   growthbook_{title Os}_{x86_64|arm64}.{tar.gz|zip}
const path = require('path');

const OS_MAP = { linux: 'Linux', darwin: 'Darwin', win32: 'Windows' };
const ARCH_MAP = { x64: 'x86_64', arm64: 'arm64' };

function assetInfo() {
  const os = OS_MAP[process.platform];
  const arch = ARCH_MAP[process.arch];
  if (!os || !arch) {
    throw new Error(
      `Unsupported platform: ${process.platform}/${process.arch}. ` +
        `Prebuilt binaries exist for linux, darwin, windows on x64/arm64. ` +
        `Install from source instead: go install github.com/growthbook/cli/cmd/growthbook@latest`
    );
  }
  const isWindows = process.platform === 'win32';
  const ext = isWindows ? 'zip' : 'tar.gz';
  const binaryName = isWindows ? 'growthbook.exe' : 'growthbook';
  return {
    os,
    arch,
    ext,
    binaryName,
    archiveName: `growthbook_${os}_${arch}.${ext}`,
  };
}

function vendorBinaryPath() {
  return path.join(__dirname, '..', 'vendor', assetInfo().binaryName);
}

module.exports = { assetInfo, vendorBinaryPath };
