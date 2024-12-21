const socket = new WebSocket('ws://localhost:8080/ws');

socket.onopen = () => {
    console.log('Websocket conn established')
}

socket.onmessage = (event)=>{
    console.log('Received from server:', event.data)
    const responseContainer = document.getElementById("response");
    responseContainer.textContent = event.data;
}

socket.onerror = (err)=>{
    console.log('Websocket error:', err)
}

socket.onclose = () =>{
    console.log('Websocket conn closed')
}

var Button = document.getElementById("submit-btn")
Button.addEventListener("click", ()=>{
    const nameInput = document.getElementById("name")
    const name = nameInput.value
    if (name){
        socket.send(name)
        console.log('sent to server:', name)
    } else {
        console.log('Name cannot be empty')
    }
})
