const fullnames = {
	'p': "Pierre",
	'f': "Feuille",
	'c': "Ciseaux",
}

let userWinCount = 0
let cpuWinCount = 0
let tieCount = 0
let totalGames = 0

function playMove(playerMove) {

	let cpuMove = "pfc"[Math.floor(Math.random() * 3)] // oui
	let winner

	switch (playerMove) {
		case cpuMove:
			winner = "noone"
			break
		case 'p':
			winner = (cpuMove === 'f') ? "cpu" : "user";
			break
		case 'f':
			winner = (cpuMove === 'c') ? "cpu" : "user";
			break
		case 'c':
			winner = (cpuMove === 'p') ? "cpu" : "user";
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
			document.getElementById("winner").textContent = "L'ordinateur à gagné..."
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
	document.getElementById("totalgames").textContent = totalGames
	document.getElementById("winrate").textContent = ((userWinCount / totalGames) * 100).toFixed(2) + "%"

	document.getElementById("results").style.visibility = 'visible'
}
