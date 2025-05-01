
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '@/context/AuthContext';
import Navbar from '@/components/Navbar';
import { Button } from '@/components/ui/button';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Switch } from '@/components/ui/switch';
import { Label } from '@/components/ui/label';
import { UserRole } from '@/types';

const Auth = () => {
  const navigate = useNavigate();
  const { login, register, isAuthenticated } = useAuth();
  const [isLoading, setIsLoading] = useState(false);
  const [role, setRole] = useState<UserRole>('student');

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
      await register(email, password, name, role);
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
            <h1 className="text-2xl font-bold text-center mb-8">
              <span className="graffiti-text">DEMOKRAT</span>
              <span className="block text-white text-lg font-normal mt-1">Account Access</span>
            </h1>
            
            <Tabs defaultValue="login" className="w-full">
              <TabsList className="grid w-full grid-cols-2 mb-8">
                <TabsTrigger value="login">Login</TabsTrigger>
                <TabsTrigger value="register">Register</TabsTrigger>
              </TabsList>
              
              <TabsContent value="login">
                <form onSubmit={handleLogin} className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="login-email">Email</Label>
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
                    <Label htmlFor="login-password">Password</Label>
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
                    {isLoading ? "Signing in..." : "Sign in"}
                  </Button>
                </form>
              </TabsContent>
              
              <TabsContent value="register">
                <form onSubmit={handleRegister} className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="register-name">Name</Label>
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
                    <Label htmlFor="register-email">Email</Label>
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
                    <Label htmlFor="register-password">Password</Label>
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
                      Register as a trainer
                    </Label>
                    <Switch 
                      id="role-switch" 
                      checked={role === 'trainer'} 
                      onCheckedChange={(checked) => setRole(checked ? 'trainer' : 'student')}
                    />
                  </div>
                  
                  <Button 
                    type="submit" 
                    className="w-full mt-6 bg-demokrat-purple hover:bg-demokrat-purple/90"
                    disabled={isLoading}
                  >
                    {isLoading ? "Creating account..." : "Create account"}
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
