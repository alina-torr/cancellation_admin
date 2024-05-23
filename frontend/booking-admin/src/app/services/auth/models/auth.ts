export interface LoginData {
  login: string;
  password: string;
}

export interface RegisterData {
  login: string;
  password: string;
  hotelName: string;
  hotelCity: string;
  hotelCountry: string;
}

export interface LoginResponse {
  access_token: string;
}
