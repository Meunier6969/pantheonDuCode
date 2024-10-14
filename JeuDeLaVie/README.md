# Jeu de la vie
Pour réussir ce défi tu dois développer le celèbre jeu de la vie de Conway. Le jeu devra être fonctionnel, la map fera 100 cellules par 100, on doit pouvoir ajouter ou supprimer une cellule n'importe ou, le déroulement se fera frame après frame, tu peux utiliser un terminal pour toute les commandes mais une interface graphique est conseillée. Pour la génération de vaisseaux : A travers un boutton ou une commande, l'utilisateur peut chosir des formes pré-concu par le développeur et les exécuter dans le jeu

## 💾 Language(s) utilisé(s)
Go

## 📦 Dépendances
- [Go 1.23.1](https://go.dev/)
- [ncurses](https://en.wikipedia.org/wiki/Ncurses)
<details>
<summary>Détails pour ncurses</summary>

**Debian/Ubuntu** : \
`apt install ncurses-dev` \
**Arch** : \
`pacman -S ncurses` \
**Windows** : \
[Bonne chance](https://github.com/rthornton128/goncurses/wiki)
	 
</details>

## 🏃 Run
```
cd JeuDeLaVie/
go run *.go
```

## 🎮 Utilisation
Ce jeu utilise les mouvement Vim pour se deplacer: \
`h` : droite \
`j` : bas \
`k` : haut \
`l` : gauche 

<details>
<summary>Contrôles</summary> 
	
`q` : quitter \
`r` : générer un jeu aléatoire \
`c` : reinitialiser le jeu à zéro \
`p` : mettre en pause 

### Quand le jeu est en marche
`hjkl` : bouger la caméra \
`s` : changer la vitesse 

### Quand le jeu est en pause
`hjkl` : bouger la caméra \
`s` : avancer d'un tour \
`t` : activer le curseur 

### Quand le curseur est activé
`hjkl` : bouger le curseur \
`az` : selectionner un pattern
`SPACE` : ajouter le pattern 
</details>

## ✏️ Note
j'ai 0 idée de si ça marche sur windows, y a surement moyen d'installer les lib ncurses, mais le plus simple c'est surement juste de tourner ça sur wsl
j'avais deja écrit un jeu de la vie en go donc ça aide quand même
oui j'utilise le vim motion pour la caméra et le curseur 
les patterns viennent de [la](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)
