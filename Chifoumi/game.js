const fullnames = {
	'r': "Requin",
	'm': "Méduse",
	'c': "Crevette",
}

let userWinCount = 0
let cpuWinCount = 0
let tieCount = 0
let totalGames = 0

function playMove(playerMove) {

	let cpuMove = "rmc"[Math.floor(Math.random() * 3)] // oui
	let winner

	switch (playerMove) {
		case cpuMove:
			winner = "noone"
			break
		case 'r':
			winner = (cpuMove === 'm') ? "cpu" : "user";
			break
		case 'm':
			winner = (cpuMove === 'c') ? "cpu" : "user";
			break
		case 'c':
			winner = (cpuMove === 'r') ? "cpu" : "user";
			break
	}

	switch (winner) {
		case "noone":
			document.getElementById("winner").textContent = "Égalité!"
			tieCount++
			break;
		case "user":
			document.getElementById("winner").textContent = "Vous avez gagné.e!"
			userWinCount++
			break;
		case "cpu":
			document.getElementById("winner").textContent = "La mer à gagné..."
			cpuWinCount++
			break;
		default:
			break;
	}

	totalGames++
	
	document.getElementById("userPlay").textContent = fullnames[playerMove]
	document.getElementById("cpuPlay").textContent = fullnames[cpuMove]

	document.getElementById("nbvictory").textContent = userWinCount
	document.getElementById("nbdefeat").textContent = cpuWinCount
	document.getElementById("nbtie").textContent = tieCount
	
	document.getElementById("winrate").textContent = ((userWinCount / totalGames) * 100).toFixed(2) + "%"

	document.getElementById("results").style.visibility = 'visible'
}
