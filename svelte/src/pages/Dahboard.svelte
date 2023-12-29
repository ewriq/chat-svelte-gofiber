<script>
  // @ts-nocheck

  export let users;
  console.log(users);

  import { io } from "socket.io-client";

  let socket;
  let message = "";
  let messages = [];

  socket = io("http://localhost:5000");

  socket.on("message", (message) => {
    messages = [...messages, message];
  });

  const sendMessage = () => {
    if (message != "") {
      socket.emit("message", message, users[0].username);
      message = "";
    } else {
      alert("Mesaj boş olmamas kardesm")
    }
  };
</script>

<main class="bg-slate-500">
  <div id="chat">
    <!-- svelte-ignore missing-declaration -->
    {#each messages as { user, text }, index}
      <p>{user}: {text}</p>
    {/each}
  </div>

  <div class=" roundex-xl   w-80" id="input-area">
    <input class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" bind:value={message} placeholder="Mesajınızı kışın" />
    <button class="text-white bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 font-medium rounded-full text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800" on:click={sendMessage}>Gönder</button>
  </div>

  <div class="">
    <h1>Kontol Paneli</h1>
    <h1>Username: {users[0].username}</h1>
    <h1>Password: {users[0].password}</h1>
    <h1>Email: {users[0].email}</h1>
  </div>
  
</main>

