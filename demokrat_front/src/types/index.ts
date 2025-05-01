
export type UserRole = 'student' | 'trainer';

export interface User {
  id: string;
  name: string;
  email: string;
  role: UserRole;
}

export interface DanceClass {
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
