const chevronDown = `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5"/></svg>`;

class Playlist extends HTMLElement {
  page;
  button;

  connectedCallback() {
    this.page = document.createElement("dialog");

    this.page.classList.add(
      "appear",
      "fixed",
      "top-0",
      "left-0",
      "-z-10",
      "w-screen",
      "h-screen",
    );
    this.page.innerHTML = `<div class="flex flex-col"
      >
      <div class="flex justify-end items-center h-16 bg-zinc-950">
       <button id="playlist-collapse" class="flex justify-center items-center h-12 w-12 text-zinc-400 p-4 mr-8 hover:bg-zinc-700 cursor-pointer rounded-full">
         <span>${chevronDown}</span>
        </button>
      </div>
      <content
        hx-get="/playlist"
        hx-swap="outerHTML"
        hx-trigger="load"
      />
      </div>`;

    this.button = document.createElement("button");
    this.button.classList.add("flex", "items-center");
    this.button.innerHTML = `<div 
        hx-get="/playlist/preview"
        hx-swap="outerHTML"
        hx-trigger="load"
      />`;

    this.append(this.page);
    this.append(this.button);

    document
      .getElementById("playlist-collapse")
      .addEventListener("click", this.toggleDialog);
    this.button.addEventListener("click", this.toggleDialog);

    this.add;
  }

  toggleDialog = () => {
    if (!!this.page.getAttribute("open")) {
      this.page.removeAttribute("open");
      this.page.setAttribute("close", "true");
      document.body.classList.remove("overflow-y-hidden");
    } else {
      this.page.setAttribute("open", "true");
      this.page.removeAttribute("close");
      document.body.classList.add("overflow-y-hidden");
    }
  };
}

customElements.define("custom-playlist", Playlist);
