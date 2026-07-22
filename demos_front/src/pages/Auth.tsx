import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '@/context/AuthContext';
import Navbar from '@/components/Navbar';
import { Button } from '@/components/ui/button';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Switch } from '@/components/ui/switch';
import Logo from "@/assets/logo-clean.svg?react";
import { Label } from '@/components/ui/label';

const Auth = () => {
  const navigate = useNavigate();
  const { login, register, isAuthenticated } = useAuth();
  const [isLoading, setIsLoading] = useState(false);
  const [isTrainer, setIsTrainer] = useState(false);
  const [activeTab, setActiveTab] = useState("login");

  // Redirect if already logged in
  React.useEffect(() => {
    if (isAuthenticated) {
      navigate('/dashboard');
    }
  }, [isAuthenticated, navigate]);

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoading(true);
    
    const formData = new FormData(e.currentTarget);
    const email = formData.get('email') as string;
    const password = formData.get('password') as string;
    
    try {
      await login(email, password);
      navigate('/dashboard');
    } catch (error) {
      console.error('Login error:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleRegister = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoading(true);
    
    const formData = new FormData(e.currentTarget);
    const name = formData.get('name') as string;
    const email = formData.get('email') as string;
    const password = formData.get('password') as string;
    
    try {
      await register(email, password, name, isTrainer);
      navigate('/dashboard');
    } catch (error) {
      console.error('Registration error:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col">
      <Navbar />
      
      <main className="flex-grow flex items-center justify-center p-4 mt-16">
        <div className="w-full max-w-md mx-auto">
          <div className="card p-8">
            <h1 className="text-2xl font-bold text-center mb-5">
              <Logo className="mx-auto mb-10 mt-5 h-20 w-auto" />
              {/*<span className="block text-white text-lg font-normal mt-1">{activeTab === "login" ? "Вход в аккаунт" : "Создание аккаунта"}</span>*/}
            </h1>
            
            <Tabs value={activeTab} onValueChange={setActiveTab} className="w-full">
              <TabsList className="grid w-full grid-cols-2 mb-8">
                <TabsTrigger value="login">Вход</TabsTrigger>
                <TabsTrigger value="register">Регистрация</TabsTrigger>
              </TabsList>
              
              <TabsContent value="login">
                <form onSubmit={handleLogin} className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="login-email">Почта</Label>
                    <input
                      id="login-email"
                      name="email"
                      type="email"
                      required
                      placeholder="your@email.com"
                      className="input-field w-full"
                    />
                  </div>
                  
                  <div className="space-y-2">
                    <Label htmlFor="login-password">Пароль</Label>
                    <input
                      id="login-password"
                      name="password"
                      type="password"
                      required
                      placeholder="••••••••"
                      className="input-field w-full"
                    />
                  </div>
                  
                  <Button 
                    type="submit" 
                    className="w-full mt-6 bg-demokrat-purple hover:bg-demokrat-purple/90"
                    disabled={isLoading}
                  >
                    {isLoading ? "Вход..." : "Войти"}
                  </Button>
                </form>
              </TabsContent>
              
              <TabsContent value="register">
                <form onSubmit={handleRegister} className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="register-name">Имя</Label>
                    <input
                      id="register-name"
                      name="name"
                      type="text"
                      required
                      placeholder="Your name"
                      className="input-field w-full"
                    />
                  </div>
                  
                  <div className="space-y-2">
                    <Label htmlFor="register-email">Почта</Label>
                    <input
                      id="register-email"
                      name="email"
                      type="email"
                      required
                      placeholder="your@email.com"
                      className="input-field w-full"
                    />
                  </div>
                  
                  <div className="space-y-2">
                    <Label htmlFor="register-password">Пароль</Label>
                    <input
                      id="register-password"
                      name="password"
                      type="password"
                      required
                      placeholder="••••••••"
                      className="input-field w-full"
                    />
                  </div>
                  
                  <div className="flex items-center justify-between pt-2">
                    <Label htmlFor="role-switch" className="cursor-pointer">
                      Зарегистрироваться как тренер
                    </Label>
                    <Switch 
                      id="role-switch" 
                      checked={isTrainer} 
                      onCheckedChange={(checked) => setIsTrainer(checked)}
                    />
                  </div>
                  
                  <Button 
                    type="submit" 
                    className="w-full mt-6 bg-demokrat-purple hover:bg-demokrat-purple/90"
                    disabled={isLoading}
                  >
                    {isLoading ? "Создание..." : "Создать аккаунт"}
                  </Button>
                </form>
              </TabsContent>
            </Tabs>
          </div>
        </div>
      </main>
    </div>
  );
};

export default Auth;
