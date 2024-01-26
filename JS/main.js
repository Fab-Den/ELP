const readline = require('readline');

const readline_interface = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

const draw_pile = {
    "A": 14,
    "B": 4,
    "C": 7,
    "D": 5,
    "E": 19,
    "F": 2,
    "G": 4,
    "H": 2,
    "I": 11,
    "J": 1,
    "K": 1,
    "L": 6,
    "M": 5,
    "N": 9,
    "O": 8,
    "P": 4,
    "Q": 1,
    "R": 10,
    "S": 7,
    "T": 9,
    "U": 8,
    "V": 2,
    "W": 1,
    "X": 1,
    "Y": 2,
    "Z": 1
};

const grids = [[], []]

const hands= [[], []]

const turn = 0

const player= 1

const end= false


function prompt(message, callback) {
    readline_interface.question(message, (a) => {
        callback(a);
    })
}



// Debut du tour
// 2 x Choix de Jarnac (pas le premier)
//  -> selectionne ligne sur la grille adverse
//  -> prompt mot + test si mot valide + modification de la grille adverse + sa grille + modification de la main
//  -> si entrer sans caractère -> back à la selection précédente
//
// Choix entre piocher une seule lettre ou echanger 3 lettres (pas la première prise de main)

// Suite du tour
// prompt choix entre
//  -> modifie un mot -> choisir une ligne / prompt le mot + tester si bonne taille + valide (pour plus tard) + modifier dans les grilles + enlever les lettres de la main
//  -> ajouter un mot -> rompt le mot + tester si bonne taille + valide (pour plus tard) + modifier dans les grilles + enlever les lettres de la main
//  -> passer (fin du tour)
// pour chaque choix -> test si fin de partie


// list functions
// -> compter les points
// -> tester la fin de partie
// -> piocher lettres + retirer de la pioche
// -> tester validité mot
// -> modification d'une ligne
// -> modification main (en paramètre les lettres retirées)
// ->

// ne pas oublier de fermer l'interface à la fin du jeu