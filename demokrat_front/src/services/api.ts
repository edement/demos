
import { DanceClass, User, UserRole } from '@/types';

const API_BASE_URL = 'http://localhost:5271';

// Функция для получения токена из localStorage
const getToken = () => localStorage.getItem('demokratToken');

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
  const token = getToken();
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
  
  // Для ответов со статусом 204 (No Content) возвращаем null
  if (response.status === 204) {
    return null as T;
  }
  
  return response.json();
};

// Сервисы для работы с авторизацией
export const authService = {
  // Регистрация пользователя
  register: (email: string, password: string, name: string, role: UserRole) => 
    apiRequest<{ user: User; token: string }>('/api/auth/register', 'POST', {
      email,
      password,
      name,
      role
    }),
  
  // Вход пользователя
  login: (email: string, password: string) => 
    apiRequest<{ user: User; token: string }>('/api/auth/login', 'POST', {
      email,
      password
    }),
    
  // Получение данных текущего пользователя
  getCurrentUser: () => 
    apiRequest<User>('/api/users/me')
};

// Сервисы для работы с занятиями
export const classesService = {
  // Получение всех занятий
  getAllClasses: () => 
    apiRequest<DanceClass[]>('/api/classes'),
  
  // Получение занятия по ID
  getClassById: (id: string) => 
    apiRequest<DanceClass>(`/api/classes/${id}`),
  
  // Создание нового занятия
  createClass: (classData: Omit<DanceClass, 'id' | 'trainer' | 'enrolledStudents'>) => 
    apiRequest<DanceClass>('/api/classes', 'POST', classData),
  
  // Обновление занятия
  updateClass: (id: string, classData: Partial<DanceClass>) => 
    apiRequest<DanceClass>(`/api/classes/${id}`, 'PUT', classData),
  
  // Удаление занятия
  deleteClass: (id: string) => 
    apiRequest<null>(`/api/classes/${id}`, 'DELETE'),

  // Получение занятий тренера
  getTrainerClasses: () => 
    apiRequest<DanceClass[]>('/api/trainers/classes'),
    
  // Получение учеников тренера
  getTrainerStudents: () => 
    apiRequest<User[]>('/api/trainers/students')
};

// Сервисы для работы с записью на занятия 
export const enrollmentService = {
  // Запись на занятие
  enrollToClass: (classId: string) => 
    apiRequest<{ id: string; classId: string; userId: string; enrollmentDate: string }>('/api/enrollments', 'POST', { classId }),
  
  // Получение записей пользователя
  getUserEnrollments: () => 
    apiRequest<DanceClass[]>('/api/enrollments'),
  
  // Отмена записи
  cancelEnrollment: (classId: string) => 
    apiRequest<null>(`/api/enrollments/${classId}`, 'DELETE')
};
