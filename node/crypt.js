
const ecies = require('eciesjs')
const encoding = require('@cosmjs/encoding')

const [mode, payload, key] = process.argv.slice(2)

if (!mode || !payload || !key) {
  throw new Error("One of mandatory arguments is no specified")
}

const decode = (input) => {
  // Replace non-url compatible chars with base64 standard chars
  input = input
      .replace(/-/g, '+')
      .replace(/_/g, '/');

  // Pad out with standard base64 required padding characters
  const pad = input.length % 4;
  if(pad) {
    if(pad === 1) {
      throw new Error('InvalidLengthError: Input base64url string is the wrong length to determine padding');
    }
    input += new Array(5-pad).join('=');
  }

  return input;
}

switch (mode) {
  case 'encrypt':
    console.log(
      encoding.toBase64(
        ecies.encrypt(
          ecies.PublicKey.fromHex(key).compressed, 
          Buffer.from(encoding.fromBase64(decode(payload)))
        )
      )
    )
    break
  case 'decrypt':
    console.log(
      ecies.decrypt(
        ecies.PrivateKey.fromHex(key).toHex(), 
        Buffer.from(encoding.fromBase64(decode(payload)))
      ).toString()
    )
    break
}


