package templates

const nativePlayer = `<!DOCTYPE html>
<html>
<head>
<title>Example Web App</title>
<style>
    body {
        width: 100%;
        height: 100%;
        background-color:#202020;
    }
</style>

</head>
<body>
    <div>
        <video src="{{ . }}" autoplay>
        </video>
    </div>
</body>
</html>`

const shaka = `<!DOCTYPE html>
<html>
<head>
<title>ShakaPlayer Web App</title>
<script src="https://ajax.googleapis.com/ajax/libs/shaka-player/3.3.0/shaka-player.ui.js"></script>
<link rel="stylesheet" href="https://ajax.googleapis.com/ajax/libs/shaka-player/3.3.0/controls.css">
<script>
	const manifestUri =
    '{{ . }}';

function initApp() {
  // Install built-in polyfills to patch browser incompatibilities.
  shaka.polyfill.installAll();

  // Check to see if the browser supports the basic APIs Shaka needs.
  if (shaka.Player.isBrowserSupported()) {
    // Everything looks good!
    initPlayer();
  } else {
    // This browser does not have the minimum set of APIs we need.
    console.error('Browser not supported!');
  }
}

async function initPlayer() {
  // Create a Player instance.
  const video = document.getElementById('video');
  const player = new shaka.Player(video);

  // Attach player to the window to make it easy to access in the JS console.
  window.player = player;

  // Listen for error events.
  player.addEventListener('error', onErrorEvent);

  // Try to load a manifest.
  // This is an asynchronous process.
  try {
    await player.load(manifestUri);
    // This runs if the asynchronous load is successful.
    console.log('The video has now been loaded!');
  } catch (e) {
    // onError is executed if the asynchronous load fails.
    onError(e);
  }
}

function onErrorEvent(event) {
  // Extract the shaka.util.Error object from the event.
  onError(event.detail);
}

function onError(error) {
  // Log the error.
  console.error('Error code', error.code, 'object', error);
}

document.addEventListener('DOMContentLoaded', initApp);
</script>
<style>
    body {
        width: 100%;
        height: 100%;
        background-color:#202020;
    }
</style>

</head>
<body>
    <div>
       <video id="video"
           width="640"
           poster="//shaka-player-demo.appspot.com/assets/poster.jpg"
           controls autoplay></video>
    </div>
</body>
</html>`
