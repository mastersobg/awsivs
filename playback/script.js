function setError(message) {
    const errorEl = document.getElementById("error");
    errorEl.innerHTML = message;
}


function attachEventListeners(player) {
    const PlayerEventType = IVSPlayer.PlayerEventType;
    player.addEventListener(PlayerEventType.ERROR, function (err) {
        setError(err.message);
    });
    player.addEventListener(PlayerEventType.TEXT_METADATA_CUE, function func(e) {
        const viewersEl = document.getElementById("viewers");
        viewersEl.innerHTML = e.text;
    });
}

function initURLInput() {
    const urlInput = document.getElementById("playback-url");
    urlInput.addEventListener("input", function (e) {
        setError("");
        player.load(e.target.value);
        player.play();
    }, true);
}

function createPlayer() {
    const player = IVSPlayer.create();
    player.attachHTMLVideoElement(document.getElementById('video-player'));
    return player;
}

function init() {
    if (!IVSPlayer.isPlayerSupported) {
        setError("Player is not supported by browser");
    }

    player = createPlayer();
    attachEventListeners(player);
    initURLInput();
}

init();
