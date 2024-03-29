
// Database structure
export interface User {
    firstname: string;
    lastname: string;
    username: string;
    password: string;
    // Closet: Closet;
}
/*
export interface Closet {
    ClosetID: number;
    UserID: number;
    ClosetName: string;
    Items: Item[];
}

export interface Item {
    ItemID: number;
    ClosetID: number;
    ItemName: string;
    ItemImage: string;
    Category: string;
    ItemTags: ItemTag[];
}

// The two tables below are separate because I want a tag to be used on many different items.
// If we consolidate into one table, then a tag will be associated with only a singular item.


// General tag
export interface Tag {
    TagID: number;
    TagName: string;
}


// Associate tag to a specific item
export interface ItemTag {
    ItemID: number;
    TagID: number;
}

export const ITEMTAGS: ItemTag[] = [
    { ItemID: 1, TagID: 1 },
    { ItemID: 2, TagID: 2 }
]

export const TAGS: Tag[] = [
    { TagID: 1, TagName: "Jeans" },
    { TagID: 2, TagName: "Blue" },
    { TagID: 3, TagName: "T-shirts" }
]

export const ITEMS: Item[] = [
    { ItemID: 1, ClosetID: 1, ItemName: "Dark wash", ItemImage: "https://media.everlane.com/image/upload/c_fill,w_640,ar_1:1,q_auto,dpr_1.0,g_face:center,f_auto,fl_progressive:steep/i/8318fe58_a828", Category: "Bottoms", ItemTags: [ITEMTAGS[0], ITEMTAGS[1]] },
    { ItemID: 2, ClosetID: 1, ItemName: "Relaxed tee", ItemImage: "https://media.everlane.com/image/upload/c_fill,w_750,ar_1:1,q_auto,dpr_1.0,g_face:center,f_auto,fl_progressive:steep/i/20def3b7_19f0", Category: "Tops", ItemTags: [ITEMTAGS[2]] }
]

export const CLOSETS: Closet[] = [
    { ClosetID: 1, UserID: 1, ClosetName: "Michelle's Closet", Items: [] }
]
*/
/*
export const USERS: User[] = [
    { UserID: 1, FirstName: "Michelle", LastName: "Taing", Username: "michelletaing", Password: "hello1234", Closet: CLOSETS[0] }
]
*/

