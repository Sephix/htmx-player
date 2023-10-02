const playSvg =
  '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.348a1.125 1.125 0 010 1.971l-11.54 6.347a1.125 1.125 0 01-1.667-.985V5.653z" /></svg>';
const pauseSvg =
  '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M15.75 5.25v13.5m-7.5-13.5v13.5" /></svg>';
const likeSvg =
  '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12z" /></svg>';
const likedSvg =
  '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#f43f5e" class="w-6 h-6" ><path d="M11.645 20.91l-.007-.003-.022-.012a15.247 15.247 0 01-.383-.218 25.18 25.18 0 01-4.244-3.17C4.688 15.36 2.25 12.174 2.25 8.25 2.25 5.322 4.714 3 7.688 3A5.5 5.5 0 0112 5.052 5.5 5.5 0 0116.313 3c2.973 0 5.437 2.322 5.437 5.25 0 3.925-2.438 7.111-4.739 9.256a25.175 25.175 0 01-4.244 3.17 15.247 15.247 0 01-.383.219l-.022.012-.007.004-.003.001a.752.752 0 01-.704 0l-.003-.001z"/></svg>';

class WebPlayer extends HTMLElement {
  constructor() {
    super();
  }

  init = false;

  button;
  audio;
  progress;
  timestamp;
  duration;
  time;

  connectedCallback() {
    const disabled = this.getAttribute("disabled") == "true";
    const src = this.getAttribute("src");

    this.audio = document.createElement("audio");
    if (src) {
      this.audio.setAttribute("src", src);
      this.audio.addEventListener("timeupdate", this.playing);
      this.audio.addEventListener("ended", this.onEnded);
      this.audio.addEventListener("canplaythrough", this.onInit);
    }

    const outterButton = document.createElement("div");
    outterButton.setAttribute(
      "class",
      "flex flex-col justify-center  rounded-full ",
    );
    this.button = document.createElement("button");
    this.button.innerHTML = playSvg;
    this.button.addEventListener("click", this.playPauseAudio);
    if (disabled) {
      this.button.setAttribute("disabled", disabled);
      this.button.setAttribute("class", "cursor-not-allowed text-zinc-400");
    } else {
      outterButton.setAttribute(
        "class",
        "h-14, w-12 rounded-full hover:bg-zinc-300 flex justify-center items-center",
      );
    }
    outterButton.append(this.button);

    // ARTIST AND LIKE
    const songTitle = this.getAttribute("title");
    const artistName = this.getAttribute("artist");
    const firstLine = document.createElement("div");
    firstLine.setAttribute("class", "flex justify-between text-sm");
    if (songTitle && artistName) {
      const artist = document.createElement("span");
      artist.innerText = `${songTitle} - ${artistName}`;
      const favorite = document.createElement("div");
      if (this.getAttribute("liked") === "true") {
        favorite.innerHTML = likedSvg;
      } else {
        favorite.innerHTML = likeSvg;
      }
      favorite.setAttribute("id", "like-svg");
      favorite.setAttribute(
        "hx-get",
        "/track/like/" + this.getAttribute("song-id"),
      );
      favorite.setAttribute("hx-swap", "innerHTML");
      favorite.setAttribute("hx-trigger", "liking-update from:body");

      const favoriteContainer = document.createElement("div");
      favoriteContainer.setAttribute(
        "hx-put",
        "/track/like/current/" + this.getAttribute("song-id"),
      );
      favoriteContainer.setAttribute("hx-swap", "innerHTML");
      favoriteContainer.setAttribute("hx-trigger", "click");
      favoriteContainer.setAttribute("hx-target", "#like-svg");
      favoriteContainer.setAttribute("class", "cursor-pointer");
      favoriteContainer.append(favorite);

      firstLine.append(artist, favoriteContainer);
    }

    // TIME
    this.progress = document.createElement("input");
    this.progress.setAttribute("type", "range");
    this.progress.setAttribute("value", "0");
    this.progress.setAttribute("max", "0");
    this.progress.setAttribute("class", "grow");
    if (disabled) {
      this.progress.setAttribute("disabled", disabled);
    }
    this.progress.addEventListener("click", this.setCurrentTime);

    this.timestamp = document.createElement("span");
    this.timestamp.innerText = "00:00";
    this.duration = document.createElement("span");
    this.duration.innerText = "00:00";

    this.time = document.createElement("div");
    this.time.append(this.timestamp);
    this.time.append(this.progress);
    this.time.append(this.duration);
    this.time.setAttribute("class", "flex basis-3/4 gap-4 items-center");

    // INFO
    const info = document.createElement("div");
    info.setAttribute("class", "flex flex-col gap-2 basis-3/4 justify-center");
    info.append(firstLine, this.time);
    this.append(this.audio);
    this.append(outterButton);
    this.append(info);
  }

  playPauseAudio = () => {
    if (this.audio.paused) {
      this.audio.play();
      this.button.innerHTML = pauseSvg;
    } else {
      this.audio.pause();
      this.button.innerHTML = playSvg;
    }
  };

  playing = ({ target }) => {
    const { currentTime } = target;
    const seconds = Math.floor(currentTime);
    this.timestamp.innerText =
      "" +
      Math.floor(seconds / 60) +
      ":" +
      String(seconds % 60).padStart(2, "0");
    this.progress.value = currentTime;
  };

  setCurrentTime = (e) => {
    this.audio.currentTime = e.target.value;
  };

  onEnded = () => {
    this.button.innerHTML = playSvg;
  };

  onInit = () => {
    if (!this.init) {
      const duration = Math.floor(
        this.audio.duration === Infinity ? 30 : this.audio.duration,
      );
      this.progress.setAttribute("max", duration);
      this.duration.innerText =
        "" +
        Math.floor(duration / 60) +
        ":" +
        String(duration % 60).padStart(2, "0");

      if (this.getAttribute("autoplay") === "true") {
        this.playPauseAudio();
      }
      this.init = true;
    }
  };
}

customElements.define("web-player", WebPlayer);
