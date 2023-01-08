import { LocalStorage } from "../Auth/LocalStorage";



const URL = "http://localhost"
const PORT = "8090"

// const ws = new WebSocket("ws://localhost:8090/ws");
//Triggered when the connection is opened
// ws.onopen = function (evt) {
//     console.log("Connection open...");
//     // ws.send("Hello WebSockets!");
//     setInterval(() => {
//         ws.send("Hello WebSockets! " + (Math.random() * 1000));
//     }, 1000);
// };
// //Triggered when a message is received
// ws.onmessage = function (evt) {
//     console.log("Received Message: " + evt.data);
// };
// //Triggered when the connection is closed
// ws.onclose = function (evt) {
//     console.log("Connection closed.");
// };

const requestAllInstances = (ID) => {
    const user = LocalStorage.get('user');
    return new Promise(async (resolve, reject) => {
        fetch(`${URL}:${PORT}/instance/read`, {
            method: 'POST',
            withCredentials: true,
            headers: { 'Content-Type': 'application/json', 'Authorization': user.token },
            body: JSON.stringify({ ID }),
        }).then((response) => {
            {
                if (response.ok) {
                    return response.json();
                }
                throw new Error('Something went wrong');
            }
        })
            .then((response) => {
                console.log("instance/read", response);
                resolve(response)
            }).catch(error => reject(error));
    })
}

const requestAllBuckets = (ID) => {
    const user = LocalStorage.get('user');
    return new Promise(async (resolve, reject) => {
        fetch(`${URL}:${PORT}/bucket/read`, {
            method: 'POST',
            withCredentials: true,
            headers: { 'Content-Type': 'application/json', 'Authorization': user.token },
            body: JSON.stringify({ ID }),
        }).then((response) => {
            {
                if (response.ok) {
                    return response.json();
                }
                throw new Error('Something went wrong');
            }
        })
            .then((response) => {
                console.log("bucket/read", response);
                resolve(response)
            }).catch(error => reject(error));
    })
}
const requestAllNetworks = (ID) => {
    const user = LocalStorage.get('user');
    return new Promise(async (resolve, reject) => {
        fetch(`${URL}:${PORT}/network/read`, {
            method: 'POST',
            withCredentials: true,
            headers: { 'Content-Type': 'application/json', 'Authorization': user.token },
            body: JSON.stringify({ ID }),
        }).then((response) => {
            {
                if (response.ok) {
                    return response.json();
                }
                throw new Error('Something went wrong');
            }
        })
            .then((response) => {
                console.log("bucket/read", response);
                resolve(response)
            }).catch(error => reject(error));
    })
}

const requestLogin = (loginBody) => {
    const user = LocalStorage.get('user');
    return new Promise(async (resolve, reject) => {
        fetch(`${URL}:${PORT}/signin`, {
            method: 'POST',
            withCredentials: true,
            headers: { 'Content-Type': 'application/json', 'Authorization': user.token },
            body: JSON.stringify(loginBody),
        }).then((response) => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Something went wrong');
        })
            .then((response) => {
                console.log("signin", response);
                resolve(response)
            }).catch(error => reject(error));
    })
}

const HTTPCalls = {
    requestAllInstances,
    requestAllBuckets,
    requestAllNetworks,
    requestLogin
}

export default HTTPCalls