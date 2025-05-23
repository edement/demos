
import React, { createContext, useState, useContext, useEffect } from 'react';
import { User } from '../types';
import { useToast } from '@/components/ui/use-toast';
import { authService } from '@/services/api';

interface AuthContextType {
  user: User | null;
  login: (email: string, password: string) => Promise<void>;
  register: (email: string, password: string, name: string, isTrainer: boolean) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const { toast } = useToast();

  useEffect(() => {
    const storedToken = localStorage.getItem('demokratAccessToken');
    
    if (storedToken) {
      fetchCurrentUser();
    } else {
      setIsLoading(false);
    }
  }, []);

  const fetchCurrentUser = async () => {
    try {
      const userData = await authService.getCurrentUser();
      setUser(userData);
    } catch (error) {
      console.error('Failed to fetch user data:', error);
      localStorage.removeItem('demokratToken');
    } finally {
      setIsLoading(false);
    }
  };

  const login = async (email: string, password: string) => {
    try {
      const { user: userData, tokens } = await authService.login(email, password);
      setUser(userData);
      localStorage.setItem('demokratAccessToken', tokens.accessToken);
      localStorage.setItem('demokratRefreshToken', tokens.refreshToken);
      localStorage.setItem('demokratUser', JSON.stringify(userData));
      toast({
        title: "Вход выполнен",
        description: `С возвращением, ${userData.name}!`,
      });
    } catch (error) {
      console.error('Login failed:', error);
      toast({
        title: "Ошибка входа",
        description: "Неверные данные для входа. Пожалуйста, попробуйте снова.",
        variant: "destructive",
      });
      throw error;
    }
  };

  const register = async (email: string, password: string, name: string, isTrainer: boolean) => {
    try {
      const { user: userData, tokens } = await authService.register(email, password, name, isTrainer);
      setUser(userData);
      localStorage.setItem('demokratAccessToken', tokens.accessToken);
      localStorage.setItem('demokratRefreshToken', tokens.refreshToken);
      localStorage.setItem('demokratUser', JSON.stringify(userData));
      toast({
        title: "Регистрация завершена",
        description: `Добро пожаловать в Demokrat, ${name}!`,
      });
    } catch (error) {
      console.error('Registration failed:', error);
      toast({
        title: "Ошибка регистрации",
        description: "Не удалось создать аккаунт. Пожалуйста, попробуйте снова.",
        variant: "destructive",
      });
      throw error;
    }
  };

  const logout = () => {
    setUser(null);
    localStorage.removeItem('demokratAccessToken');
    localStorage.removeItem('demokratRefreshToken');
    localStorage.removeItem('demokratUser');
    toast({
      title: "Выход выполнен",
      description: "Вы успешно вышли из системы.",
    });
  };

  const value = {
    user,
    login,
    register,
    logout,
    isAuthenticated: !!user,
  };

  if (isLoading) {
    return <div>Загрузка...</div>;
  }

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
