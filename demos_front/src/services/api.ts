import { Class, User, Tokens } from '@/types/types';
import { mockTrainer, mockUser, mockTokens, mockClasses, mockEnrollments } from '@/mocks/mocks';

const API_BASE_URL = 'http://localhost:5088';
const useMock = import.meta.env.VITE_USE_MOCK === 'true';

// Функция для получения токена из localStorage
const getAccessToken = () => localStorage.getItem('demokratAccessToken');
const getRefreshToken = () => localStorage.getItem('demokratRefreshToken');

// Универсальная функция для запросов к API
const apiRequest = async <T>(
  endpoint: string, 
  method: string = 'GET', 
  data?: any
): Promise<T> => {
  const url = `${API_BASE_URL}${endpoint}`;
  
  const headers: HeadersInit = {
    'Content-Type': 'application/json',
  };
  
  // Добавляем токен авторизации, если он есть
  const token = getAccessToken();
  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  
  const options: RequestInit = { 
    method, 
    headers,
    credentials: 'include',
  };
  
  if (data && method !== 'GET') {
    options.body = JSON.stringify(data);
  }
  
  const response = await fetch(url, options);
  
  // Если ответ не успешен, генерируем ошибку
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText || `API Error: ${response.status}`);
  }

  // Если access просрочен
  if (response.status === 401) {
    //обновляем токены
    //повторный запрос
  }
  
  // Для ответов со статусом 204 (No Content) возвращаем null
  if (response.status === 204) {
    return null as T;
  }
  
  return response.json();
};

// Сервисы для работы с авторизацией
export const authService = {
  // Регистрация пользователя
  register: (email: string, password: string, name: string, isTrainer: boolean) => {
      if (useMock) {
        return Promise.resolve({
          user: mockTrainer,
          tokens: mockTokens
        });
      }

      return apiRequest<{ user: User; tokens: Tokens }>('/api/auth/registration', 'POST', {
        email, password, name, isTrainer
      });
    },
  
  // Вход пользователя
  login: (email: string, password: string) => {
      if (useMock) {
        return Promise.resolve({
          user: mockUser,
          tokens: mockTokens
        });
      }

      return apiRequest<{ user: User; tokens: Tokens }>('/api/auth/login', 'POST', {
        email,
        password
      });
  },
    
  // Получение данных текущего пользователя
  getCurrentUser: () => {
    if (useMock) {
      return Promise.reject({
        status: 401,
        message: "Unauthorized"
      });
    }
    apiRequest<User>('/api/users/me');
  },

  // Обновление токенов
  refreshTokens: () => 
    apiRequest<Tokens>('/api/auth/refresh'),
};

// Сервисы для работы с занятиями
export const classesService = {
  // Получение всех занятий
  getAllClasses: () => {
    if (useMock) {
      return Promise.resolve(mockClasses);
    }
    apiRequest<Class[]>('/api/classes');
  },
  
  // Получение занятия по ID
  getClassById: (id: string) => 
    apiRequest<Class>(`/api/classes/${id}`),
  
  // Создание нового занятия
  createClass: (classData: Omit<Class, 'id' | 'trainer' | 'enrolledStudents'>) => 
    apiRequest<null>('/api/classes', 'POST', classData),
  
  // Обновление занятия
  updateClass: (id: string, classData: Partial<Class>) => 
    apiRequest<Class>(`/api/classes/${id}`, 'PUT', classData),
  
  // Удаление занятия
  deleteClass: (id: string) => 
    apiRequest<null>(`/api/classes/${id}`, 'DELETE'),

  // Получение занятий тренера
    getTrainerClasses: () => {
      if (useMock) {
        let trainerClasses: Class[] = [];
        for(let danceClass of mockClasses) {
          if(danceClass.trainer.id === mockTrainer.id) {
            trainerClasses.push(danceClass);
          }
        }
        return trainerClasses;
      }
      return apiRequest<Class[]>('/api/trainers/classes')
    },
    
  // Получение учеников тренера НЕ ИСПОЛЬЗУЕТСЯ                                               
  getTrainerStudents: () => 
    apiRequest<User[]>('/api/trainers/students')
};

// Сервисы для работы с записью на занятия 
export const enrollmentService = {
  // Запись на занятие
  enrollToClass: (classId: string) => 
    apiRequest<null>('/api/classes/enrollments', 'POST', { classId }),
  
  // Получение записей пользователя
  getUserEnrollments: () => {
    if (useMock) {
      let userEnrollments: Class[] = [];
      for(let enroll of mockEnrollments) {
        if(enroll.enrolledStudents.includes(mockUser.id)) {
          userEnrollments.push(enroll);
        }
      }
      return userEnrollments
    }
    return apiRequest<Class[]>('/api/classes/enrollments');
  },
  
  // Отмена записи
  cancelEnrollment: (classId: string) => 
    apiRequest<null>(`/api/classes/enrollments/${classId}`, 'DELETE')
};
