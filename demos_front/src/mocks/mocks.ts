import { Class, User, Tokens } from '@/types/types';

export const mockTrainer: User = {
  id: 'trainer1',
  name: 'trainer',
  email: 'trainer@mail.ru',
  isTrainer: true,
};

export const mockUser: User = {
  id: 'user1',
  name: 'user',
  email: 'user@mail.ru',
  isTrainer: false,
};

export const mockTokens: Tokens = {
  accessToken: "mock-access-token",
  refreshToken: "mock-refresh-token"
};

export const mockClasses: Class[] = [
    {
        id: 'class1',
        date: '2024-07-01',
        time: '18:00',
        location: 'Dance Studio A',
        price: 20,
        trainer: {
            id: 'trainer1',
            name: 'trainer',
        },
    },
    {
        id: 'class2',
        date: '2024-07-02',
        time: '19:00',
        location: 'Dance Studio B',
        price: 25,
        trainer: {
            id: 'trainer1',
            name: 'trainer',
        },
    }
];

export const mockEnrollments: Class[] = [
    {
        id: 'class1',
        date: '2024-07-01',
        time: '18:00',
        location: 'Dance Studio A',
        price: 20,
        trainer: {
            id: 'trainer1',
            name: 'trainer',
        },
        enrolledStudents: ['user1'],
    }
]