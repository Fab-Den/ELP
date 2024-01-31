const readline = require('readline');

const readline_interface = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let draw_pile = {
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

let grids = [[], [["A", "B", "C"]]]

let hands= [[], ["D"]]

let turn = 0;

let end= false

let number_jarnac = 0


function prompt(message, callback) {
    readline_interface.question(message, (a) => {
        callback(a);
    })
}


async function check_yes_no(str, callback_if_yes, callback_if_no, callback_if_none){
    return new Promise((resolve, reject) => {
        if (str === "yes"){
            callback_if_yes().then(resolve)
        } else if (str === "no"){
            callback_if_no().then(resolve)
            // console.log("callback no")
            // resolve()
        } else {
            console.log("Bad input")
            // callback_if_none()
            console.log("callback bad input")
            resolve()
        }
    })
}





async function jarnac(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Jarnac (yes / no) : ", (str) => {
            check_yes_no(str, question_jarnac_line, resolve, jarnac).then(resolve)
        })


    })

}



async function question_jarnac_line(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Chose a line (number / no) : ", (str) => {
            if (str === ""){
                jarnac().then(resolve)
            } else if (!isNaN(Number(str)) && Number.isInteger(Number(str))){
                if (0 <= Number(str) && Number(str) < grids[(turn+1)%2].length){
                    input_word(hands[(turn+1)%2], grids[(turn+1)%2][Number(str)], question_jarnac_line).then(resolve)
                } else {
                    console.log("Bad line")
                    question_jarnac_line().then(resolve)
                }
            } else if (str === "no"){
                input_word(hands[(turn+1)%2], [], question_jarnac_line).then(resolve)

            } else {
                console.log("Bad input")
                question_jarnac_line().then(resolve)
            }
        })
    })
}



async function input_word(hand_letters, grid_letters, back_callback){

    return new Promise((resolve, reject) => {


        readline_interface.question("Input a word : ", (str) => {

            if (str === ""){
                back_callback().then(resolve)
            } else {
                if (str.length > grid_letters.length && all_char_in_str(grid_letters, str) && all_char_in_str([...str], grid_letters.concat(hand_letters)) && str.length >= 3 && str.length <= 9){
                    console.log("Mot valide")
                    resolve()
                } else {
                    console.log("Mot non valide")
                    resolve()
                }
            }

        })
    })
}


function all_char_in_str(list, str){
    let returned = true
    list.forEach(element => {count(list, element) > count([...str], element) ? returned = false : returned=true})
    return returned
}

function count(list, element){
    return list.reduce((count, list_element) => {
        return list_element === element ? count + 1 : count;
    }, 0);
}


function test_line_number_input(str, player_number){

    if (Number.isInteger(Number(str))){
        return Number(str) < hands[player_number].length;
    } else {
        return false
    }
}

async function start_turn(){
    console.log("Player 1 :")

    await jarnac()

    if (!end){
        turn += 1
        start_turn()
    } else {
        readline_interface.close()
    }
}

start_turn()


// async function question_exchange_letters()


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