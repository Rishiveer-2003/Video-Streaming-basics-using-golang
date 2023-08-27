const videoPlayer = document.getElementById('videoPlayer');
videoPlayer.addEventListener('error', () =>{
    console.error('video playback error');
})

videoPlayer.addEventListener('ended', () =>{
    console.log('video ended');
})



