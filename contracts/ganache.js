const fs = require('fs');
const spawn = require('child_process').spawn;

const psKey = fs.readFileSync('../private.key').toString().replace(/\r?\n|\r/g, "");
const udKey = fs.readFileSync('../private.userdelegate.key').toString().replace(/\r?\n|\r/g, "");
const ganache = spawn('ganache-cli', ['--account', `0x${psKey},1000000000000000000000000`, '--account', `0x${udKey},1000000000000000000000000`]);

ganache.stdout.on('data', (data) => {
    console.log(data.toString());
});

ganache.stderr.on('data', (data) => {
    console.log(data.toString());
});
