# Jeu de la vie
Pour réussir ce défi tu dois développer le celèbre jeu de la vie de Conway. Le jeu devra être fonctionnel, la map fera 100 cellules par 100, on doit pouvoir ajouter ou supprimer une cellule n'importe ou, le déroulement se fera frame après frame, tu peux utiliser un terminal pour toute les commandes mais une interface graphique est conseillée. Pour la génération de vaisseaux : A travers un boutton ou une commande, l'utilisateur peut chosir des formes pré-concu par le développeur et les exécuter dans le jeu

## 💾 Language(s) utilisé(s)
Go

## 📦 Dépendances
- [Go 1.23.1](https://go.dev/)
- [ncurses](https://en.wikipedia.org/wiki/Ncurses)
<details>
<summary>Détails pour ncurses</summary>

Debian/Ubuntu :
`apt install ncurses-dev` \
Arch :
`pacman -S ncurses` \
Windows :
[Bonne chance](https://github.com/rthornton128/goncurses/wiki)
	 
</details>

## 🏃 Run
```
go run main.go
```

## 🎮 Utilisation

<kdb>q</kdb> : quitter \
<kdb>r</kdb> : générer un jeu aléatoire \
<kdb>c</kdb> : reinitialiser le jeu à zéro \
<kdb>p</kdb> : mettre en pause 

### Quand le jeu est en marche
<kdb>h</kdb><kdb>j</kdb><kdb>k</kdb><kdb>l</kdb> : bouger la caméra \
<kdb>s</kdb> : changer la vitesse 

### Quand le jeu est en pause
<kdb>h</kdb><kdb>j</kdb><kdb>k</kdb><kdb>l</kdb> : bouger la caméra \
<kdb>s</kdb> : avancer d'un tour \
<kdb>t</kdb> : activer le curseur 

### Quand le curseur est activé
<kdb>h</kdb><kdb>j</kdb><kdb>k</kdb><kdb>l</kdb> : bouger le curseur \
<kdb>SPACE</kdb> : allumer/éteindre la cellule séléctionée


## ✏️ Note
j'ai 0 idée de si ça marche sur windows, y a surement moyen d'installer les lib ncurses, mais le plus simple c'est surement juste de tourner ça sur wsl
j'avais deja écrit un jeu de la vie en go donc ça aide quand même
oui j'utilise le vim motion pour la caméra et le curseur 