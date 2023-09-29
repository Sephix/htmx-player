const playSvg =
  '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.348a1.125 1.125 0 010 1.971l-11.54 6.347a1.125 1.125 0 01-1.667-.985V5.653z" /></svg>';
const pauseSvg =
  '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M15.75 5.25v13.5m-7.5-13.5v13.5" /></svg>';
const likeSvg =
  '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12z" /></svg>';
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
      favorite.innerHTML = likeSvg;
      firstLine.append(artist, favorite);
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
    console.log(currentTime);
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
