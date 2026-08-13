export interface User {
  id: string;
  username: string;
  email: string;
  isTrainer: boolean;
}

export interface Class {
  id: string;
  datetime: string;
  location: string;
  price: number;
  trainer: {
    id: string;
    username: string;
  };
  enrolledStudents?: string[];
}

export interface Tokens {
  accessToken: string;
  refreshToken: string;
}