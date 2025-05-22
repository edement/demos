export interface User {
  id: string;
  name: string;
  email: string;
  isTrainer: boolean;
}

export interface Class {
  id: string;
  date: string;
  time: string;
  location: string;
  price: number;
  trainer: {
    id: string;
    name: string;
  };
  enrolledStudents?: string[];
}

export interface Tokens {
  accessToken: string;
  refreshToken: string;
}