#!/usr/bin/env node
'use strict';
// Launcher: exec the vendored native binary, forwarding args, stdio, and exit code.
const fs = require('fs');
const { spawnSync } = require('child_process');
const { vendorBinaryPath } = require('../lib/platform');

const bin = vendorBinaryPath();
if (!fs.existsSync(bin)) {
  console.error(
    'growthbook: native binary not found. Reinstall the package, or run the ' +
      'postinstall step manually: `node node_modules/growthbook/install.js`.'
  );
  process.exit(1);
}

const result = spawnSync(bin, process.argv.slice(2), { stdio: 'inherit' });
if (result.error) {
  console.error(`growthbook: ${result.error.message}`);
  process.exit(1);
}
process.exit(result.status === null ? 1 : result.status);
