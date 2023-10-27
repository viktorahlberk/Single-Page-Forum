import { pages } from "../frontend/templates.js"
import { newWebSocketConnection } from "../frontend/websocket.js"

// TODO
// Check creating webSocketConnections when user logs in with email

let mouseDown = false;
let socket; // holds web socket connection



window.addEventListener("DOMContentLoaded", async () => {
  const authUserResponse = await authenticateUser();

  if (authUserResponse.authenticated) {
    goToMainPage();
    socket = newWebSocketConnection(authUserResponse.nickname)
  } else {
    goToLoginPage();
  }
})


function addGreetingsToHeader(username) {
  const div = document.getElementById('greetings')
  div.innerText = `Hi, ${username}`
}


async function goToNewPostPage() {
  const authUserResponse = await authenticateUser();
  const loggedInUserNickname = authUserResponse.nickname

  if (!authUserResponse.authenticated) {
    console.log("Not logged in!")
    goToLoginPage();
    return
  }

  document.body.innerHTML = pages.header + pages.newPost
  addGreetingsToHeader(loggedInUserNickname)



  const newPostContainer = document.querySelector('.newPostContainer')
  const userContainer = await createUserList(loggedInUserNickname);
  newPostContainer.appendChild(userContainer)

  let submitPostBtn = document.getElementById('lBtn')
  submitPostBtn.addEventListener('click', submitNewPost)
  setEventListeners()

}


async function createUserList(loggedInUserNickname) {

  const onlineUsers = await getOnlineUsers();
  console.log(onlineUsers)

  const userlistContainer = document.createElement('div')
  userlistContainer.id = 'userlistContainer'
  userlistContainer.className = 'userlistContainer'

  let otherUsers = await getOtherUsers(loggedInUserNickname)

  for (let user of otherUsers) {
    const userDiv = document.createElement('div')
    userDiv.id = user.NickName
    userDiv.innerHTML = `<div>${user.NickName}</div>`

    onlineUsers.forEach(onlineUser => {
      if (onlineUser.UserName === user.NickName) {
        userDiv.classList.add('onlineIndicator')
      }
    })

    userDiv.addEventListener("click", openChat)
    userlistContainer.appendChild(userDiv)
  }

  return userlistContainer
}
let loggedInUserNickname;
async function goToMainPage() {
  const authUserResponse = await authenticateUser();
  //let loggedInUserNickname;

  if (authUserResponse.authenticated) {
    loggedInUserNickname = authUserResponse.nickname
  } else {
    console.log("Not logged in!")
    goToLoginPage();
    return
  }

  document.body.innerHTML = pages.header + pages.mainpage
  addGreetingsToHeader(loggedInUserNickname)

  const mainContainer = document.querySelector('#main')


  // POSTS
  const posts = await getPosts()
  const postsContainer = document.querySelector('#posts')

  const div = document.createElement("div")
  for (let post of posts) {
    let datetime = convertTimestampForFrontend(post.Created)
    div.innerHTML +=
      `
          <div class="singlePost">
            <div class="postContainerHeader">   
              <h2 id="title" data-postId=${post.ID}>${post.Title}</h2>
            </div>
            <div class="postContainerFooter">
              <span>Created by ${post.AuthorName} at ${datetime}.</span>
            </div>
          </div>
        `
    postsContainer.appendChild(div)
  }

  // USER LIST
  const usersDiv = await createUserList(loggedInUserNickname, mainContainer);
  const usersContainer = document.querySelector("#users");
  usersContainer.appendChild(usersDiv)

  setEventListeners()

  // CREATE EVENT LISTENERS FOR ADDED POST TITLES
  const postTitles = document.querySelectorAll("#title");
  postTitles.forEach(postTitle => postTitle.addEventListener("click", async () => {
    const authUserResponse = await authenticateUser();

    if (!authUserResponse.authenticated) {
      console.log("Not logged in!")
      goToLoginPage();
      return
    }

    // ADD INITIAL HTML, GET POST ELEMENTS AND ID
    document.body.innerHTML = pages.header + pages.postpage;
    addGreetingsToHeader(loggedInUserNickname)
    const currentPostTitle = document.querySelector(".post-title");
    const currentPostBody = document.querySelector(".post-text");
    const postID = postTitle.dataset.postid;

    // GET CURRENT POST
    const currentPost = await getOnePost(postID)

    // SET CURRENT POST TITLE AND BODY
    currentPostTitle.textContent = currentPost.Title;
    currentPostBody.textContent = currentPost.Body;

    // GET ALL COMMENTS
    const postComments = await getComments(postID)

    // CHECK IF POST HAS ANY COMMENTS
    if (postComments.length !== 0) {
      addCommentsToHtml(postComments)
    }

    // GET USERS LIST
    const usersDiv = await createUserList(loggedInUserNickname, mainContainer);
    const usersContainer = document.querySelector("#users");
    usersContainer.appendChild(usersDiv)

    setEventListeners();

    // CREATE EVENT LISTENER FOR POST ADD COMMENT BUTTON
    const createCommentBtn = document.querySelector("#createCommentBtn");
    createCommentBtn.addEventListener("click", async () => {
      const authUserResponse = await authenticateUser();

      if (!authUserResponse.authenticated) {
        console.log("Not logged in!")
        goToLoginPage();
        return
      }

      // CREATE THE COMMENT (SEND IT TO SERVER)
      await createComment(Number(postID))

      // CHECK COMMENTS CONTAINER
      let commentsContainer = document.querySelector("#comments");
      if (commentsContainer === null) {
        commentsContainer = document.createElement("div");
        commentsContainer.id = "comments";
      } else {
        commentsContainer.innerHTML = "";
      }

      // GET ALL COMMENTS
      const postComments = await getComments(postID)
      addCommentsToHtml(postComments)
      setEventListeners();
    })
  }))
}

function goToRegPage() {
  document.body.innerHTML = pages.register
  let loginBack = document.getElementById('backToLoginBtn')
  loginBack.addEventListener('click', goToLoginPage)

  regForm.onsubmit = async (e) => {
    e.preventDefault();

    let response = await fetch('/reg', {
      method: 'POST',
      body: new FormData(regForm)
    });

    let result = await response.text();

    alert(result);
    if (result === "Registered") {
      goToLoginPage();
    } else {
      goToRegPage();
    }
  };
}

function goToLoginPage() {
  document.body.innerHTML = pages.login
  let regButton = document.getElementById('reg-btn')
  regButton.addEventListener('click', goToRegPage)
  let loginBtn = document.getElementById('loginBtn')
  loginBtn.addEventListener("click", async function () {

    let user = {
      login: document.getElementById('login').value,
      password: document.getElementById('password').value,
    };
    let response = await fetch('/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json;charset=utf-8'
      },
      body: JSON.stringify(user)
    });
    let result = await response.json();
    //console.log((result));
    if (result.IsAuthorised) {
      socket = newWebSocketConnection(user.login)
      goToMainPage();
    } else {
      goToLoginPage();
    }
  })
}

function setEventListeners() {

  let logoBtn = document.getElementById('logoBtn')
  logoBtn.addEventListener('click', goToMainPage)

  let btn = document.getElementById('logoutBtn')
  btn.addEventListener('click', logout)

  let postBtn = document.getElementById("newPostBtn")
  postBtn.addEventListener('click', goToNewPostPage)
}



function logout() {
  fetch('/logout')
    .then((response) => response.text())
    .then((value) => {
      console.log(value)
      socket.close(1000, "User logged out")
      goToLoginPage()

    })
    .catch(console.log)

}


function submitNewPost() {

  submitNewPostForm.onsubmit = async (e) => {
    e.preventDefault();

    let response = await fetch('/newpost', {
      method: 'POST',
      body: new FormData(submitNewPostForm)
    });

    let result = await response.text();

    alert(result);
    if (result === "Submitted") {
      goToMainPage();
    } else if (response.status === 400) {
      goToLoginPage();
    } else {
      goToNewPostPage();

    }
  };
}
function convertTimestampForFrontend(timestamp) {
  let date = new Date(timestamp)
  return date.toLocaleString()
}



export function createMessageDiv(message, loggedInUserNickname) {

  let messageDiv = document.createElement("div");
  let messageTitle = document.createElement("p");
  let messageText = document.createElement("p");
  let span = document.createElement("span");

  span.style.marginLeft = "0.25rem";

  messageText.textContent = message.content;
  messageTitle.textContent = message.sender;

  span.textContent = message.created;
  messageTitle.appendChild(span)


  messageDiv.appendChild(messageTitle)
  messageDiv.appendChild(messageText)
  if (message.sender === loggedInUserNickname) {
    messageDiv.style.backgroundColor = "gray";
  }

  return messageDiv
}

async function displayPreviousMessages(loggedInUserNickname, openedChatWith, messagesLoaded) {
  const messagingAreaDiv = document.querySelector(".messagingArea");
  const messagesDiv = document.createElement("div");
  messagesDiv.className = "messages";


  let previousMessages = await getLast10Messages(loggedInUserNickname, openedChatWith, messagesLoaded)

  for (let message of previousMessages) {
    let messageDiv = createMessageDiv(message, loggedInUserNickname)
    messagesDiv.appendChild(messageDiv)
  }

  messagingAreaDiv.prepend(messagesDiv)
}


async function openChat(e) {
  const authUserResponse = await authenticateUser();
  let loggedInUserNickname;
  let openedChatWith = e.target.textContent;

  if (authUserResponse.authenticated) {
    loggedInUserNickname = authUserResponse.nickname
  } else {
    console.log("Not logged in!")
    goToLoginPage();
    return
  }
  document.body.innerHTML = pages.header + pages.chatpage
  addGreetingsToHeader(loggedInUserNickname)

  let usersDiv = await createUserList(loggedInUserNickname)
  let usersContainer = document.querySelector("#users");
  usersContainer.appendChild(usersDiv)
  // const pageMainContent = document.querySelector('.chatAndUsers')
  // pageMainContent.appendChild(userListContainer)

  // CHAT BOX
  // Connected with
  var connectedWith = document.getElementById("connectedWith");
  connectedWith.textContent = openedChatWith;

  // Messaging area
  const messagingAreaDiv = document.querySelector(".messagingArea");
  messagingAreaDiv.messagesLoaded = 0;
  messagingAreaDiv.openedChatWith = openedChatWith;
  messagingAreaDiv.loggedInUserNickname = loggedInUserNickname

  // Messaging area event listeners
  messagingAreaDiv.addEventListener("scroll", debounce(loadMessages, 200))

  messagingAreaDiv.addEventListener("mousedown", () => {
    mouseDown = true;
  })
  messagingAreaDiv.addEventListener("mouseup", async () => {

    // If mouse button is released and messages were requested we will display them now
    if (messagingAreaDiv.requestMessages) {
      messagingAreaDiv.messagesLoaded += 10;
      let oldScrollHeight = messagingAreaDiv.scrollHeight

      await displayPreviousMessages(
        messagingAreaDiv.loggedInUserNickname,
        messagingAreaDiv.openedChatWith,
        messagingAreaDiv.messagesLoaded
      )

      messagingAreaDiv.scrollTop = messagingAreaDiv.scrollHeight - oldScrollHeight;
      messagingAreaDiv.requestMessages = false;
    }

    mouseDown = false;
  })

  await displayPreviousMessages(loggedInUserNickname, openedChatWith, messagingAreaDiv.messagesLoaded)
  messagingAreaDiv.scrollTop = messagingAreaDiv.scrollHeight; // Set scroll position at bottom

  setEventListeners()


  let sendBtn = document.getElementById('sendBtn')
  sendBtn.addEventListener('click', async function () {


    let messageObj = {
      sender: loggedInUserNickname,
      reciever: openedChatWith,
      content: input.value,
      created: new Date().toLocaleString()
    }
    socket.send(JSON.stringify(messageObj));
    input.value = "";

    usersDiv.remove()

    usersDiv = await createUserList(loggedInUserNickname)
    //usersContainer = document.querySelector("#users");
    usersContainer.appendChild(usersDiv)

    /*userListContainer = await createUserList(loggedInUserNickname)
    pageMainContent = document.getElementById('main')
    pageMainContent.appendChild(userListContainer)*/
  })

}

const debounce = (f, ms) => {
  let timeout;
  return function executedFunction() {
    const context = this;
    const args = arguments;
    const later = function () {
      timeout = null;
      f.apply(context, args);
    };
    clearTimeout(timeout);
    timeout = setTimeout(later, ms);
  };
};

async function loadMessages() {
  const messagingAreaDiv = document.querySelector(".messagingArea");

  // If scroll bar is moved to the top without mouse we display previous messages
  if (messagingAreaDiv.scrollTop === 0 && mouseDown == false) {

    this.messagesLoaded += 10;
    let oldScrollHeight = messagingAreaDiv.scrollHeight

    await displayPreviousMessages(
      messagingAreaDiv.loggedInUserNickname,
      messagingAreaDiv.openedChatWith,
      messagingAreaDiv.messagesLoaded
    )

    messagingAreaDiv.scrollTop = messagingAreaDiv.scrollHeight - oldScrollHeight;

  }

  // If scroll bar is moved to the top with mouse we will not display previous messages instantly
  if (messagingAreaDiv.scrollTop === 0 && mouseDown == true) {
    messagingAreaDiv.requestMessages = true;
  } else {
    messagingAreaDiv.requestMessages = false;
  }
  console.log("Scroll event!")

}

// ADDING HTML
function addCommentsToHtml(postComments) {
  const commentsContainer = document.querySelector("#comments");

  for (let comment of postComments) {
    let timestamp = convertTimestampForFrontend(comment.created)
    const commentDiv = document.createElement("div");
    commentDiv.id = "comment"
    commentDiv.innerHTML +=
      `          
              <p>${comment.body}</p>
              <div class="postContainerFooter">
              <span>Commented by ${comment.authorName} at ${timestamp}.</span>
              </div>                
      `
    commentsContainer.appendChild(commentDiv)
  }
}

async function getOnlineUsers() {
  try {
    const response = await fetch("/getonline")
    const data = await response.json();
    //console.log(data)
    return data

  } catch (e) {
    console.log(e)
    return e
  }
}

async function addUserlistToHtml() {
  const chatContainer = document.querySelector('.chatContainer')
  const userlistContainer = document.createElement('div')
  userlistContainer.id = 'userlistContainer'
  userlistContainer.className = 'userlistContainer'
  let otherUsers = await getOtherUsers(loggedInUserNickname)

  for (let user of otherUsers) {
    const userDiv = document.createElement('div')
    userDiv.innerHTML =
      `
      <div>${user.NickName}</div>
    `
    userDiv.addEventListener("click", openChat)
    userlistContainer.appendChild(userDiv)
  }
  chatContainer.appendChild(userlistContainer)
  document.body.appendChild(chatContainer)
}


// API REQUESTS
async function authenticateUser() {
  try {
    const response = await fetch("/userauth");
    const data = await response.json();
    return data

  } catch (e) {
    console.log(e)
  }
}

async function createComment(id) {
  const commentTextBox = document.querySelector("#commentTextArea");

  // If user tries to post empty comment - just return.
  if (commentTextBox.value === "") {
    return
  }

  try {
    await fetch(`/commentpost`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ "postid": id, "comment": commentTextBox.value })
    })
  } catch (e) {
    console.log(e)
    return e
  }
  commentTextBox.value = ""
}

async function getComments(postid) {
  try {
    const response = await fetch(`/getcomments?id=${postid}`)
    const data = await response.json();
    return data;

  } catch (e) {
    console.log(e)
    return e
  }
}

async function getPosts() {
  try {
    const response = await fetch("/allposts")
    const data = await response.json();
    return data

  } catch (e) {
    console.log(e)
    return e
  }
}

async function getOnePost(postid) {

  try {
    const response = await fetch(`/post?id=${postid}`)
    const data = await response.json();
    return data;

  } catch (e) {
    console.log(e)
    return e
  }
}


async function getAllUsers() {
  try {
    const response = await fetch("/getallusers")
    const data = await response.json();
    //console.log(data)
    return data

  } catch (e) {
    console.log(e)
    return e
  }
}

async function getOtherUsers(loggedInUser) {
  try {
    const response = await fetch(`/getotherusers?user=${loggedInUser}`)
    const data = await response.json();
    //console.log(data)
    return data

  } catch (e) {
    console.log(e)
    return e
  }
}

async function getLast10Messages(sender, reciever, messagesLoaded) {
  try {
    const response = await fetch(`/getlast10messages?sender=${sender}&reciever=${reciever}&messages_loaded=${messagesLoaded}`)
    const data = await response.json()
    return data
  } catch (e) {
    console.log(e)
  }
}

//function sortAlphabetically(){
//
//}
