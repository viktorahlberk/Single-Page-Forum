import { createMessageDiv } from "./app.js";
let alertTextDiv = `<div class="alert alert-warning alert-dismissible fade show" role="alert">
  <span id="alertMsg"></span>
  <button type="button" class="close" data-dismiss="alert" aria-label="Close">
    <span aria-hidden="true">&times;</span>
  </button>
</div>`
export function newWebSocketConnection(loggedInUserNickname) {
    let socket = new WebSocket("ws://localhost:8081/ws");

    socket.onopen = function () {
        console.log("Connection established with", loggedInUserNickname)
    };

    socket.onerror = function (e) {
        console.log("Socket error", e)
    }

    socket.onmessage = function (e) {
        let messageObj = JSON.parse(e.data);
        console.log("messageObj ->", messageObj)
        if (messageObj.userLeft) {
            // Update user list
            const disconnectedUserDiv = document.querySelector(`#${messageObj.userLeft}`)
            disconnectedUserDiv.classList.remove("onlineIndicator")
            return
        } else if (messageObj.userJoined) {
            const connectedUserDiv = document.querySelector(`#${messageObj.userJoined}`)
            connectedUserDiv.classList.add("onlineIndicator")
            return
        }


        let messagingAreaDiv = document.querySelector(".messagingArea");

        // If user hasnt opened chat we display an alert and do nothing more
        if (messagingAreaDiv === null) {
            sendChatNotification(messageObj.sender)
            return
        }

        // Get the name of the user who we opened chat with
        let chatOpenWith = document.querySelector("#connectedWith").textContent;

        // Checks if user scrollbar was scrolled (at the bottom)
        let scrollbarWasScrolled = false;
        if (Math.abs(messagingAreaDiv.scrollTop + messagingAreaDiv.clientHeight - messagingAreaDiv.scrollHeight) < 2) {
            scrollbarWasScrolled = true;
        }

        // Only display message if sender has opened a chat with reciever or
        // message reciver has opened the chat with message sender
        if (messageObj.sender === chatOpenWith || messageObj.reciever === chatOpenWith) {
            const messageDiv = createMessageDiv(messageObj, loggedInUserNickname);
            messagingAreaDiv.lastElementChild.append(messageDiv);
        } else {
            sendChatNotification(messageObj.sender)
            return
        }

        if (scrollbarWasScrolled) {
            // Places scrollbar to the bottom
            messagingAreaDiv.scrollTop = messagingAreaDiv.scrollHeight;
        }
    };

    return socket
}

function sendChatNotification(sender) {
    let main = document.querySelector("main");

    // Insert the alertDivText after the navbar
    main.insertAdjacentHTML("afterbegin", alertTextDiv)

    // Select and create message for alert
    let alertMsg = document.querySelector("#alertMsg")
    alertMsg.textContent = `New message from ${sender}`
}