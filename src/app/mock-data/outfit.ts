import { Item } from "./item"

export interface Outfit {
    Tops: Item;
    TopID: number;
    Bottoms: Item;
    BottomID: number;
    OnePieces: Item;
    OnePieceID: number;
    Accessories: Item;
    AccessoriesID: number;
    Shoes: Item;
    ShoesID: number;
}