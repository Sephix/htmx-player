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

  playButton;
  nextButton;
  audio;
  progress;
  timestamp;
  duration;
  time;

  connectedCallback() {
    const src = this.getAttribute("src");
    if (src) {
      this.audio = document.createElement("audio");
      this.audio.setAttribute("src", src);
      this.audio.addEventListener("timeupdate", this.playing);
      this.audio.addEventListener("ended", this.onEnded);
      this.audio.addEventListener("canplaythrough", this.onInit);

      this.playButton = document.getElementById("player-play-button");
      this.playButton.addEventListener("click", this.playPauseAudio);

      this.nextButton = document.getElementById("player-next-button");
      this.nextButton.addEventListener("click", this.playNext);

      // TIME
      this.progress = document.getElementById("player-progress-bar");
      this.progress.addEventListener("click", this.setCurrentTime);

      this.timestamp = document.getElementById("player-progress-time");
      this.duration = document.getElementById("player-duration-time");
      // INFO
      this.append(this.audio);
    }
  }

  playPauseAudio = () => {
    if (this.audio.paused) {
      this.audio.play();
      this.playButton.innerHTML = pauseSvg;
    } else {
      this.audio.pause();
      this.playButton.innerHTML = playSvg;
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
    this.playButton.innerHTML = playSvg;
    this.playNext();
  };

  playNext = () => {
    fetch("/playlist/next", { method: "POST" }).then(() =>
      document.body.dispatchEvent(new Event("play-song")),
    );
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
