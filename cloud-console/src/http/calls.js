const URL = "http://localhost"
const PORT = "8090"

const ws = new WebSocket("ws://localhost:8090/ws");
//Triggered when the connection is opened
ws.onopen = function (evt) {
    console.log("Connection open...");
    // ws.send("Hello WebSockets!");
    setInterval(() => {
        ws.send("Hello WebSockets! " + (Math.random() * 1000));
    }, 1000);
};
//Triggered when a message is received
ws.onmessage = function (evt) {
    console.log("Received Message: " + evt.data);
};
//Triggered when the connection is closed
ws.onclose = function (evt) {
    console.log("Connection closed.");
};

const requestAllInstances = (ID) => {
    return new Promise(async (resolve, reject) => {
        fetch(`${URL}:${PORT}/instance/read`, {
            method: 'post',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ ID }),
        }).then((resonse) => resonse.json())
            .then((resonse) => {
                console.log("instance/read", resonse);
                resolve(resonse)
            });
    })
}
const requestAllBuckets = (ID) => {
    return new Promise(async (resolve, reject) => {
        fetch(`${URL}:${PORT}/bucket/read`, {
            method: 'post',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ ID }),
        }).then((resonse) => resonse.json())
            .then((resonse) => {
                console.log("bucket/read", resonse);
                resolve(resonse)
            });
    })
}

const HTTPCalls = {
    requestAllInstances,
    requestAllBuckets,
}

export default HTTPCalls