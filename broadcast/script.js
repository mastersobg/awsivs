init();

async function init() {
    try {
        const streamConfig = IVSBroadcastClient.BASIC_LANDSCAPE;
        const client = IVSBroadcastClient.create({
            streamConfig: streamConfig,
        });
        window.client = client;
        const previewEl = document.getElementById('preview');
        client.attachPreview(previewEl);

        const devices = await navigator.mediaDevices.enumerateDevices();
        const videoDevices = devices.filter((d) => d.kind === 'videoinput');
        const audioDevices = devices.filter((d) => d.kind === 'audioinput');

        const cameraStream = await navigator.mediaDevices.getUserMedia({
            video: {
                deviceId: videoDevices[0].deviceId,
                width: {
                    ideal: streamConfig.maxResolution.width,
                    max: streamConfig.maxResolution.width,
                },
                height: {
                    ideal: streamConfig.maxResolution.height,
                    max: streamConfig.maxResolution.height,
                },
            },
        });
        client.addVideoInputDevice(cameraStream, 'camera1', {index: 0});

        const microphoneStream = await navigator.mediaDevices.getUserMedia({
            audio: {deviceId: audioDevices[0].deviceId},
        });
        client.addAudioInputDevice(microphoneStream, 'mic1');
    } catch (err) {
        setError(err.message);
    }
}

function setError(message) {
    if (Array.isArray(message)) {
        message = message.join("<br/>");
    }
    const errorEl = document.getElementById("error");
    errorEl.innerHTML = message;
}

async function startBroadcast() {
    const streamKeyEl = document.getElementById("stream-key");
    const endpointEl = document.getElementById("ingest-endpoint");
    const start = document.getElementById("start");

    try {
        start.disabled = true;
        await window.client.startBroadcast(streamKeyEl.value, endpointEl.value);
    } catch (err) {
        start.disabled = false;
        setError(err.toString());
    }
}

async function stopBroadcast() {
    const start = document.getElementById("start");
    start.disabled = false;
    try {
        await window.client.stopBroadcast();
    } catch (err) {
        setError(err.toString());
    }
}
