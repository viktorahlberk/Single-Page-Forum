export const pages = {

  login:
    `<section class="vh-100" style="background-color: #eee;">
    <div class="container h-100">
        <div class="row d-flex justify-content-center align-items-center h-100">
          <div class="col-lg-12 col-xl-11">
            <div class="card text-black" style="border-radius: 25px;">
              <div class="card-body p-md-5">
                <div class="row justify-content-center">
                  <div class="col-md-10 col-lg-6 col-xl-5 order-2 order-lg-1">
                    <p class="text-center h1 fw-bold mb-5 mx-1 mx-md-4 mt-4">Single Page App</p>

                    

                        <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                                <input type="text" id="login" class="form-control" placeholder="login or Email" name="login" required/>
                            </div>
                        </div>
                        <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                                <input type="password" id="password" class="form-control" placeholder="Password" name="password"required/>
                            </div>
                        </div>
                        <div class="d-flex justify-content-center mx-4 mb-3 mb-lg-4">
                            <button type="submit" id="loginBtn" class="btn btn-primary btn-lg">Submit</button>
                          </div>
                          <h4>Dont have account? <span id="reg-btn"><u>Register</u></span></h4>
                    
                  </div>
                  <div class="col-md-10 col-lg-6 col-xl-7 d-flex align-items-center order-1 order-lg-2">

                    <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-registration/draw1.webp" class="img-fluid" alt="Sample image">
    
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
    </div>
 
</section>
              
     
    `,
  header:
    `<div class="navbar navbar-light bg-light shadow fixed-top">       
            <a id="logoBtn" class="navbar-brand font-weight-bold">Coding Forum</a>
            <a id="greetings"></a>          
            <a id="newPostBtn">New Post</a>                 
            <a id="logoutBtn" class="nav-link">Logout</a>
      </div>`,

  register:
    `  <section class="vh-100" style="background-color: #eee;">
    <div class="container h-100">
        <div class="row d-flex justify-content-center align-items-center h-100">
          <div class="col-lg-12 col-xl-11">
            <div class="card text-black" style="border-radius: 25px;">
              <div class="card-body p-md-5">
                <div class="row justify-content-center">
                  <div class="col-md-10 col-lg-6 col-xl-5 order-2 order-lg-1">
                    <p class="text-center h1 fw-bold mb-5 mx-1 mx-md-4 mt-4">Sign up</p>

                    <form id="regForm" method="POST" action="/reg" class="mx-1 mx-md-4">

                        <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                                <input type="text" id="form3Example1c" class="form-control" placeholder="Nick Name"  name="nickname"/>
                            </div>
                        </div>
                        <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                                <input type="text" id="form3Example1c" class="form-control" placeholder="First Name" name="fname"/>
                               
                            </div>
                        </div>

                        <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                                <input type="text" id="form3Example1c" class="form-control" placeholder="Last Name" name="lname"/>
                            </div>
                        </div>
                        <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                                <input type="text" id="form3Example1c" class="form-control" placeholder="Age" name="age"/>
                            </div>
                        </div>
                        <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                                <input type="text" id="form3Example1c" class="form-control" placeholder="Gender" name="gender"/>
                            </div>
                        </div>
                        <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                                <input type="email" id="form3Example1c" class="form-control" placeholder="Email" name="email"/>
                            </div>
                        </div>
                        <div class="d-flex flex-row align-items-center mb-4">
                            <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                            <div class="form-outline flex-fill mb-0">
                                <input type="password" id="form3Example1c" class="form-control" placeholder="Password" name="passw"/>
                            </div>
                        </div>
                        <div class="d-flex justify-content-center mx-4 mb-3 mb-lg-4">
                            <button type="submit" class="btn btn-primary btn-lg">Register</button>
                          </div>
                  </form>
                          <h4 id="backToLoginBtn">Back to <u>Log in</u></h4>                    
                  </div>
                  <div class="col-md-10 col-lg-6 col-xl-7 d-flex align-items-center order-1 order-lg-2">

                    <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-registration/draw1.webp" class="img-fluid" alt="Sample image">
    
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
    </div>
 
</section>
              
    </div>`,
  newPost:
    `
    <main id="main">
    <div class="newPostContainer">
      <div class="container">
          <div class="row d-flex justify-content-center">
            <div class="col-lg-12 col-xl-11">
              <div class="card text-black" style="border-radius: 25px;">
                <div class="card-body ">
                  <div class="row justify-content-center">
                    <div class="col-md-10 col-lg-6 col-xl-5 order-2 order-lg-1">
                      <p class="text-center h1 fw-bold mb-5 mx-1 mx-md-4 mt-4">Create Post</p>

                      <form id="submitNewPostForm">
                          <div class="d-flex flex-row align-items-center mb-4">
                              <i class="fas fa-user fa-lg me-3 fa-fw"></i>
                              <div class="form-outline flex-fill mb-0">
                                  <input type="text" id="login" class="form-control" placeholder="Title.." name="title" required/>
                              </div>
                          </div>
                          <p class="font-italic">Choose Categories.</p>
                          <div class="form-check">
                            <input class="form-check-input" type="checkbox" value="sport " id="flexCheckDefault" name="cat1">
                            <label class="form-check-label" for="flexCheckDefault">
                              Sport
                            </label>
                          </div>
                          <div class="form-check">
                            <input class="form-check-input" type="checkbox" value="news " id="flexCheckChecked" name="cat2">
                            <label class="form-check-label" for="flexCheckChecked">
                             News
                            </label>
                          </div>    
                          <div class="form-check">
                            <input class="form-check-input" type="checkbox" value="art " id="flexCheckDefault" name="cat3">
                            <label class="form-check-label" for="flexCheckDefault">
                              Art
                            </label>
                          </div>
                          <div class="form-check">
                            <input class="form-check-input" type="checkbox" value="music " id="flexCheckChecked" name="cat4">
                            <label class="form-check-label" for="flexCheckChecked">
                              Music
                            </label>
                          </div>    
                          <div class="form-check">
                            <input class="form-check-input" type="checkbox" value="animals " id="flexCheckDefault"name="cat5">
                            <label class="form-check-label" for="flexCheckDefault">
                              Animals
                            </label>
                          </div>
                          <div class="form-check">
                            <input class="form-check-input" type="checkbox" value="games " id="flexCheckChecked" name="cat6">
                            <label class="form-check-label" for="flexCheckChecked">
                              Games
                            </label>
                          </div>
    
                          <div class="md-form mt-4">
                            <i class="fas fa-pencil-alt prefix"></i>
                            <textarea id="form10" class="md-textarea form-control" rows="3" placeholder="Write something here.." name="body" required></textarea>                           
                          </div>    
    
                          <div class="d-flex justify-content-center m-4">
                              <button type="submit" id="lBtn" class="btn btn-primary btn-md">Submit</button>
                          </div>
                      </form>

                  </div>
                </div>
              </div>
            </div>
          </div>
    </div>
    
    </main>`,
  minipost:
    ` <div class="card card-1">
    
                  <div class="card-header">
                    USask computer scientist studying e-sports gamers’ ability to succeed or fail
                  </div>
                  <div class="card-body">
                    <div class=" card-title justify-content-start ml-2"><span class="d-block font-weight-bold name"><h5>Games</h5></span></div>
                    <p class="card-text">With supporting text below as a natural lead-in to additional content.</p>
                    <div class="games">
                      <div class="bg-white p-2">
                        <div class="d-flex flex-row user-info">
                            <div class="d-flex flex-column justify-content-start ml-2"><span class="d-block font-weight-bold name">USask computer scientist studying e-sports gamers’ ability to succeed or fail</span></div>
                        </div>
                        <div class="mt-2">
                            <p class="comment-text">Competitive gaming has been steadily gaining popularity around the world over the years, as millions are playing games such as League of Legends and Fortnite.
    
    
                              Professional gamers pursue the grand title of e-sports champion and with that, in many cases, there are cash prizes as they compete in massive tournaments around the world.
                              But what makes a great player? Or one who is still seeking to become better? How do they acquire the necessary skill sets to succeed?
    
                            </p>
                        </div>
                    </div>
                    <div class="bg-white">
                        <div class="d-flex flex-row fs-12">
                            <div class="like p-2 cursor"><i class="fa fa-thumbs-o-up"></i><span class="ml-1">Like</span></div>
                            <div class="like p-2 cursor"><i class="fa fa-commenting-o"></i><span class="ml-1">Comment</span></div>
                            <div class="like p-2 cursor"><i class="fa fa-share"></i><span class="ml-1">Share</span></div>
                        </div>
                    </div>
                    <div class="bg-light p-2">
                        <div class="d-flex flex-row align-items-start"> <textarea class="form-control ml-1 shadow-none textarea"></textarea></div>
                        <div class="mt-2 text-right"><button class="btn btn-primary btn-sm shadow-none" type="button">Post comment</button></div>
                    </div>
                  </div>
                  </div>
                </div>
                `,
  postpage: `

  <main id="postpageMain">

  <div id="postpageContent">
  <div id="postAndComments">
    <div id="onePost">
      <div class="card card-1">
        <div class="card-body">
            <div class=" card-title justify-content-start ml-2"><span class="d-block font-weight-bold name">
                    <h5 class="post-title">Post title</h5>
                </span></div>
            <div class="games">
                <div class="bg-white p-2">
                    <p class="post-text"></p>
                </div>
                
                <form class="bg-light p-2">
                    <div class="d-flex flex-row align-items-start"> <textarea
                            class="form-control ml-1 shadow-none textarea" id="commentTextArea"></textarea></div>
                    <div class="mt-2 text-right"><button class="btn btn-primary btn-sm shadow-none" id="createCommentBtn"type="button">Post
                            comment</button></div>
                </form>
            </div>
        </div>
      </div>
    </div>

    <div id="comments"></div>
  </div>
  
  <div id="users"><div/>
  </div>
  </main>`,
  comment:
    `<div class="card card-1" style="margin-top: 85px;">
        <div class="card-body">
            <div class=" card-title justify-content-start ml-2"><span class="d-block font-weight-bold name">
                    <h5 class="post-title"></h5>
                </span></div>
            <div class="games">
                <div class="bg-white p-2">
                    <p id="commentBody" class="post-text">Comment-body</p>
                </div>
                <div class="bg-white">
                    <div class="d-flex flex-row fs-12">
                        <div class="like p-2 cursor"><i class="fa fa-commenting-o"></i>
                          <span class="ml-1">Author</span>
                        </div>
                </div>
            </div>

        </div>
    </div>`,
  chatpage:
    ` 
    <main id="main">
     

  <div class="chatAndUsers">
    <div id="chat">
      <h3>Connected with <span id="connectedWith"></span></h3>
      <div class="messagingArea" style="overflow:auto; height: 60vh"></div>
      <div id="sendMsg">
          <input id="input" type="text"/>
          <button id="sendBtn">Send</button>
        </div>
    </div>

    <div id="users"></div>
  </div>
    </main>
    
    `,

  mainpage: `
    <main id="main">
   
      <div class="postsAndUsers">
        <div id="posts"></div>
        <div id="users"></div>
      </div>
    </main>
    
    `
} 