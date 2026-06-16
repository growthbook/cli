#!/usr/bin/env node
'use strict';
// postinstall: download the release archive for this platform, verify its
// sha256 against checksums.txt, and extract the binary into vendor/.
// Zero runtime deps — uses Node built-ins plus the system `tar` (bsdtar ships
// on Linux, macOS, and Windows 10+, and extracts both .tar.gz and .zip).
const fs = require('fs');
const os = require('os');
const path = require('path');
const https = require('https');
const crypto = require('crypto');
const { execFileSync } = require('child_process');
const { assetInfo, vendorBinaryPath } = require('./lib/platform');

const REPO = 'growthbook/cli';
const VERSION = require('./package.json').version;
const TAG = `v${VERSION}`;
const BASE = `https://github.com/${REPO}/releases/download/${TAG}`;

function get(url) {
  return new Promise((resolve, reject) => {
    https
      .get(url, { headers: { 'User-Agent': 'growthbook-cli-installer' } }, (res) => {
        if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
          res.resume();
          resolve(get(res.headers.location));
          return;
        }
        if (res.statusCode !== 200) {
          res.resume();
          reject(new Error(`GET ${url} failed: HTTP ${res.statusCode}`));
          return;
        }
        const chunks = [];
        res.on('data', (c) => chunks.push(c));
        res.on('end', () => resolve(Buffer.concat(chunks)));
      })
      .on('error', reject);
  });
}

function sha256(buf) {
  return crypto.createHash('sha256').update(buf).digest('hex');
}

async function main() {
  const { archiveName, binaryName } = assetInfo();
  const dest = vendorBinaryPath();

  // Allow opting out of the network download (e.g. air-gapped CI that vendors
  // the binary another way) and skip work if it's already present.
  if (process.env.GROWTHBOOK_CLI_SKIP_DOWNLOAD) return;
  if (fs.existsSync(dest)) return;

  console.error(`growthbook: downloading ${archiveName} (${TAG})...`);
  const [archive, checksums] = await Promise.all([
    get(`${BASE}/${archiveName}`),
    get(`${BASE}/checksums.txt`),
  ]);

  const expected = checksums
    .toString('utf8')
    .split('\n')
    .map((l) => l.trim().split(/\s+/))
    .find(([, name]) => name === archiveName);
  if (!expected) throw new Error(`checksums.txt has no entry for ${archiveName}`);
  const got = sha256(archive);
  if (got !== expected[0]) {
    throw new Error(`checksum mismatch for ${archiveName}: expected ${expected[0]}, got ${got}`);
  }

  const tmp = fs.mkdtempSync(path.join(os.tmpdir(), 'growthbook-'));
  const archivePath = path.join(tmp, archiveName);
  fs.writeFileSync(archivePath, archive);

  const vendorDir = path.dirname(dest);
  fs.mkdirSync(vendorDir, { recursive: true });
  // bsdtar autodetects gzip and zip; -C extracts into vendorDir.
  execFileSync('tar', ['-xf', archivePath, '-C', vendorDir, binaryName], { stdio: 'inherit' });
  fs.rmSync(tmp, { recursive: true, force: true });

  if (!fs.existsSync(dest)) throw new Error(`extraction did not produce ${binaryName}`);
  if (process.platform !== 'win32') fs.chmodSync(dest, 0o755);
  console.error('growthbook: installed.');
}

main().catch((err) => {
  console.error(`\ngrowthbook: install failed: ${err.message}`);
  process.exit(1);
});
