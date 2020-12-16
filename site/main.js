function formatDuration(duration) {
	let minutes = (duration / 60 | 0);
	let seconds = (duration - minutes * 60) | 0;

	return `${minutes}` + ":" + `${seconds}`.padStart(2, "0");
}

var currentUrl = new URL(window.location.href);

var genreStr = currentUrl.searchParams.get("g");

var blockSeek = false;

if (!genreStr || genreStr == undefined || genreStr == "" || genreStr.length == 0) {
	// genreStr = "pop";
	window.location.replace("/genres");
}

var player = new Audio();

function download() {
	let url = player.src.replace('/get', '/download');
	window.open(url, '_blank');
}

async function getTrackInfo() {
	document.getElementById("trackTitle").innerHTML = "Загрузка...";
	document.getElementById("trackArtist").innerHTML = " ";

	let url = "/getRandomTrackInfoByGenre?g=" + genreStr;
	let response = await fetch(url);
	let commits = await response.json();

	document.getElementById("trackTitle").innerHTML = commits.Title;
	document.getElementById("trackArtist").innerHTML = commits.Artist;
	document.getElementById("genreList").innerHTML = commits.Genres.join(", ");
	player.src = "/get?url=" + commits.DownloadUrl;
}

async function nextTrack() {
	await getTrackInfo();
	await player.play();
}

player.addEventListener('playing', function () {
	document.querySelector("#playPauseButton").innerHTML = '<img src="/pause.svg">';
});

player.addEventListener('pause', function () {
	document.querySelector("#playPauseButton").innerHTML = '<img src="/play.svg">';
});

player.addEventListener('waiting', function () {
	document.querySelector("#playPauseButton").innerHTML = '<img src="/loading.svg">';
});

player.addEventListener('ended', function () {
	nextTrack();
});

player.addEventListener('error', function () {
	nextTrack();
});

player.addEventListener('durationchange', function () {
	document.querySelector("#myRange").max = player.duration | 0;
	document.getElementById("totalTime").innerHTML = formatDuration(player.duration);
})

player.addEventListener('timeupdate', function () {
	document.querySelector("#myRange").value = player.currentTime | 0;
	document.getElementById("currentTime").innerHTML = formatDuration(player.currentTime);
})

document.querySelector("#myRange").addEventListener('input', function () {
	if (!blockSeek) {
		blockSeek = true;
		//videoWasPaused = player.paused;
		player.pause();
	}
	player.currentTime = document.querySelector("#myRange").value;
}, false);

document.querySelector("#myRange").addEventListener('change', function () {
	//if (!videoWasPaused) {
	player.play();
	//}
	blockSeek = false;
}, false);

function playPauseTrack() {
	if (player.paused) {
		player.play();
	} else {
		player.pause()
	}
}

nextTrack();
