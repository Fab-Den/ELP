module Test exposing (..)


import Browser
import Html exposing (..)
import Html.Events exposing (..)
import Random
import Http
import Array
import List



-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL

type State = Loading
    | Failure
    | Success String

type alias Model =
  {
  word_list : List String,
  current_word : String,
  random_index : Int
  }



init : () -> (Model, Cmd Msg)
init _ =
  ( Model [] "" 0
  , Http.get
          { url = "https://perso.liris.cnrs.fr/tristan.roussillon/GuessIt/thousand_words_things_explainer.txt"
          , expect = Http.expectString GotText
          }
  )


roll: Int -> Random.Generator Int
roll a =
  Random.int 0 a

-- UPDATE


type Msg
  = Roll
  | GetElementFromList Int
  | GotText (Result Http.Error String)


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    Roll ->
      ( model
      , Random.generate GetElementFromList (roll (List.length model.word_list))
      )

    GetElementFromList index ->
      (
        case model.word_list |> Array.fromList |> Array.get index of
          Maybe.Just a -> {model|current_word = a}
          _ -> model
        , Cmd.none
      )

    GotText result ->
        case result of
            Ok fullText ->
                ( {model|word_list = String.split " " fullText} , Cmd.none)

            Err _ ->
                (model, Cmd.none)

-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW
view : Model -> Html Msg
view model =
  div []
    [ h1 [] [ text (model.current_word) ]
    , button [ onClick Roll ] [ text "Roll" ]
    ]