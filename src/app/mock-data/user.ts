export interface User {
    ID?: number;
    firstname: string;
    lastname: string;
    username: string;
    password: string;
    // Closet: Closet;
}

export interface LoginRequest {
    username: string;
    password: string;
}