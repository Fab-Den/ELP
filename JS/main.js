const { fonction1, fonction2 } = require('./test');
fonction1();


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

let grids = [[], [["B", "O", "N", "J", "O", "U"]]]

let hands= [[], ["O", "N", "B"]]

let turn = 0;

let end= false

let number_jarnac = 0


function step1(){
    // jarnac ou pas
    if (turn !== 0){
        step1_chose_jarnac().then(step2)

    } else {
        step2()
    }

}


function step2(){
    // tirer 1 lettre ou echanger

    if (turn > 1){
        step2_chose_exchange().then(step3)
    } else {
        step3()
    }

}


function step3(){
    // jouer mot
    step3_chose_do().then(() => {turn += 1; change_turn()})

}

function change_turn(){

    if (grids[turn%2].length !== 8){
        console.log("Turn of player " + ((turn%2)+1).toString() + " : ")
        step1()
    } else {
        // faire des choses pour compter le score + affichage gagnant
    }


}

function step3_chose_do(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Do you want to input a word ? (yes/no) ", (str) => {
            if (str === "yes"){
                step3_chose_line().then(resolve)
            } else if (str === "no"){
                resolve()
            } else {
                step3_chose_do().then(resolve)
            }
        })
    })
}


function step3_chose_line(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Chose a line (number / no) : ", (str) => {
            if (str === ""){
                step3_chose_do().then(resolve)
            } else if (!isNaN(Number(str)) && Number.isInteger(Number(str))){
                if (0 <= Number(str) && Number(str) < grids[turn%2].length){
                    step3_input_word(hands[turn%2], grids[turn%2][Number(str)]).then(resolve)
                } else {
                    console.log("Bad line")
                    step3_chose_line().then(resolve)
                }
            } else if (str === "no"){
                if (grids[turn%2].length !== 8){
                    step3_input_word(hands[turn%2], []).then(resolve)
                }
                else {
                    console.log("There is no more line on which we can add a word")
                }

            } else {
                console.log("Bad input")
                question_jarnac_line().then(resolve)
            }
        })
    })
}

function step3_input_word(hand_letters, grid_letters){
    return new Promise((resolve, reject) => {
        readline_interface.question("Chose a word : ", (str) => {
            console.log("can_write_with([...str], hand_letters.concat(grid_letters)) && can_write_with(grid_letters, [...str]) : ", can_write_with([...str], hand_letters.concat(grid_letters)))
            if (str === ""){
                step3_chose_line().then(resolve)
            } else if (str.length < 3 || str.length > 9){
                console.log("Bad word length")
                step3_input_word(hand_letters, grid_letters).then(resolve)
            } else if (can_write_with([...str], hand_letters.concat(grid_letters)) && can_write_with(grid_letters, [...str]) && str.length > grid_letters.length){
                // qqch quand c ok
                console.log("YOUPIII")
            } else {
                console.log("Bad word")
                step3_input_word(hand_letters, grid_letters).then(resolve)
            }
        })
    })
}



function step1_chose_jarnac(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Jarnac (yes / no) : ", (str) => {

            if (str === "yes"){
                step1_chose_line().then(resolve)

            } else if (str === "no"){
                resolve()

            } else {
                console.log("Bad input")
                step1_chose_jarnac().then(resolve)
            }
        })
    })
}



function step1_chose_line(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Chose a line (number / no) : ", (str) => {
            if (str === ""){
                step1_chose_jarnac().then(resolve)

            } else if (!isNaN(Number(str)) && Number.isInteger(Number(str))){

                if (0 <= Number(str) && Number(str) < grids[(turn+1)%2].length){
                    step1_input_word(hands[(turn+1)%2], grids[(turn+1)%2][Number(str)], Number(str)).then(resolve)
                } else {
                    console.log("Bad line")
                    question_jarnac_line().then(resolve)
                }
            } else if (str === "no"){
                step1_input_word(hands[(turn+1)%2], []).then(resolve)

            } else {
                console.log("Bad input")
                question_jarnac_line().then(resolve)
            }
        })
    })
}



function step1_input_word(hand_letters, grid_letters, number){

    return new Promise((resolve, reject) => {


        readline_interface.question("Input a word : ", (str) => {

            if (str === ""){
                step1_chose_line().then(resolve)
            } else {
                if (str.length > grid_letters.length && can_write_with(grid_letters, [...str]) && can_write_with([...str], grid_letters.concat(hand_letters)) && str.length >= 3 && str.length <= 9){
                    console.log("Mot valide")
                    // fonction qui vient du fichier functions
                    // grids, hands = updateJarnac(grids, number, turn, hands, grid_letters, [...str])
                    resolve()
                } else {
                    console.log("Mot non valide")
                    step1_input_word(hand_letters, grid_letters).then(resolve)
                }
            }

        })
    })
}




function can_write_with(list1, list2){
    for (let i = 0; i < list1.length; i++){
        if (count(list1, list1[i]) > count(list2, list1[i])){
            return false
        }
    }
    return true
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

function step2_chose_exchange(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Exchange 3 letters instead of drawing ? (yes/no) ", (str) => {
            if (str === "yes"){
                step2_chose_letters().then(resolve)

            } else if (str === "no"){
                resolve()
            }
            else {
                step2_chose_exchange().then(resolve)
            }
        })

    })
}

function step2_chose_letters(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Chose 3 letters to exchange : ", (str) => {
            if (str === ""){
                step2_chose_exchange().then(resolve)

            } else if (str.length !== 3){
                console.log("You must select exactly 3 letters (no spaces)")
                step2_chose_letters().then(resolve)

            } else {
                if (can_write_with([...str], hands[turn])){
                    // DO SOMETHING
                    resolve()
                } else {
                    console.log("Invalid letters")
                    step2_chose_letters().then(resolve)
                }

            }
        })
    })
}


function display_grid(index){

    let border_top = "┌" + function (){let s = ""; for (let i=0; i<8; i++){s+= "─┬"} return s}() + "─┐"
    let border_in = "├" + function (){let s = ""; for (let i=0; i<8; i++){s+= "─┼"} return s}() + "─┤"
    let border_bottom = "└" + function (){let s = ""; for (let i=0; i<8; i++){s+= "─┴"} return s}() + "─┘"

    let text = "GRID PLAYER " + (index%2+1).toString() + " : \n"
    text += border_top + "\n"
    for (let i = 0; i<8; i++){
        text += "│"
        if (i < grids[index].length){
            for (let j = 0; j < 9; j++){
                if (j<grids[index][i].length){
                    text += grids[index][i][j] + "|"
                } else {
                    text += " |"
                }

            }
        } else {
            for (let j = 0; j < 9; j++){
                text += " |"
            }
        }

        text += "\n" + (i !== 7 ? border_in : border_bottom) + "\n"

    }
    console.log(text)
}

function display_letters(index){
    let text = "LETTERS OF PLAYER " + (index%2+1).toString() + " : "

    for (let i = 0; i < hands[index].length; i++){
        text += hands[index][i] + " "
    }
    console.log(text)
}

// change_turn()

display_grid(1)
display_letters(1)

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